package route

import (
	"errors"
	"fmt"
	"slp_cd/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateDeployRecord(DeployQueue chan int64) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

func getBranch(ref string) (string, error) {
	s := strings.Split(ref, "/")
	if s[1] == "head" {
		return s[2], nil
	} else {
		return "", errors.New("no branch")
	}
}
