package router

import (
	"github.com/gin-gonic/gin"
	"is-deploy-agent/domain/deploy"
	"is-deploy-agent/domain/healthCheck"
	"is-deploy-agent/domain/loadBalance"
	"is-deploy-agent/domain/log"
	"is-deploy-agent/domain/setting"
	"is-deploy-agent/domain/update"
	"is-deploy-agent/utils"
	"net/http"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	lb := router.Group("/api/v1/load-balance")
	{
		lb.GET("", func(context *gin.Context) {
			lbStatus, err := loadBalance.CheckLbStatus()
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
		})

		lb.PUT("/exclude", func(context *gin.Context) {
			worker := context.Query("worker")
			err := loadBalance.Exclude(worker)
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
			err := loadBalance.Restore()
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

	dp := router.Group("/api/v1/deploy")
	{
		dp.PUT("/shell", func(context *gin.Context) {
			worker := context.Query("worker")
			arguments := context.Query("arguments")

			output, err := deploy.Deploy(worker, arguments)

			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"message": output,
				})
			}
		})
	}

	set := router.Group("/api/v1/setting")
	{
		set.PUT("", func(context *gin.Context) {
			body, _ := context.GetRawData()
			err := setting.SyncSettingJson(string(body))
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

	up := router.Group("/api/v1/update")
	{
		up.PUT("/:version", func(context *gin.Context) {
			version := context.Params.ByName("version")
			err := update.AgentUpdate(version)
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

	lg := router.Group("/api/v1/logs")
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

	hc := router.Group("/api/v1/health-check")
	{
		hc.GET("/agent", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "정상적으로 연결되었습니다",
			})
		})

		hc.GET("/tomcat", func(context *gin.Context) {
			tomcat := context.Query("tomcat")

			output, err := healthCheck.TomcatHealthCheck(tomcat)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"message": output,
				})
			}
		})
	}

	return router
}
