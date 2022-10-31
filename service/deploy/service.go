package deploy

import (
	"fmt"
	"is-deploy-agent/utils"
	"os/exec"
)

func Deploy(worker string) {
	executeShell(worker)
}

func executeShell(worker string) {
	shellPath := getShellPath(worker)
	cmd := exec.Command(shellPath)
	output, _ := cmd.Output()

	fmt.Println(string(output))
	//todo log로 변경
}

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
