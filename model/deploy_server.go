package model

import (
	"log"
	"time"
)

type DeployServer struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Branch    string    `json:"branch"`
	Dir       string    `json:"dir"`
	Cmd       string    `json:"cmd"`
	CreatedAt time.Time `json:"created_at"`
}

func FindServerByNameAndBranch(name, branch string) *DeployServer {
	db := GetDBConn()
	deployServer := DeployServer{}
	row := db.QueryRow("SELECT id FROM deploy_servers WHERE name = $1 and branch = $2", name, branch)
	queryErr := row.Scan(&deployServer.ID)
	if queryErr != nil {
		log.Println(queryErr)
	}
	return &deployServer
}
