package rest

import (
	"csh-api/cmd/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

func StartServer () {
	server := defineServerProperties()
	restPort := ":" + strconv.Itoa(config.GetConfig().Rest.Port)
	server.Run(restPort)
}

func defineServerProperties() *gin.Engine {
	server := gin.Default()
	server.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "PONG",
		})
	})
	return server
}