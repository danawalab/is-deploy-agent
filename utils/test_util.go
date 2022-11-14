package utils

import (
	"encoding/json"
	"is-deploy-agent/model"
	"log"
	"os"
)

func GetJsonToTest() model.Model {
	path, err := os.Open("../../setting.json")
	if err != nil {
		log.Println(err)
	}

	var models model.Model

	decoder := json.NewDecoder(path)
	err = decoder.Decode(&models)
	if err != nil {
		log.Println(err)
	}

	defer path.Close()
	return models
}
