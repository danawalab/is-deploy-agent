package utils

import (
	"encoding/json"
	"is-deploy-agent/model"
	"log"
	"os"
)

func GetJsonToTest() []model.Model {
	path, err := os.Open("../../setting.json")
	if err != nil {
		log.Fatal(err)
	}

	var model []model.Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&model)

	return model
}
