package router

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service/deploy"
	"is-deploy-agent/service/loadbalance"
	"net/http"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	lb := router.Group("/loadbalance")
	{
		lb.PUT("/exclude", func(context *gin.Context) {
			worker := context.Query("worker")
			loadbalance.Exclude(worker)
			context.String(http.StatusOK, "Router exclude Ready %s", worker)
		})
		lb.PUT("/restore", func(context *gin.Context) {
			loadbalance.Restore()
			context.String(http.StatusOK, "Router restore Ready")
		})
	}

	dp := router.Group("/webapp")
	{
		dp.PUT("/deploy", func(context *gin.Context) {
			service := context.Query("service")
			worker := context.Query("worker")
			deploy.Deploy(worker)
			context.String(http.StatusOK, "Router deploy Ready %s %s", service, worker)
		})
	}

	return router
}
