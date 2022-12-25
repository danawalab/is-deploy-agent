package router

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/api"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	lb := router.Group("/api/v1/load-balance")
	{
		lb.GET("", api.CheckLbStatus)
		lb.PUT("/exclude", api.Exclude)
		lb.PUT("/restore", api.Restore)
	}

	dp := router.Group("/api/v1/deploy")
	{
		dp.PUT("/shell", api.ShellDeploy)
	}

	set := router.Group("/api/v1/setting")
	{
		set.PUT("", api.SyncSettingJson)
	}

	up := router.Group("/api/v1/update")
	{
		up.PUT("/:version", api.AgentUpdate)
		up.GET("", api.GetAgentVersion)
	}

	lg := router.Group("/api/v1/logs")
	{
		lg.GET("/tail/n", api.GetTomcatLog)
	}

	hc := router.Group("/api/v1/health-check")
	{
		hc.GET("/agent", api.AgentHealthCheck)
		hc.GET("/tomcat", api.TomcatHealthCheck)
	}

	return router
}
