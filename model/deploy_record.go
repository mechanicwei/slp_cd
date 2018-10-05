package model

import (
	"log"
	"time"
)

// Valid status is in ["waiting", "processing", "processed", "failed"]
type DeployRecord struct {
	ID        int64     `json:"id"`
	Status    string    `json:"status"`
	ServerID  int64     `json:"server_id" db:"server_id"`
	Commit    string    `json:"commit"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (dc *DeployRecord) DeployServer() *DeployServer {
	return FindDeployServerByID(dc.ServerID)
}

func (dc *DeployRecord) Save() bool {
	db := GetDBConn()
	insertSql := "INSERT INTO deploy_records (status, server_id, commit, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err := db.QueryRow(insertSql, dc.Status, dc.ServerID, dc.Commit, dc.CreatedAt.Format(time.RFC3339)).Scan(&dc.ID)
	if err != nil {
		log.Printf("SaveDeployRecord failed: %v", err)
		return false
	}
	return true
}

func (dc *DeployRecord) UpdateStatus(newStatus string) bool {
	db := GetDBConn()
	updateSql := `
		UPDATE deploy_records
		SET status = $1
		WHERE id=$2;
	`
	_, err := db.Exec(updateSql, newStatus, dc.ID)
	if err != nil {
		log.Printf("UpdateStatus failed: %v", err)
		return false
	}

	return true
}

func (dc *DeployRecord) Exec() {
	log.Printf("Exec DeployRecord #%d\n", dc.ID)
	dc.UpdateStatus("processing")

	if dc.DeployServer().runCmd() {
		dc.UpdateStatus("processed")
	} else {
		dc.UpdateStatus("failed")
	}
}

func FindDeployRecordByID(id int64) *DeployRecord {
	db := GetDBConn()
	dr := DeployRecord{}
	queryErr := db.Get(&dr, "SELECT * FROM deploy_records WHERE id=$1", id)
	if queryErr != nil {
		log.Println(queryErr)
	}
	return &dr
}
