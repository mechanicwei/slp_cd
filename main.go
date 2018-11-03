package main

import (
	"slp_cd/model"
	route "slp_cd/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var DeployQueue = make(chan int64)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PATCH", "POST"},
		AllowHeaders:     []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	go consumeDeployQueue()

	router.GET("/deploy_servers", route.GetDeployServers)
	router.POST("/deploy/:server", route.CreateDeployRecord(DeployQueue))

	router.POST("/deploy_servers", route.CreateDeployServer)
	router.PATCH("/deploy_servers/:id", route.UpdateDeployServer)
	router.GET("/deploy_servers/:server_id/records", route.GetDeployRecords)
	router.Run(":8080")
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
