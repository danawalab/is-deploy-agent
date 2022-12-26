package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"is-deploy-agent/router"
	"os"
)

func main() {
	port := os.Args
	server := router.SetRouter()

	gin.SetMode(gin.ReleaseMode)

	if len(port) == 1 {
		fmt.Println("port를 지정해 주세요")
	} else if len(port) == 2 {
		err := server.Run(":" + port[1])
		if err != nil {
			fmt.Println("port 에러 ", err)
		}
	} else {
		fmt.Println("port가 여러개 입니다")
	}
}
