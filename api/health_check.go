package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/domain"
	"net/http"
)

func AgentHealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "정상적으로 연결되었습니다",
	})
}

func TomcatHealthCheck(context *gin.Context) {
	tomcat := context.Query("tomcat")

	output, err := domain.HealthCheck(tomcat)
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
