package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()

	router.POST("/deploy/:server", func(c *gin.Context) {
		server := c.Param("server")
		ref := c.PostForm("ref")
		branch, err := getBranch(ref)
		if err != nil {
			return
		}
		fmt.Printf("Deploying %s with %s", server, branch)

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
