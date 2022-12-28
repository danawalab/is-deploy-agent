package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service"
	"net/http"
)

func GetLog(context *gin.Context) {
	worker := context.Query("worker")
	line := context.Query("line")
	logs, err := service.GetLogTailFlagN(worker, line)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		context.String(http.StatusOK, logs)
	}
}
