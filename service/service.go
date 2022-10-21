package service

import (
	"encoding/json"
	"is-deploy-agent/model"
	"log"
	"os"
)

func Restore() {
	path := getPropertiesPath(0)
	loadbalancerMap := getLoadbalancerMap(0)
	lbLength := len(loadbalancerMap)

	if isLengthOne(lbLength) {
		lb := getWorkerMapResult(loadbalancerMap)
		writeFileString(path, lb)
	} else {
		writeFileArray(path, loadbalancerMap, lbLength)
	}
}

func Exclude(worker string) {
	path := getPropertiesPath(0)
	excludeMap := getExcludeMap(0, worker)
	exLength := len(excludeMap)

	if isLengthOne(exLength) {
		ex := getWorkerMapResult(excludeMap)
		writeFileString(path, ex)
	} else {
		writeFileArray(path, excludeMap, exLength)
	}
}

func getWorkerMapResult(workerMap []model.WorkerMap) string {
	key := workerMap[0].Key
	value := workerMap[0].Value
	result := key + "=" + value

	return result
}

func getExcludeMap(node int, worker string) []model.WorkerMap {
	models := readJson()
	podLength := len(models[0].NodeList[node].PodList)
	var excludeMap []model.WorkerMap

	for pods := 0; pods < podLength; pods++ {
		pod := models[0].NodeList[node].PodList[pods]
		name := pod.Name

		if name == worker {
			exLength := len(pod.ExcludeMap)

			for excludeMaps := 0; excludeMaps < exLength; excludeMaps++ {
				key := pod.ExcludeMap[excludeMaps].Key
				value := pod.ExcludeMap[excludeMaps].Value

				excludeMap = append(excludeMap, model.WorkerMap{Key: key, Value: value})
			}
			break
		}
	}
	return excludeMap
}

func getLoadbalancerMap(node int) []model.WorkerMap {
	models := readJson()
	modelLength := len(models[0].NodeList[node].LbMap)
	var loadbalancerMap []model.WorkerMap

	if isLengthOne(modelLength) {
		key := models[0].NodeList[node].LbMap[0].Key
		value := models[0].NodeList[node].LbMap[0].Value

		loadbalancerMap = append(loadbalancerMap, model.WorkerMap{Key: key, Value: value})
		return loadbalancerMap
	} else {
		for loadbalancer := 0; loadbalancer < modelLength; loadbalancer++ {
			key := models[0].NodeList[node].LbMap[loadbalancer].Key
			value := models[0].NodeList[node].LbMap[loadbalancer].Value

			loadbalancerMap = append(loadbalancerMap, model.WorkerMap{Key: key, Value: value})
		}
		return loadbalancerMap
	}
}

func writeFileString(path string, workerMap string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	file.Write([]byte(workerMap))
	defer file.Close()
}

func writeFileArray(path string, workerMaps []model.WorkerMap, length int) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	for workerMap := 0; workerMap < length; workerMap++ {
		key := workerMaps[workerMap].Key
		value := workerMaps[workerMap].Value

		lb := key + "=" + value + "\n"
		file.Write([]byte(lb))
	}
	defer file.Close()
}

func isLengthOne(length int) bool {
	if length == 1 {
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

	defer path.Close()
	return models
}
