package main

import (
	"is-deploy-agent/router"
)

func main() {
	server := router.SetRouter()
	server.Run(":5000")
}
