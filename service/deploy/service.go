package deploy

import (
	"is-deploy-agent/utils"
	"log"
	"os/exec"
)

// Deploy
// setting.json에 지정된 shell 경로를 통해 shell 실행 하여 배포
func Deploy(worker string) error {
	err := executeShell(worker)
	if err != nil {
		log.Println(err)
	}

	return err
}

// shell을 실행
func executeShell(worker string) error {
	shellPath, err := getShellPath(worker)
	if err != nil {
		log.Println(err)
	}
	cmd := exec.Command(shellPath)
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}

	log.Println("Execute Shell Script : ", string(output))
	return err
}

// shell의 경로를 반환
func getShellPath(worker string) (string, error) {
	json, err := utils.GetJson()
	if err != nil {
		log.Println(err)
	}
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

	return shellPath, err
}
