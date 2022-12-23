package tests

import (
	"encoding/json"
	"is-deploy-agent/domain"
	"log"
	"os"
)

func GetJsonToTest() domain.Node {
	path, err := os.Open("../../setting.json")
	if err != nil {
		log.Println(err)
	}

	var models domain.Node

	decoder := json.NewDecoder(path)
	err = decoder.Decode(&models)
	if err != nil {
		log.Println(err)
	}

	defer path.Close()
	return models
}
