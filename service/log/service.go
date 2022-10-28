package log

import (
	"bufio"
	"github.com/hpcloud/tail"
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

func GetLogTailFlagN(worker string, line string) string {
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
	cmd := exec.Command("tail", "-n", line, logPath)
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	return string(output)
}

func GetLogTailFlagF(worker string) *tail.Tail {
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

	t, _ := tail.TailFile(logPath, tail.Config{Follow: true, ReOpen: true, MustExist: true, Poll: true, Location: &tail.SeekInfo{0, 2}})

	return t
}

func isNameEqual(name string, worker string) bool {
	if name == worker {
		return true
	}
	return false
}
