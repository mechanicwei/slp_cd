package route

import (
	"fmt"
	"net/http"
	"slp_cd/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateDeployRepo(c *gin.Context) {
	var deployRepo model.DeployRepo

	if err := c.ShouldBindJSON(&deployRepo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if deployRepo.Save() {
		c.JSON(201, deployRepo)
	} else {
		c.JSON(422, gin.H{
			"status": "failed",
		})
	}
}

func UpdateDeployRepo(c *gin.Context) {
	deployRepoId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	deployRepo := model.FindDeployRepoByID(deployRepoId)

	if deployRepo.ID == 0 {
		message := fmt.Sprintf("Can't find DeployRepo with id %d", deployRepoId)
		c.JSON(404, gin.H{
			"error": message,
		})
		return
	}

	if err := c.ShouldBindJSON(&deployRepo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if deployRepo.Update() {
		c.JSON(200, deployRepo)
	} else {
		c.JSON(422, gin.H{
			"status": "failed",
		})
	}
}

func GetDeployRepos(c *gin.Context) {
	deployRepos := model.AllDeployRepos()

	c.JSON(200, deployRepos)
}
