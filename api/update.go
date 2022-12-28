package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service"
	"is-deploy-agent/utils"
	"net/http"
)

func AgentUpdate(context *gin.Context) {
	version := context.Params.ByName("version")
	err := service.AgentUpdate(version)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
}

func GetAgentVersion(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": utils.Version,
	})
}
