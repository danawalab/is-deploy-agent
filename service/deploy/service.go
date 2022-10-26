package deploy

import (
	"encoding/json"
	"fmt"
	"is-deploy-agent/model"
	"log"
	"os"
	"os/exec"
)

func Deploy(node int, worker string) {
	executeShell(node, worker)
}

func executeShell(node int, worker string) {
	shellPath := getShellPath(node, worker)
	cmd := exec.Command(shellPath)
	output, _ := cmd.Output()

	fmt.Println(string(output))
}

func getShellPath(node int, worker string) string {
	models := readJson()
	podLength := len(models[0].NodeList[node].PodList)

	var shellPath string
	for pods := 0; pods < podLength; pods++ {
		pod := models[0].NodeList[node].PodList[pods]
		name := pod.Name

		if isNameEqual(name, worker) {
			shellPath = pod.ShPath

			break
		}
	}

	return shellPath
}

func isNameEqual(name string, worker string) bool {
	if name == worker {
		return true
	}
	return false
}

func readJson() []model.Model {
	path, err := os.Open("./setting.json")
	if err != nil {
		log.Fatal(err)
	}

	var models []model.Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&models)

	defer path.Close()
	return models
}
