package model

import (
	"fmt"
	"log"
	"time"
)

type DeployRepo struct {
	ID            int64    `json:"id"`
	Name          string   `json:"name"`
	GithubUrl     string   `json:"github_url" db:"github_url"`
	WebhookSecret string   `json:"webhook_secret" db:"webhook_secret"`
	Openids       string   `json:"openids"`
	CreatedAt     JsonTime `json:"created_at" db:"created_at"`
}

func (dr *DeployRepo) Save() bool {
	db := GetDBConn()
	defer db.Close()
	if dr.CreatedAt.IsZero() {
		dr.CreatedAt = JsonTime{time.Now()}
	}
	insertSql := `
		INSERT INTO deploy_repos (name, github_url, webhook_secret, openids, created_at)
		VALUES (:name, :github_url, :webhook_secret, :openids, :created_at)
		RETURNING id
	`
	fmt.Println(dr)
	nstmt, _ := db.PrepareNamed(insertSql)
	err := nstmt.Get(&dr.ID, dr)
	if err != nil {
		log.Printf("SaveDeployRepo failed: %v", err)
		return false
	}
	return true
}

func FindDeployRepoByID(id int64) *DeployRepo {
	db := GetDBConn()
	defer db.Close()
	dr := DeployRepo{}
	queryErr := db.Get(&dr, "SELECT * FROM deploy_repos WHERE id=$1", id)
	if queryErr != nil {
		log.Println(queryErr)
	}
	return &dr
}

func (ds *DeployRepo) Update() bool {
	db := GetDBConn()
	defer db.Close()

	updateSql := `
		UPDATE deploy_repos
		SET name=:name, github_url=:github_url, webhook_secret=:webhook_secret, openids=:openids
		WHERE id = :id
	`
	_, err := db.NamedExec(updateSql, ds)
	if err != nil {
		log.Printf("updateDeployRepo failed: %v", err)
		return false
	}
	return true
}

func AllDeployRepos() []DeployRepo {
	selectSql := `SELECT * FROM deploy_repos ORDER BY id asc`
	deployRepos := []DeployRepo{}
	db := GetDBConn()
	defer db.Close()
	err := db.Select(&deployRepos, selectSql)

	if err != nil {
		fmt.Println(err)
	}
	return deployRepos
}
