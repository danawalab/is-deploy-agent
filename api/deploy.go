package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service"
	"net/http"
)

func Deploy(context *gin.Context) {
	worker := context.Query("worker")
	arguments := context.Query("arguments")

	output, err := service.Deploy(worker, arguments)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "배포 쉘 스크립트가 실행 되었습니다.",
			"output":  output,
		})
	}
}
