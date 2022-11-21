package utils

import (
	"encoding/json"
	"is-deploy-agent/model"
	"log"
	"os"
)

// GetJson
// setting.json을 읽어서 반환
func GetJson() (model.Node, error) {
	path, err := os.Open("./setting.json")
	if err != nil {
		log.Println(err)
	}
	defer path.Close()

	var models model.Node
	err = json.NewDecoder(path).Decode(&models)
	if err != nil {
		log.Println(err)
	}

	return models, err
}

// IsNameEqual
// setting.json에서 지정한 node의 name 또는 podList에 name이 worker 인자와 같은 경우 true 반환
func IsNameEqual(name string, worker string) bool {
	if name == worker {
		return true
	}
	return false
}
