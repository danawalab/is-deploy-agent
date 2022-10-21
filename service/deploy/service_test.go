package deploy

import (
	"bufio"
	"encoding/json"
	"fmt"
	"is-deploy-agent/model"
	"log"
	"os"
	"testing"
)

func TestGetLog(t *testing.T) {
	models := getJson()
	logPath := models[0].NodeList[0].PodList[0].LogPath
	logName := "catalina.2022-10-21.log"

	logs, err := os.Open(logPath + logName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(logs)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}

func getJson() []model.Model {
	path, err := os.Open("../../setting.json")
	if err != nil {
		log.Fatal(err)
	}

	var model []model.Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&model)

	return model
}
