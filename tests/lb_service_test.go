package tests

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func ExcludeTestReadJsonValue(t *testing.T) {
	node := GetJsonToTest()
	arrayLength := node.PodList[0].LbMap
	fmt.Println(len(arrayLength))

	for i := 0; i < len(arrayLength); i++ {
		key := node.PodList[0].LbMap[i].Key
		value := node.PodList[0].LbMap[i].Value

		fmt.Printf("TestReadJsonValue, %s = %s", key, value)
		fmt.Println()
	}
}

func ExcludeTestJsonValueSave(t *testing.T) {
	node := GetJsonToTest()
	arrayLength := node.PodList[0].LbMap

	fmt.Println(len(arrayLength))

	var newArray []ExcludeMap
	for i := 0; i < len(arrayLength); i++ {
		key := node.PodList[0].LbMap[i].Key
		value := node.PodList[0].LbMap[i].Value

		newArray = append(newArray, ExcludeMap{key, value})
	}

	fmt.Println("TestJsonValueSave = ", newArray, len(newArray))
}

func ExcludeTestFindByName(t *testing.T) {
	worker := "WAS1"
	node := GetJsonToTest()

	length := len(node.PodList)
	var newArray []ExcludeMap

	for i := 0; i < length; i++ {
		pod := node.PodList[i]
		name := pod.Name

		if worker == name {
			exLength := len(pod.LbMap)

			for j := 0; j < exLength; j++ {
				key := pod.LbMap[j].Key
				value := pod.LbMap[j].Value

				newArray = append(newArray, ExcludeMap{key, value})
			}
			break
		}
	}

	fmt.Println("TestFindByName = ", newArray, len(newArray))
}

func TestReadProperties(t *testing.T) {
	path := GetJsonToTest().Path
	file, _ := os.Open(path)
	defer file.Close()

	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	resultLength := len(result)

	node := GetJsonToTest()

	length := len(node.PodList)
	newArray := make([]string, 0)
	nameArray := make([]string, 0)
	podName := ""

	for i := 0; i < length; i++ {
		println("Pod Search")
		pod := node.PodList[i]
		exLength := len(pod.LbMap)

		for j := 0; j < exLength; j++ {
			key := pod.LbMap[j].Key
			value := pod.LbMap[j].Value

			newArray = append(newArray, key+"="+value)
		}
		newArrayLength := len(newArray)
		//println("newArrayLength ", i, " =", newArrayLength)

		if resultLength == newArrayLength {
			for k := 0; k < resultLength; k++ {
				if result[k] == newArray[k] {
					podName = pod.Name
					nameArray = append(nameArray, podName)
					println(podName)
				} else {
					break
					//podName = ""
					//nameArray = append(nameArray, podName)
				}
			}

			a := make([]string, 0)
			b := make(map[string]struct{})

			for _, val := range nameArray {
				if _, ok := b[val]; !ok {
					b[val] = struct{}{}
					println(val)
					a = append(a, val)
				}
			}

			if len(a) != 1 {
				podName = ""
			}
		}
		newArray = newArray[len(newArray):]
	}

	if podName == "" {
		println("Node Search")
		nodeLbLength := len(node.LbMap)

		if resultLength == nodeLbLength {
			for x := 0; x < nodeLbLength; x++ {
				newArray = append(newArray, node.LbMap[x].Key+"="+node.LbMap[x].Value)
			}

			for y := 0; y < resultLength; y++ {
				if result[y] == newArray[y] {
					podName = node.Name
				} else {
					break
				}
			}
		}
	}

	if podName == "" {
		podName = "Not Match"
	}

	println("Name", podName)
}

type ExcludeMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
