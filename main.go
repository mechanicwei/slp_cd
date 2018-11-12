package main

import (
	_ "slp_cd/config"

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
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"X-Pagination-Total"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	go consumeDeployQueue()

	authMiddleware := route.NewAuthMiddleware()
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/deploy/:server", route.CreateDeployRecord(DeployQueue))

	api := router.Group("/api")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/deploy_servers", route.GetDeployServers)

		api.POST("/deploy_servers", route.CreateDeployServer)
		api.PATCH("/deploy_servers/:id", route.UpdateDeployServer)
		api.GET("/deploy_servers/:server_id/records", route.GetDeployRecords)
	}

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
