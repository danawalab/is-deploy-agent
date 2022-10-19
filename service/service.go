package service

import (
	"encoding/json"
	"is-deploy-agent/model"
	"log"
	"os"
)

func Restore(node int) {
	path := getPropertiesPath(0)
	loadbalancerMap := getLbMap(node)
	lbLength := len(loadbalancerMap)

	if isLengthOne(lbLength) {
		key := loadbalancerMap[0].Key
		value := loadbalancerMap[0].Value
		lb := key + "=" + value

		writeFileString(path, lb)
	} else {
		writeFileArray(path, loadbalancerMap, lbLength)
	}
}

func Exclude(node int, pod int) {
	//path := getPropertiesPath(0)
}

func writeFileString(path string, lb string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	file.Write([]byte(lb))
	file.Close()
}

func writeFileArray(path string, lb []model.WorkerMap, length int) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < length; i++ {
		key := lb[i].Key
		value := lb[i].Value

		lb := key + "=" + value + "\n"
		file.Write([]byte(lb))
	}
	file.Close()
}

func getLbMap(node int) []model.WorkerMap {
	jsons := readJson()
	jsonLength := len(jsons[0].NodeList[node].LbMap)
	var loadbalancerMap []model.WorkerMap

	if isLengthOne(jsonLength) {
		key := jsons[0].NodeList[node].LbMap[0].Key
		value := jsons[0].NodeList[node].LbMap[0].Value

		loadbalancerMap = append(loadbalancerMap, model.WorkerMap{Key: key, Value: value})
		return loadbalancerMap
	} else {
		for i := 0; i < jsonLength; i++ {
			key := jsons[0].NodeList[node].LbMap[i].Key
			value := jsons[0].NodeList[node].LbMap[i].Value

			loadbalancerMap = append(loadbalancerMap, model.WorkerMap{Key: key, Value: value})
		}
		return loadbalancerMap
	}
}

func isLengthOne(jsonLength int) bool {
	if jsonLength == 1 {
		return true
	}
	return false
}

func getPropertiesPath(node int) string {
	jsons := readJson()
	return jsons[0].NodeList[node].Path
}

func readJson() []model.Model {
	path, err := os.Open("./setting.json")
	if err != nil {
		log.Fatal(err)
	}

	var models []model.Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&models)

	return models
}
