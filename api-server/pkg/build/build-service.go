package build

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/workhorse/apiserver/pkg/db"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/workhorse/api"
	"github.com/workhorse/apiserver/pkg/config"
)

type BuildService struct {
}

func (bs *BuildService) StartNewBuild(build api.Build) {

	config := config.GetAppConfig()
	var conninfo string = fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable",
		config.Database.Name, config.Database.User, config.Database.Password, config.Database.Host)

	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	insertStmt := `
	INSERT INTO build (status, project_id, created_ts)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	buildId := 0
	tx.QueryRow(insertStmt, "Pending", build.ProjectId, time.Now()).Scan(&buildId)

	for _, step := range build.Steps {
		insertStmt1 := `
		INSERT INTO build_steps (build_id, name, image,status,created_ts)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`

		var buildStepId int
		tx.QueryRow(insertStmt1, buildId, step.Name, step.Image, "Pending", time.Now()).Scan(&buildStepId)

		for _, cmd := range step.Commands {
			insertStmt2 := `
			INSERT INTO build_steps_command (step_id, command)
			VALUES ($1, $2)
			`

			_, err := tx.Exec(insertStmt2, buildStepId, cmd.Command)
			if err != nil {
				panic(err)
			}
		}

	}

	tx.Commit()

}

func (bs *BuildService) BindToNode(binding *api.BuildNodeBinding) {
	config := config.GetAppConfig()
	var conninfo string = fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable",
		config.Database.Name, config.Database.User, config.Database.Password, config.Database.Host)

	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	insertStmt := `
	INSERT INTO build_node_binding (build_id, node_id)
	VALUES ($1, $2)
	RETURNING id
	`

	_, err = db.Exec(insertStmt, binding.BuildId, binding.NodeId)
	if err != nil {
		log.Println(err)
	}

}

func (bs *BuildService) UpdateBuildStepStatus(stepId int, status string) {
	config := config.GetAppConfig()
	var conninfo string = fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable",
		config.Database.Name, config.Database.User, config.Database.Password, config.Database.Host)

	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	updateStmt := `
	UPDATE build_steps 
	SET status=$2
	WHERE id=$1
	`

	_, err = db.Exec(updateStmt, stepId, status)
	if err != nil {
		log.Println(err)
	}

}

func (bs *BuildService) BindBuildStepToNode(step *api.BuildStepNodeBinding) {
	config := config.GetAppConfig()
	var conninfo string = fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable",
		config.Database.Name, config.Database.User, config.Database.Password, config.Database.Host)

	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	updateStmt := `
	INSERT INTO build_step_node_binding
	(step_id, ip_address)
	VALUES($1, $2)
	`

	_, err = db.Exec(updateStmt, step.StepId, step.IpAddress)
	if err != nil {
		log.Println(err)
	}

}

func (bs *BuildService) GetBuild(buildId int) (api.Build, error) {

	build := api.Build{}

	err := db.Run(func(db *sql.DB) error {
		selectStmt := `
select * 
from build b 
where b.id = $1
	`
		row := db.QueryRow(selectStmt, buildId)
		var id int
		var status string
		var projectId int64
		var createdTs time.Time
		var startTs, endTs sql.NullTime

		err := row.Scan(&id, &status, &projectId, &createdTs, &startTs, &endTs)
		if err != nil {
			return err
		}

		build.Id = id
		build.Status = status
		build.ProjectId = projectId
		build.CreatedTs = createdTs
		if startTs.Valid {
			build.StartTs = startTs.Time
		}
		if endTs.Valid {
			build.EndTs = endTs.Time
		}

		selectStmt = `
select * from build_steps 
where build_id=$1
	`
		rows, err := db.Query(selectStmt, buildId)
		if err != nil {
			return err
		}

		for rows.Next() {
			var id int
			var buildId int
			var name, image, status string
			var createdTs time.Time
			var startTs, endTs sql.NullTime

			err := rows.Scan(&id, &buildId, &name, &image, &status, &createdTs, &startTs, &endTs)
			if err != nil {
				return err
			}

			bs := api.BuildStep{
				Id:        id,
				BuildId:   buildId,
				Name:      name,
				Image:     image,
				Status:    status,
				CreatedTs: createdTs,
			}

			if startTs.Valid {
				bs.StartTs = startTs.Time
			}
			if endTs.Valid {
				bs.EndTs = endTs.Time
			}

			build.Steps = append(build.Steps, bs)
		}

		return nil
	})

	return build, err
}

func (bs *BuildService) GetStep(stepId int) (api.BuildStep, error) {

	step := api.BuildStep{}

	err := db.Run(func(db *sql.DB) error {
		selectStmt := `
select * 
from build_steps 
where id = $1
	`
		row := db.QueryRow(selectStmt, stepId)
		var id int
		var buildId int
		var name, image, status string
		var createdTs time.Time
		var startTs, endTs sql.NullTime

		err := row.Scan(&id, &buildId, &name, &image, &status, &createdTs, &startTs, &endTs)
		if err != nil {
			return err
		}

		step.Id = id
		step.BuildId = buildId
		step.Name = name
		step.Image = image
		step.Status = status
		step.CreatedTs = createdTs
		if startTs.Valid {
			step.StartTs = startTs.Time
		}
		if endTs.Valid {
			step.EndTs = endTs.Time
		}

		selectStmt = `
select * from build_steps_command 
where step_id=$1
	`
		rows, err := db.Query(selectStmt, stepId)
		if err != nil {
			return err
		}

		for rows.Next() {
			var id int
			var stepId int
			var command string

			err := rows.Scan(&id, &stepId, &command)
			if err != nil {
				return err
			}

			bs := api.BuildStepCommand{
				Id:      id,
				Command: command,
				StepId:  stepId,
			}

			step.Commands = append(step.Commands, bs)
		}

		return nil
	})

	return step, err
}
