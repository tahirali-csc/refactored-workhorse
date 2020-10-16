package api

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type BuildStep struct {
	Id        int
	BuildId   int
	Name      string `json:"name"`
	Image     string `json:"image"`
	Status    string
	CreatedTs time.Time
	StartTs   *time.Time
	EndTs     *time.Time
	LogInfo   LogStorageProperties
	Commands  []BuildStepCommand `json:"commands"`
	Node      NodeInfo
}

type LogStorageProperties map[string]interface{}

func (p LogStorageProperties) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *LogStorageProperties) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		//return errors.New("Type assertion .([]byte) failed.")
		return nil
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}
