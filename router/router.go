package router

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service/deploy"
	"is-deploy-agent/service/loadbalance"
	"is-deploy-agent/service/sync"
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
			service := context.Query("service")
			worker := context.Query("worker")
			deploy.Deploy(0, worker)
			context.String(http.StatusOK, "Router deploy Ready %s %s", service, worker)
		})
	}

	sc := router.Group("/sync")
	{
		sc.PUT("", sync.FetchJson)
	}

	lg := router.Group("/logs")
	{
		lg.GET("")
	}

	hp := router.Group("/health-check")
	{
		hp.GET("", func(context *gin.Context) {
			context.String(http.StatusOK, "Health Good")
		})
	}

	return router
}
