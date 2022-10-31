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
			loadbalance.Exclude(0, worker)
			context.String(http.StatusOK, "Router exclude Ready %s", worker)
		})
		lb.PUT("/restore", func(context *gin.Context) {
			loadbalance.Restore(0)
			context.String(http.StatusOK, "Router restore Ready")
		})
	}

	dp := router.Group("/webapp")
	{
		dp.PUT("/deploy", func(context *gin.Context) {
			worker := context.Query("worker")
			deploy.Deploy(0, worker)
			context.String(http.StatusOK, "Router deploy Ready %s", worker)
		})
	}

	sc := router.Group("/sync")
	{
		sc.PUT("", fetch.FetchJson)
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
			context.String(http.StatusOK, "Health Good")
		})
	}

	return router
}
