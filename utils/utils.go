package utils

import (
	"encoding/json"
	"fmt"
	"is-deploy-agent/model"
	"os"
)

// GetJson
// setting.json을 읽어서 반환
func GetJson() model.Model {
	path, err := os.Open("./setting.json")
	if err != nil {
		fmt.Println(err)
	}

	var models model.Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&models)

	defer path.Close()
	return models
}

// IsNameEqual
// setting.json에서 지정한 node의 name 또는 podList에 name이 worker 인자와 같은 경우 true 반환
func IsNameEqual(name string, worker string) bool {
	if name == worker {
		return true
	}
	return false
}
