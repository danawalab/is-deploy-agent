package log

import (
	"bufio"
	"is-deploy-agent/utils"
	"log"
	"os"
)

func GetLogAll(worker string) *bufio.Scanner {
	models := utils.GetJson()
	logLength := len(models[0].NodeList[0].PodList)

	var logPath string
	for pods := 0; pods < logLength; pods++ {
		pod := models[0].NodeList[0].PodList[pods]
		name := pod.Name

		if isNameEqual(name, worker) {
			logPath = pod.LogPath

			break
		}
	}

	logs, err := os.Open(logPath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(logs)

	return scanner
}

func isNameEqual(name string, worker string) bool {
	if name == worker {
		return true
	}
	return false
}
