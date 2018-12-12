package model

import (
	"bytes"
	"log"
	"os"
	"slp_cd/notification"
	"strings"
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

// Valid status is in ["waiting", "processing", "processed", "failed"]
type DeployRecord struct {
	ID           int64      `json:"id"`
	Status       string     `json:"status"`
	ServerID     int64      `json:"server_id" db:"server_id"`
	Compare      string     `json:"compare"`
	CreatedAt    JsonTime   `json:"created_at" db:"created_at"`
	EndedAt      JsonTime   `json:"ended_at" db:"ended_at"`
	DeployUser   DeployUser `json:"deploy_user" db:"deploy_user"`
	Stdout       NullString `json:"stdout"`
	Stderr       NullString `json:"stderr"`
	DeployServer *DeployServer
}

func (dc *DeployRecord) SetDeployServer() {
	dc.DeployServer = FindDeployServerByID(dc.ServerID)
}

func (dc *DeployRecord) Save() bool {
	if dc.CreatedAt.IsZero() {
		dc.CreatedAt = JsonTime{time.Now()}
	}

	db := GetDBConn()
	defer db.Close()

	insertSql := `
		INSERT INTO deploy_records (status, server_id, compare, created_at, deploy_user)
		VALUES (:status, :server_id, :compare, :created_at, :deploy_user)
		RETURNING id
	`
	nstmt, err := db.PrepareNamed(insertSql)
	err = nstmt.Get(&dc.ID, dc)
	if err != nil {
		log.Printf("SaveDeployRecord failed: %v", err)
		return false
	}
	return true
}

func (dc *DeployRecord) UpdateStatus(newStatus string) bool {
	db := GetDBConn()
	defer db.Close()

	var ended_at JsonTime
	if newStatus != "processing" {
		ended_at = JsonTime{time.Now()}
	}

	updateSql := `
		UPDATE deploy_records
		SET status = $1, ended_at = $2
		WHERE id=$3;
	`
	_, err := db.Exec(updateSql, newStatus, ended_at, dc.ID)
	if err != nil {
		log.Printf("UpdateStatus failed: %v", err)
		return false
	}

	dc.Status = newStatus
	return true
}

func (dc *DeployRecord) Exec() {
	log.Printf("Exec DeployRecord #%d\n", dc.ID)
	dc.UpdateStatus("processing")
	dc.SetDeployServer()
	cmd := dc.DeployServer.buildCmd()

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	contextLogger := DeployLogger.WithFields(logrus.Fields{
		"DeployServer": dc.ServerID,
	})
	contextLogger.Info("Deploy starting")

	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if err != nil {
		contextLogger.Warn(err)
		dc.UpdateStatus("failed")
	} else {
		contextLogger.Info("Deploy done")
		dc.UpdateStatus("processed")
	}
	dc.storeCmdLog(outStr, errStr)
	dc.notify()
}

func FindDeployRecordByID(id int64) *DeployRecord {
	db := GetDBConn()
	defer db.Close()

	dr := DeployRecord{}
	queryErr := db.Get(&dr, "SELECT * FROM deploy_records WHERE id=$1", id)
	if queryErr != nil {
		log.Println(queryErr)
	}
	return &dr
}

func (dc *DeployRecord) storeCmdLog(outStr, errStr string) {
	db := GetDBConn()
	defer db.Close()

	updateSql := `
		UPDATE deploy_records
		SET stdout = $1, stderr = $2
		WHERE id=$3;
	`
	_, err := db.Exec(updateSql, outStr, errStr, dc.ID)
	if err != nil {
		log.Printf("storeCmdLog failed: %v", err)
	}
}

func (dr *DeployRecord) notify() {
	deployServer := dr.DeployServer
	deployServer.SetDeployRepo()
	deployRepo := deployServer.DeployRepo

	openidsArr := strings.Split(deployRepo.Openids, ";")
	if len(openidsArr) == 0 {
		return
	}

	options := make(map[string]string)
	if dr.Status == "processed" {
		options["first"] = "部署成功"
	} else {
		options["first"] = "部署失败"
	}

	options["keyword1"] = deployRepo.Name
	options["keyword2"] = deployServer.Name
	options["remark"] = "点击查看详情"

	if notification.NotifyBySkylark(openidsArr, options) {
		log.Printf("Succeed to notify admin of DeployServer#%d", dr.ServerID)
	}
}
