package model

import (
	"bytes"
	"log"
	"os/exec"
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

func FindDeployServerByID(id int64) *DeployServer {
	db := GetDBConn()
	dr := DeployServer{}
	row := db.QueryRow("SELECT * FROM deploy_servers WHERE id=$1", id)
	queryErr := row.Scan(&dr.ID, &dr.Name, &dr.Branch, &dr.Dir, &dr.Cmd, &dr.CreatedAt)
	if queryErr != nil {
		log.Println(queryErr)
	}
	return &dr
}

func (ds *DeployServer) runCmd() bool {
	cmd := exec.Command(ds.Cmd)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Dir = ds.Dir
	err := cmd.Run()
	if err != nil {
		log.Printf("DeployServer #%d failed to run: %v\n", ds.ID, err)
		return false
	}

	// fmt.Printf("out:\n%s\nerr:\n%s\n", stdout.String(), stderr.String())

	return true
}
