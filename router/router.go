package router

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/service"
	"net/http"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	lb := router.Group("/loadbalance")
	{
		lb.PUT("/exclude", func(context *gin.Context) {
			worker := context.Query("worker")

			context.String(http.StatusOK, "Router exclude Ready %s", worker)
		})
		lb.PUT("/restore", func(context *gin.Context) {
			service.Restore(0)
			context.String(http.StatusOK, "Router restore Ready")
		})
	}

	return router
}
