package loadbalance

import (
	"fmt"
	"is-deploy-agent/model"
	"is-deploy-agent/utils"
	"os"
)

func Restore() {
	path := getPropertiesPath()
	loadbalancerMap := getLoadbalancerMap()
	lbLength := len(loadbalancerMap)

	if isLengthOne(lbLength) {
		lb := getWorkerMapResult(loadbalancerMap)
		writeFileString(path, lb)
	} else {
		writeFileArray(path, loadbalancerMap, lbLength)
	}
}

func Exclude(worker string) {
	path := getPropertiesPath()
	excludeMap := getExcludeMap(worker)
	exLength := len(excludeMap)

	if isLengthOne(exLength) {
		ex := getWorkerMapResult(excludeMap)
		writeFileString(path, ex)
	} else {
		writeFileArray(path, excludeMap, exLength)
	}
}

func getWorkerMapResult(workerMap []model.UriMap) string {
	key := workerMap[0].Key
	value := workerMap[0].Value
	result := key + "=" + value

	return result
}

func getExcludeMap(worker string) []model.UriMap {
	json := utils.GetJson()
	podLength := len(json.Node.PodList)
	var excludeMap []model.UriMap

	for pods := 0; pods < podLength; pods++ {
		pod := json.Node.PodList[pods]
		name := pod.Name

		if utils.IsNameEqual(name, worker) {
			exLength := len(pod.LbMap)

			for excludeMaps := 0; excludeMaps < exLength; excludeMaps++ {
				key := pod.LbMap[excludeMaps].Key
				value := pod.LbMap[excludeMaps].Value

				excludeMap = append(excludeMap, model.UriMap{Key: key, Value: value})
			}
			break
		}
	}
	return excludeMap
}

func getLoadbalancerMap() []model.UriMap {
	json := utils.GetJson()
	modelLength := len(json.Node.LbMap)
	var loadbalancerMap []model.UriMap

	if isLengthOne(modelLength) {
		key := json.Node.LbMap[0].Key
		value := json.Node.LbMap[0].Value

		loadbalancerMap = append(loadbalancerMap, model.UriMap{Key: key, Value: value})
		return loadbalancerMap
	} else {
		for loadbalancer := 0; loadbalancer < modelLength; loadbalancer++ {
			key := json.Node.LbMap[loadbalancer].Key
			value := json.Node.LbMap[loadbalancer].Value

			loadbalancerMap = append(loadbalancerMap, model.UriMap{Key: key, Value: value})
		}
		return loadbalancerMap
	}
}

func writeFileString(path string, workerMap string) {
	file, err := getFile(path)

	_, err = file.Write([]byte(workerMap))
	if err != nil {
		fmt.Println(err)
		//todo log로 변경
	}
	defer file.Close()
}

func writeFileArray(path string, workerMaps []model.UriMap, length int) {
	file, err := getFile(path)

	for workerMap := 0; workerMap < length; workerMap++ {
		key := workerMaps[workerMap].Key
		value := workerMaps[workerMap].Value

		lb := key + "=" + value + "\n"
		_, err = file.Write([]byte(lb))
		if err != nil {
			return
		}
	}
	defer file.Close()
}

func getFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		//todo log로 변경
	}
	return file, err
}

func isLengthOne(length int) bool {
	if length == 1 {
		return true
	}
	return false
}

func getPropertiesPath() string {
	json := utils.GetJson()
	return json.Node.Path
}
