package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/domain"
	"net/http"
)

func SyncSettingJson(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}

	err = domain.SyncSettingJson(string(body))
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
