package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service"
	"net/http"
)

func AgentHealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "정상적으로 연결되었습니다",
	})
}

func TomcatHealthCheck(context *gin.Context) {
	worker := context.Query("worker")

	output, err := service.TomcatHealthCheck(worker)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": output,
		})
	}
}
