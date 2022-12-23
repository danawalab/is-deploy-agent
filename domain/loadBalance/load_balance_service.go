package loadBalance

import (
	"bufio"
	"is-deploy-agent/domain"
	"is-deploy-agent/utils"
	"log"
	"os"
)

// CheckLbStatus
// uriworkeermap.properties의 값과 setting.json의 lbMap 값과 비교하여 같은 값이 있다면 해당 podName을 반환
// 또는 nodeName을 반환 해당하는 podName과 nodeName이 없으면 Not Match로 반환
func CheckLbStatus() (string, error) {
	path, err := getPropertiesPath()
	if err != nil {
		log.Println(err)
		return "", err
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println("uriworkermap.properties를 찾지 못했습니다. ", err)
		return "", err
	}
	defer file.Close()

	uriWorkerMap := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		uriWorkerMap = append(uriWorkerMap, scanner.Text())
	}
	uriWorkerMapLength := len(uriWorkerMap)

	node, err := utils.GetSettingJson()
	if err != nil {
		log.Println(err)
		return "", err
	}
	podLength := len(node.PodList)

	podMap := make([]string, 0)
	// setting.json에 pod의 lbMap이 2개 이상일 경우 해당 슬라이스를 통해 다른 결괏값이 나오는 걸 방지
	checkPodMap := make([]string, 0)
	podName := ""

	// pod 길이 만큼 uriworekr.properties의 값과 pod의 lbMap 값을 비교한다
	for pod := 0; pod < podLength; pod++ {
		podLbLength := len(node.PodList[pod].LbMap)

		for lb := 0; lb < podLbLength; lb++ {
			key := node.PodList[pod].LbMap[lb].Key
			value := node.PodList[pod].LbMap[lb].Value

			podMap = append(podMap, key+"="+value)
		}
		newArrayLength := len(podMap)

		// uriwokermap.properties의 길이와 setting.json에 pod의 lbMap의 길이가 같을 경우
		if uriWorkerMapLength == newArrayLength {
			for mapLength := 0; mapLength < uriWorkerMapLength; mapLength++ {
				if uriWorkerMap[mapLength] == podMap[mapLength] {
					podName = node.PodList[pod].Name
					// checkPodMap에 PodName을 넣어준다
					checkPodMap = append(checkPodMap, podName)
					return podName, err
				} else {
					break
				}
			}

			// checkPodMap의 중복 값을 제거 후 길이가 1가 같지 않으면 podName = ""로 초기화
			// 초기화를 안 하면 pod의 lbMap이 2개 이상일 경우 다른 결괏값이 나올 수 있음
			if len(deleteDuplicate(checkPodMap)) != 1 {
				podName = ""
			}
		}
		// podMap의 배열을 비어준다
		podMap = podMap[len(podMap):]
	}

	// Pod에서 podName을 못 찾았을 경우
	if podName == "" {
		nodeLbLength := len(node.LbMap)

		// uriwokermap.properties의 길이와 setting.json에 node의 lbMap의 길이가 같을 경우
		if uriWorkerMapLength == nodeLbLength {
			for nodeLb := 0; nodeLb < nodeLbLength; nodeLb++ {
				key := node.LbMap[nodeLb].Key
				value := node.LbMap[nodeLb].Value

				podMap = append(podMap, key+"="+value)
			}

			for mapLength := 0; mapLength < uriWorkerMapLength; mapLength++ {
				if uriWorkerMap[mapLength] == podMap[mapLength] {
					podName = node.Name
					return podName, err
				} else {
					break
				}
			}
		}
	}

	// node와 pod 둘 다 못 찾을 경우 Not Match로 리턴
	if err == nil {
		return "매칭되는 거 없음", err
	} else {
		return err.Error(), err
	}
}

func deleteDuplicate(array []string) []string {
	newArray := make([]string, 0)
	checkArray := make(map[string]struct{})

	for _, val := range array {
		if _, ok := checkArray[val]; !ok {
			checkArray[val] = struct{}{}
			newArray = append(newArray, val)
		}
	}
	return newArray
}

// Restore
// uriworkermap.properties 설정을 setting.json에 설정한 원래 값으로 복구
func Restore() error {
	path, err := getPropertiesPath()
	if err != nil {
		log.Println(err)
		return err
	}

	loadbalancerMap, err := getNodeLbMap()
	if err != nil {
		log.Println(err)
		return err
	}
	lbLength := len(loadbalancerMap)

	if isLengthOne(lbLength) {
		lb := getLbMapResult(loadbalancerMap)
		err = writeFileString(path, lb)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		err = writeFileArray(path, loadbalancerMap, lbLength)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return err
}

// Exclude
// worker와 setting,json에 podList의 name 같으면 해당 pod는 uriworkermap.propertis를 수정하여 로드밸런서에 제외
func Exclude(worker string) error {
	path, err := getPropertiesPath()
	if err != nil {
		log.Println(err)
		return err
	}

	excludeMap, err := getPodLbMap(worker)
	if err != nil {
		log.Println(err)
		return err
	}
	exLength := len(excludeMap)

	if isLengthOne(exLength) {
		ex := getLbMapResult(excludeMap)
		err = writeFileString(path, ex)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		err = writeFileArray(path, excludeMap, exLength)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return err
}

// lbMap이 1개면 바로 key=value로 반환
func getLbMapResult(lbMap []domain.UriMap) string {
	key := lbMap[0].Key
	value := lbMap[0].Value
	result := key + "=" + value

	return result
}

// setting.json podList의 lbMap 반환
func getPodLbMap(worker string) ([]domain.UriMap, error) {
	node, err := utils.GetSettingJson()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	podLength := len(node.PodList)
	var excludeMap []domain.UriMap

	for pods := 0; pods < podLength; pods++ {
		pod := node.PodList[pods]
		name := pod.Name

		if utils.IsNameEqual(name, worker) {
			exLength := len(pod.LbMap)

			for excludeMaps := 0; excludeMaps < exLength; excludeMaps++ {
				key := pod.LbMap[excludeMaps].Key
				value := pod.LbMap[excludeMaps].Value

				excludeMap = append(excludeMap, domain.UriMap{Key: key, Value: value})
			}
			break
		}
	}
	return excludeMap, err
}

// setting.json node의 lbMap 반환
func getNodeLbMap() ([]domain.UriMap, error) {
	node, err := utils.GetSettingJson()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	modelLength := len(node.LbMap)
	var loadbalancerMap []domain.UriMap

	for loadbalancer := 0; loadbalancer < modelLength; loadbalancer++ {
		key := node.LbMap[loadbalancer].Key
		value := node.LbMap[loadbalancer].Value

		loadbalancerMap = append(loadbalancerMap, domain.UriMap{Key: key, Value: value})
	}
	return loadbalancerMap, err
}

// lbMap이 1개일 경우
func writeFileString(path string, workerMap string) error {
	file, err := getUriWorkerMapFile(path)
	defer file.Close()

	_, err = file.Write([]byte(workerMap))
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

// lbMap이 2개 이상일 경우
func writeFileArray(path string, workerMaps []domain.UriMap, length int) error {
	file, err := getUriWorkerMapFile(path)
	defer file.Close()

	for workerMap := 0; workerMap < length; workerMap++ {
		key := workerMaps[workerMap].Key
		value := workerMaps[workerMap].Value

		lb := key + "=" + value + "\n"
		_, err = file.Write([]byte(lb))
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return err
}

// uriworkermap.properties 반환
func getUriWorkerMapFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return nil, err
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
func getPropertiesPath() (string, error) {
	node, err := utils.GetSettingJson()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return node.Path, err
}
