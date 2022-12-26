package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service"
	"net/http"
)

func CheckLbStatus(context *gin.Context) {
	lbStatus, err := service.CheckLbStatus()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else if lbStatus == "매칭되는 거 없음" {
		context.JSON(http.StatusOK, gin.H{
			"error": lbStatus,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": lbStatus,
		})
	}
}

func Exclude(context *gin.Context) {
	worker := context.Query("worker")
	err := service.Exclude(worker)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": worker + " 제외되었습니다",
		})
	}
}

func Restore(context *gin.Context) {
	err := service.Restore()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "연결이 복원되었습니다",
		})
	}
}
