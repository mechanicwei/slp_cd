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
		repoID, _ := strconv.ParseInt(c.Query("repo_id"), 10, 64)
		ref := c.PostForm("ref")
		branch, err := getBranch(ref)
		if err != nil {
			return
		}

		deployServer := model.FindServerByRepoIDAndBranch(repoID, branch)

		if deployServer.ID == 0 {
			fmt.Printf("Can't find a server for #%d with %s branch\n", repoID, branch)
			return
		}
		fmt.Printf("Deploying #%d with %s", repoID, branch)

		var deployUser model.DeployUser
		sender := c.PostFormMap("sender")
		deployUser.Name = sender["login"]
		deployUser.AvatarUrl = sender["avatar_url"]
		deployUser.GithubUrl = sender["html_url"]

		deployRecord := model.DeployRecord{
			Status:     "waiting",
			ServerID:   deployServer.ID,
			Compare:    c.PostForm("compare"),
			DeployUser: deployUser,
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

	totalCount := strconv.Itoa(deployServer.TotalDeployRecordsCount())
	c.Header("X-Pagination-Total", totalCount)

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
