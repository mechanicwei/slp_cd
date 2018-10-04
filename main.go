package main

import (
	"errors"
	"fmt"
	"slp_cd/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var DeployQueue = make(chan int64)

func main() {
	router := gin.Default()

	go consumeDeployQueue()

	router.POST("/deploy/:server", func(c *gin.Context) {
		server := c.Param("server")
		ref := c.PostForm("ref")
		branch, err := getBranch(ref)
		if err != nil {
			return
		}

		deployServer := model.FindServerByNameAndBranch(server, branch)

		if deployServer.ID == 0 {
			fmt.Printf("Can't find a server for %s with %s branch\n", server, branch)
			return
		}
		fmt.Printf("Deploying %s with %s", server, branch)

		deployRecord := model.DeployRecord{
			Status:    "waiting",
			ServerID:  deployServer.ID,
			Commit:    c.PostForm("head_commit"),
			CreatedAt: time.Now(),
		}

		if !deployRecord.Save() {
			return
		}

		DeployQueue <- deployRecord.ID

		c.JSON(200, gin.H{
			"status": "received",
		})
	})
	router.Run(":8080")
}

func getBranch(ref string) (string, error) {
	s := strings.Split(ref, "/")
	if s[1] == "head" {
		return s[2], nil
	} else {
		return "", errors.New("no branch")
	}
}

func consumeDeployQueue() {
	var deployRecordID int64
	for {
		select {
		case deployRecordID = <-DeployQueue:
			deployRecord := model.FindDeployRecordByID(deployRecordID)
			deployRecord.Exec()
		}
	}
}
