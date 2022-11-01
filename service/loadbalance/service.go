package loadbalance

import (
	"fmt"
	"is-deploy-agent/model"
	"is-deploy-agent/utils"
	"os"
)

// Restore
// uriworkermap.properties 설정을 setting.json에 설정한 원래 값으로 복구
func Restore() {
	path := getPropertiesPath()
	loadbalancerMap := getNodeLbMap()
	lbLength := len(loadbalancerMap)

	if isLengthOne(lbLength) {
		lb := getLbMapResult(loadbalancerMap)
		writeFileString(path, lb)
	} else {
		writeFileArray(path, loadbalancerMap, lbLength)
	}
}

// Exclude
// worker와 setting,json에 podList의 name 같으면 해당 pod는 uriworkermap.propertis를 수정하여 로드밸런서에 제외
func Exclude(worker string) {
	path := getPropertiesPath()
	excludeMap := getPodLbMap(worker)
	exLength := len(excludeMap)

	if isLengthOne(exLength) {
		ex := getLbMapResult(excludeMap)
		writeFileString(path, ex)
	} else {
		writeFileArray(path, excludeMap, exLength)
	}
}

// lbMap이 1개면 바로 key=value로 반환
func getLbMapResult(lbMap []model.UriMap) string {
	key := lbMap[0].Key
	value := lbMap[0].Value
	result := key + "=" + value

	return result
}

// setting.json podList의 lbMap 반환
func getPodLbMap(worker string) []model.UriMap {
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

// setting.json node의 lbMap 반환
func getNodeLbMap() []model.UriMap {
	json := utils.GetJson()
	modelLength := len(json.Node.LbMap)
	var loadbalancerMap []model.UriMap

	for loadbalancer := 0; loadbalancer < modelLength; loadbalancer++ {
		key := json.Node.LbMap[loadbalancer].Key
		value := json.Node.LbMap[loadbalancer].Value

		loadbalancerMap = append(loadbalancerMap, model.UriMap{Key: key, Value: value})
	}
	return loadbalancerMap
}

// lbMap이 1개일 경우
func writeFileString(path string, workerMap string) {
	file, err := getFile(path)

	_, err = file.Write([]byte(workerMap))
	if err != nil {
		fmt.Println(err)
		//todo log로 변경
	}
	defer file.Close()
}

// lbMap이 2개 이상일 경우
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

// uriworkermap.properties 반환
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

// setting.json의 uriworkermap.properties 경로 반환
func getPropertiesPath() string {
	json := utils.GetJson()
	return json.Node.Path
}
