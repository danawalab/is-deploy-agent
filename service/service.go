package service

import (
	"encoding/json"
	"is-deploy-agent/model"
	"log"
	"os"
)

func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func writeFile(path string, lb string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	file.Write([]byte(lb))
	file.Close()
}

func readJson() string {
	path, err := os.Open("./setting.json")
	if err != nil {
		log.Fatal(err)
	}

	var model []model.Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&model)

	key := model[0].PodList[0].LbMap[0].Key
	value := model[0].PodList[0].LbMap[0].Value

	result := key + "=" + value

	return result
}

func Restore() {
	path := "../ex.properties"
	lb := readJson()
	readFile(path)
	writeFile(path, lb)
	readFile(path)
}

func Exclude() {

}
