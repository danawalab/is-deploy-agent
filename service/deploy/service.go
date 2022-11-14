package deploy

import (
	"is-deploy-agent/utils"
	"log"
	"os/exec"
)

// Deploy
// setting.json에 지정된 shell 경로를 통해 shell 실행 하여 배포
func Deploy(worker string) {
	executeShell(worker)
}

// shell을 실행
func executeShell(worker string) {
	shellPath := getShellPath(worker)
	cmd := exec.Command(shellPath)
	output, _ := cmd.Output()

	log.Println("Execute Shell Script : ", string(output))
}

// shell의 경로를 반환
func getShellPath(worker string) string {
	json := utils.GetJson()
	podLength := len(json.Node.PodList)
	var shellPath string

	for pods := 0; pods < podLength; pods++ {
		pod := json.Node.PodList[pods]
		name := pod.Name

		if utils.IsNameEqual(name, worker) {
			shellPath = pod.ShPath
			break
		}
	}

	return shellPath
}
