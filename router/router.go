package router

import (
	"github.com/gin-gonic/gin"
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
			context.String(http.StatusOK, "Router restore Ready")
		})
	}

	return router
}