package router

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service/deploy"
	"is-deploy-agent/service/fetch"
	"is-deploy-agent/service/loadbalance"
	"is-deploy-agent/service/log"
	"net/http"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	lb := router.Group("/load-balance")
	{
		lb.PUT("/exclude", func(context *gin.Context) {
			worker := context.Query("worker")
			loadbalance.Exclude(worker)
			context.JSON(http.StatusOK, gin.H{
				"message": worker + "is Exclude Complete",
			})
		})
		lb.PUT("/restore", func(context *gin.Context) {
			loadbalance.Restore()
			context.JSON(http.StatusOK, gin.H{
				"message": "Restore Complete",
			})
		})
	}

	dp := router.Group("/webapp")
	{
		dp.PUT("/deploy", func(context *gin.Context) {
			worker := context.Query("worker")
			deploy.Deploy(worker)
			context.JSON(http.StatusOK, gin.H{
				"message": "Deploy Complete",
			})
		})
	}

	sc := router.Group("/sync")
	{
		sc.GET("", func(context *gin.Context) {
			json := fetch.GetSettingJson()
			context.JSON(http.StatusOK, gin.H{
				"data": json,
			})
		})

		sc.PUT("", func(context *gin.Context) {
			fetch.SyncSettingJson()
			context.JSON(http.StatusOK, gin.H{
				"message": "setting.json sync complete",
			})
		})
	}

	lg := router.Group("/logs")
	{
		lg.GET("/all", func(context *gin.Context) {
			worker := context.Query("worker")
			logs := log.GetLogAll(worker)
			for logs.Scan() {
				context.String(http.StatusOK, "%s\n", logs.Text())
			}
		})

		lg.GET("/tail/n", func(context *gin.Context) {
			worker := context.Query("worker")
			line := context.Query("line")
			logs := log.GetLogTailFlagN(worker, line)
			context.String(http.StatusOK, logs)
		})
	}

	hp := router.Group("/health-check")
	{
		hp.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Health Good",
			})
		})
	}

	return router
}
