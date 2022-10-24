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

	lb := router.Group("/loadbalance")
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

	fetch := router.Group("/fetch")
	{
		fetch.PUT("", sync.FetchJson)
	}

	return router
}
