package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/domain"
	"net/http"
)

func GetTomcatLog(context *gin.Context) {
	tomcat := context.Query("tomcat")
	line := context.Query("line")
	logs, err := domain.GetTomcatLog(tomcat, line)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		context.String(http.StatusOK, logs)
	}
}
