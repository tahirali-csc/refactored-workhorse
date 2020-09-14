package api

import "time"

type NodeInfo struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	LastHeartBeatTS time.Time `json:"lastHeartBeatTS"`
}
