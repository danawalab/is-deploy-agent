package log

import (
	"bufio"
	"fmt"
	"is-deploy-agent/utils"
	"os"
	"os/exec"
)

// GetLogAll
// 모든 로그를 읽어들이고 나서 반환
// Deprecated
func GetLogAll(worker string) *bufio.Scanner {
	json := utils.GetJson()
	logLength := len(json.Node.PodList)

	var logPath string
	for pods := 0; pods < logLength; pods++ {
		pod := json.Node.PodList[pods]
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

// GetLogTailFlagN
// tail -n을 사용하여 로그를 끝에서 부터 N번째 줄까지 반환
func GetLogTailFlagN(worker string, line string) string {
	json := utils.GetJson()
	logLength := len(json.Node.PodList)

	var logPath string
	for pods := 0; pods < logLength; pods++ {
		pod := json.Node.PodList[pods]
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
