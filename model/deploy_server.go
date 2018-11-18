package model

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

var DeployLogger = logrus.New()

func init() {
	f, err := os.OpenFile("deploy.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	DeployLogger.Out = f
}

type DeployServer struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name" binding:"required"`
	Branch    string   `json:"branch" binding:"required"`
	Dir       string   `json:"dir" binding:"required"`
	Cmd       string   `json:"cmd" binding:"required"`
	CreatedAt JsonTime `json:"created_at" db:"created_at"`
}

func FindServerByNameAndBranch(name, branch string) *DeployServer {
	db := GetDBConn()
	defer db.Close()
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
	defer db.Close()
	dr := DeployServer{}
	queryErr := db.Get(&dr, "SELECT * FROM deploy_servers WHERE id=$1", id)
	if queryErr != nil {
		log.Println(queryErr)
	}
	return &dr
}

func (ds *DeployServer) Save() bool {
	db := GetDBConn()
	defer db.Close()
	if ds.CreatedAt.IsZero() {
		ds.CreatedAt = JsonTime{time.Now()}
	}
	insertSql := `
		INSERT INTO deploy_servers (name, branch, dir, cmd, created_at)
		VALUES (:name, :dir, :branch, :cmd, :created_at)
		RETURNING id
	`
	nstmt, err := db.PrepareNamed(insertSql)
	err = nstmt.Get(&ds.ID, ds)
	if err != nil {
		log.Printf("SaveDeployServer failed: %v", err)
		return false
	}
	return true
}

func (ds *DeployServer) Update() bool {
	db := GetDBConn()
	defer db.Close()

	updateSql := `
		UPDATE deploy_servers
		SET name=:name, branch=:branch, dir=:dir, cmd=:cmd
		WHERE id = :id
	`
	_, err := db.NamedExec(updateSql, ds)
	if err != nil {
		log.Printf("updateDeployServer failed: %v", err)
		return false
	}
	return true
}

func AllDeployServers() []DeployServer {
	selectSql := `SELECT * FROM deploy_servers ORDER BY id asc`
	deployServers := []DeployServer{}
	db := GetDBConn()
	defer db.Close()
	err := db.Select(&deployServers, selectSql)

	if err != nil {
		fmt.Println(err)
	}
	return deployServers
}

func (ds *DeployServer) PaginatedDeployRecords(page, perPage int) []DeployRecord {
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 24
	}

	selectSql := `SELECT * FROM deploy_records WHERE server_id=$1 ORDER BY id desc LIMIT $2 OFFSET $3`
	deployRecords := []DeployRecord{}
	db := GetDBConn()
	defer db.Close()
	err := db.Select(&deployRecords, selectSql, ds.ID, perPage, (page-1)*perPage)
	if err != nil {
		fmt.Println(err)
	}
	return deployRecords
}

func (ds DeployServer) TotalDeployRecordsCount() int {
	querySql := `SELECT count(*) FROM deploy_records WHERE server_id=$1`
	db := GetDBConn()
	defer db.Close()
	var totalCount int
	db.QueryRow(querySql, ds.ID).Scan(&totalCount)
	return totalCount
}

func (ds *DeployServer) runCmd() bool {
	contextLogger := DeployLogger.WithFields(logrus.Fields{
		"DeployServer": ds.ID,
	})

	contextLogger.Info("Deploy starting")
	cmd := exec.Command(ds.Cmd)

	cmd.Dir = ds.Dir
	err := cmd.Run()
	if err != nil {
		contextLogger.Warn(err)
		return false
	}
	contextLogger.Info("Deploy done")
	return true
}
