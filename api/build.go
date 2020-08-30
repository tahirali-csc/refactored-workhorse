package api

import "time"

type Build struct {
	Id        int    `json:"id"`
	Status    string `json:"status"`
	ProjectId int64  `json:"projectId"`
	CreatedTs time.Time
	StartTs   time.Time
	EndTs     time.Time
	Steps     []BuildStep `json:"steps"`
}
