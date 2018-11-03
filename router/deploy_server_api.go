package route

import (
	"fmt"
	"net/http"
	"slp_cd/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateDeployServer(c *gin.Context) {
	var deployServer model.DeployServer

	if err := c.ShouldBindJSON(&deployServer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if deployServer.Save() {
		c.JSON(201, deployServer)
	} else {
		c.JSON(422, gin.H{
			"status": "failed",
		})
	}
}

func UpdateDeployServer(c *gin.Context) {
	deployServerId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	deployServer := model.FindDeployServerByID(deployServerId)

	if deployServer.ID == 0 {
		message := fmt.Sprintf("Can't find DeployServer with id %d", deployServerId)
		c.JSON(404, gin.H{
			"error": message,
		})
		return
	}

	if err := c.ShouldBindJSON(&deployServer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if deployServer.Update() {
		c.JSON(200, deployServer)
	} else {
		c.JSON(422, gin.H{
			"status": "failed",
		})
	}
}

func GetDeployServers(c *gin.Context) {
	deployServers := model.AllDeployServers()

	c.JSON(200, deployServers)
}
