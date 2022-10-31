package log

import (
	"bufio"
	"fmt"
	"is-deploy-agent/utils"
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

		if utils.IsNameEqual(name, worker) {
			logPath = pod.LogPath
			break
		}
	}

	logs, err := os.Open(logPath)
	if err != nil {
		fmt.Println(err)
		//todo log로 변경
	}
	scanner := bufio.NewScanner(logs)
	defer logs.Close()
	return scanner
}

func GetLogTailFlagN(worker string, line string) string {
	models := utils.GetJson()
	logLength := len(models[0].NodeList[0].PodList)

	var logPath string
	for pods := 0; pods < logLength; pods++ {
		pod := models[0].NodeList[0].PodList[pods]
		name := pod.Name

		if utils.IsNameEqual(name, worker) {
			logPath = pod.LogPath
			break
		}
	}

	cmd := exec.Command("tail", "-n", line, logPath)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		//todo log로 변경
	}
	return string(output)
}
