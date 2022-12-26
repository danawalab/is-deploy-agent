package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service"
	"net/http"
)

func SyncSettingJson(context *gin.Context) {
	body, _ := context.GetRawData()
	err := service.SyncSettingJson(string(body))
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "setting.json 동기화 완료",
		})
	}
}
