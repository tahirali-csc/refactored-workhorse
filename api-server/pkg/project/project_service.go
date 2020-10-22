package project

import (
	"database/sql"
	"github.com/workhorse/api"
	"github.com/workhorse/apiserver/pkg/db"
)

type ProjectService struct {
}

func (ps ProjectService) Create(project *api.Project) (*api.Project, error) {
	res := &api.Project{}

	err := db.Run(func(db *sql.DB) error {
		insertStmt := `
	insert into project
	(name, private_key, clone_url)
	values($1, $2, $3)
	RETURNING id
	`

		id := 0
		err := db.QueryRow(insertStmt, project.Name, project.PrivateKey, project.CloneURL).Scan(&id)
		if err != nil {
			return err
		}

		res.Id = id
		res.Name = project.Name
		res.PrivateKey = project.PrivateKey
		res.CloneURL = project.CloneURL
		return nil
	})

	return res, err
}
