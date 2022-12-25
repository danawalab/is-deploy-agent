package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/domain"
	"net/http"
)

func ShellDeploy(context *gin.Context) {
	tomcat := context.Query("tomcat")
	arguments := context.Query("arguments")

	output, err := domain.Deploy(tomcat, arguments)

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
