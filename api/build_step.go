package api

import "time"

type BuildStep struct {
	Id        int
	BuildId   int
	Name      string `json:"name"`
	Image     string `json:"image"`
	Status    string
	CreatedTs time.Time
	StartTs   time.Time
	EndTs     time.Time
	Commands  []BuildStepCommand `json:"commands"`
}
