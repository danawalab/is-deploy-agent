package utils

import (
	"encoding/json"
	"fmt"
	"is-deploy-agent/model"
	"os"
)

func GetJson() []model.Model {
	path, err := os.Open("./setting.json")
	if err != nil {
		fmt.Println(err)
	}

	var models []model.Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&models)

	defer path.Close()
	return models
}

func IsNameEqual(name string, worker string) bool {
	if name == worker {
		return true
	}
	return false
}
