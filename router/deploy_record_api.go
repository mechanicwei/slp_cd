package route

import (
	"errors"
	"fmt"
	"slp_cd/model"
	"strconv"
	"strings"

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
			Status:   "waiting",
			ServerID: deployServer.ID,
			Commit:   c.PostForm("head_commit"),
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

func GetDeployRecords(c *gin.Context) {
	deployServerId, _ := strconv.ParseInt(c.Param("server_id"), 10, 64)
	deployServer := model.FindDeployServerByID(deployServerId)

	if deployServer.ID == 0 {
		message := fmt.Sprintf("Can't find DeployServer with id %d", deployServerId)
		c.JSON(404, gin.H{
			"error": message,
		})
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("per_page"))

	deployRecords := deployServer.PaginatedDeployRecords(page, perPage)

	c.JSON(200, deployRecords)
}

func getBranch(ref string) (string, error) {
	s := strings.Split(ref, "/")
	if s[1] == "head" {
		return s[2], nil
	} else {
		return "", errors.New("no branch")
	}
}
