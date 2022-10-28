package log

import (
	"bufio"
	"is-deploy-agent/utils"
	"log"
	"os"
	"os/exec"
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

func GetLogTailFlagN(worker string) string {
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
	cmd := exec.Command("tail", "-n 10", logPath)
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	return string(output)
}

func isNameEqual(name string, worker string) bool {
	if name == worker {
		return true
	}
	return false
}
