package main

import (
	//	"github.com/gin-gonic/gin"
	"is-deploy-agent/router"
)

func main() {

	server := router.SetRouter()

	server.Run()
}
