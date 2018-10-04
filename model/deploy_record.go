package model

import (
	"log"
	"time"
)

// Valid status is in ["waiting", "processing", "processed", "failed"]
type DeployRecord struct {
	ID        int64     `json:"id"`
	Status    string    `json:"status"`
	ServerID  int64     `json:"server_id"`
	Commit    string    `json:"commit"`
	CreatedAt time.Time `json:"created_at"`
}

func (dc *DeployRecord) Save() bool {
	db := GetDBConn()
	insertSql := "INSERT INTO deploy_records (status, server_id, commit, created_at) VALUES ($1, $2, $3, $4)"
	result, err := db.Exec(insertSql, dc.Status, dc.ServerID, dc.Commit, dc.CreatedAt.Format(time.RFC3339))
	if err != nil {
		log.Printf("SaveDeployRecord failed: %v", err)
		return false
	}

	dc.ID, _ = result.LastInsertId()
	return true
}
