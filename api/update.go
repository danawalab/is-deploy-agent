package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/domain"
	"is-deploy-agent/utils"
	"net/http"
)

func AgentUpdate(context *gin.Context) {
	version := context.Params.ByName("version")
	err := domain.AgentUpdate(version)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		// 에러가 없으면 에이전트 종료 후 삭제하는데 반환 값은 의미가 없는가 아닌가?
		context.JSON(http.StatusOK, gin.H{
			"message": "에이전트 업데이트 완료",
		})
	}
}

func GetAgentVersion(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": utils.Version,
	})
}
