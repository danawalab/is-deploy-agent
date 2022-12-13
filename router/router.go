package router

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service/deploy"
	"is-deploy-agent/service/fetch"
	"is-deploy-agent/service/loadbalance"
	"is-deploy-agent/service/log"
	"is-deploy-agent/utils"
	"net/http"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	lb := router.Group("/load-balance")
	{
		lb.GET("", func(context *gin.Context) {
			lbStatus, err := loadbalance.CheckLbStatus()
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else if lbStatus == "Not Match" {
				context.JSON(http.StatusOK, gin.H{
					"error": lbStatus,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"message": lbStatus,
				})
			}
		})

		lb.PUT("/exclude", func(context *gin.Context) {
			worker := context.Query("worker")
			err := loadbalance.Exclude(worker)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"message": worker + " 가 제외되었습니다",
				})
			}
		})

		lb.PUT("/restore", func(context *gin.Context) {
			err := loadbalance.Restore()
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"message": "연결이 복원되었습니다",
				})
			}
		})
	}

	dp := router.Group("/webapp")
	{
		dp.PUT("/deploy", func(context *gin.Context) {
			worker := context.Query("worker")
			err := deploy.Deploy(worker)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"message": worker + " 가 배포되었습니다",
				})
			}
		})
	}

	sc := router.Group("/sync")
	{
		// deprecated
		// GET Method 사용 안하고 있음
		sc.GET("", func(context *gin.Context) {
			json, err := fetch.GetSettingJson()
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"data": json,
				})
			}
		})

		sc.PUT("", func(context *gin.Context) {
			body, _ := context.GetRawData()
			err := fetch.SyncSettingJson(string(body))
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"message": "setting.json 동기화 완료",
				})
			}
		})
	}

	up := router.Group("/update")
	{
		up.PUT("/:version", func(context *gin.Context) {
			version := context.Params.ByName("version")
			err := fetch.UpdateAgent(version)
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
		})

		up.GET("/version", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": utils.Version,
			})
		})
	}

	lg := router.Group("/logs")
	{
		lg.GET("/tail/n", func(context *gin.Context) {
			worker := context.Query("worker")
			line := context.Query("line")
			logs, err := log.GetLogTailFlagN(worker, line)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				context.String(http.StatusOK, logs)
			}
		})
	}

	hp := router.Group("/health-check")
	{
		hp.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "정상적으로 연결되었습니다",
			})
		})
	}

	return router
}
