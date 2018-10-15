package route

import (
	"net/http"
	"slp_cd/model"

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
		c.JSON(201, gin.H{
			"status": "created",
		})
	} else {
		c.JSON(422, gin.H{
			"status": "failed",
		})
	}
}
