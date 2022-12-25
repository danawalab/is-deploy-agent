package main

import (
	"fmt"
	"is-deploy-agent/router"
	"os"
)

func main() {
	port := os.Args
	server := router.SetRouter()

	if len(port) == 1 {
		fmt.Println("Please specify port")
	} else if len(port) == 2 {
		err := server.Run(":" + port[1])
		if err != nil {
			fmt.Println("port error ", err)
		}
	} else {
		fmt.Println("There are multiple ports, Please write only one port.")
	}
}
