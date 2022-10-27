package log

import (
	"bufio"
	"is-deploy-agent/utils"
	"log"
	"os"
)

func GetLogAll() *bufio.Scanner {
	models := utils.GetJson()

	logPath := models[0].NodeList[0].PodList[0].LogPath

	logs, err := os.Open(logPath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(logs)

	return scanner
}
