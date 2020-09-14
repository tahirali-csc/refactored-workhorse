package node

import (
	"database/sql"
	"github.com/workhorse/api"
	"github.com/workhorse/apiserver/pkg/db"
)

type NodeInfoService struct {
}

func (ns *NodeInfoService) UpdateNode(info *api.NodeInfo) error {

	return db.Run(func(db *sql.DB) error {
		insertStmt := `
	insert into node_info
	(name, last_heart_beat)
	values($1, $2)
	on conflict(name) do update set last_heart_beat = EXCLUDED.last_heart_beat
	`

		_, err := db.Exec(insertStmt, info.Name, info.LastHeartBeatTS)
		if err != nil {
			return err
		}

		return nil
	})
}
