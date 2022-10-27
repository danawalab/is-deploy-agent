package utils

import (
	"encoding/json"
	"is-deploy-agent/model"
	"log"
	"os"
)

func GetJson() []model.Model {
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
