package node

import (
	"database/sql"
	"github.com/workhorse/api"
	"github.com/workhorse/apiserver/pkg/db"
	"time"
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

func (ns *NodeInfoService) ListNodes() ([]api.NodeInfo, error) {

	var nodeList []api.NodeInfo

	err := db.Run(func(db *sql.DB) error {
		selectStmt := `
	select id, name, last_heart_beat
	from node_info
	`
		rows, err := db.Query(selectStmt)
		defer rows.Close()

		if err != nil {
			return err
		}

		for rows.Next() {
			var id int
			var name string
			var lastHeartBeat time.Time

			err := rows.Scan(&id, &name, &lastHeartBeat)
			if err != nil {
				return err
			}

			nodeList = append(nodeList,
				api.NodeInfo{
					Id:              id,
					Name:            name,
					LastHeartBeatTS: lastHeartBeat,
				})
		}

		return nil
	})

	return nodeList, err
}
