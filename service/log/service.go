package log

import (
	"is-deploy-agent/utils"
	"log"
	"os/exec"
)

// GetLogTailFlagN
// tail -n을 사용하여 로그를 끝에서 부터 N번째 줄까지 반환
func GetLogTailFlagN(worker string, line string) (string, error) {
	node, err := utils.GetJson()
	if err != nil {
		log.Println(err)
		return "", err
	}
	logLength := len(node.PodList)

	var logPath string
	for pods := 0; pods < logLength; pods++ {
		pod := node.PodList[pods]
		name := pod.Name

		if utils.IsNameEqual(name, worker) {
			logPath = pod.LogPath
			break
		}
	}

	cmd := exec.Command("tail", "-n", line, logPath)
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(output), err
}
