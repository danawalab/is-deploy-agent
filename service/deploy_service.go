package service

import (
	"is-deploy-agent/utils"
	"log"
	"os/exec"
	"strings"
)

// Deploy
// setting.json에 지정된 쉘 스크립트 경로를 통해 쉘 스크립트 실행 하여 배포
func Deploy(worker string, arguments string) (string, error) {
	argument := strings.Split(arguments, " ")

	output, err := executeShell(worker, argument...)
	if err != nil {
		log.Println(err)
		return output, err
	}

	return output, err
}

// shell을 실행
func executeShell(worker string, arguments ...string) (string, error) {
	shellPath, err := getShellPath(worker)
	if err != nil {
		log.Println(err)
		return "", err
	}
	cmd := exec.Command(shellPath, arguments...)
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println("Execute ShellScript Script : ", string(output))
	return string(output), err
}

// shell의 경로를 반환
func getShellPath(worker string) (string, error) {
	node, err := utils.GetSettingJson()
	if err != nil {
		log.Println(err)
		return "", err
	}
	podLength := len(node.PodList)
	var shellPath string

	for pods := 0; pods < podLength; pods++ {
		pod := node.PodList[pods]
		name := pod.Name

		if utils.IsNameEqual(name, worker) {
			shellPath = pod.ShPath
			break
		}
	}

	return shellPath, err
}
