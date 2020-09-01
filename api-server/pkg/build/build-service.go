package build

import (
	"context"
	"database/sql"
	"fmt"
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
	// if row == nil {
	// 	panic(row)
	// }

	for _, step := range build.Steps {
		insertStmt1 := `
		INSERT INTO build_steps (build_id, name, image,status,created_ts)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`

		var buildStepId int
		tx.QueryRow(insertStmt1, buildId, step.Name, step.Image, "Pending", time.Now()).Scan(&buildStepId)
		// if row == nil {
		// 	panic(row)
		// }

		for _, cmd := range step.Commands {
			insertStmt2 := `
			INSERT INTO build_steps_command (step_id, command)
			VALUES ($1, $2)
			`

			_ , err := tx.Exec(insertStmt2, buildStepId, cmd.Command)
			if err != nil {
				panic (err)
			}
		}

	}

	tx.Commit()

}
