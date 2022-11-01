package loadbalance

import (
	"fmt"
	"is-deploy-agent/utils"
	"testing"
)

func TestReadJsonValue(t *testing.T) {
	json := utils.GetJsonToTest()
	arrayLength := json.Node.PodList[0].LbMap
	fmt.Println(len(arrayLength))

	for i := 0; i < len(arrayLength); i++ {
		key := json.Node.PodList[0].LbMap[i].Key
		value := json.Node.PodList[0].LbMap[i].Value

		fmt.Printf("TestReadJsonValue, %s = %s", key, value)
		fmt.Println()
	}
}

func TestJsonValueSave(t *testing.T) {
	json := utils.GetJsonToTest()
	arrayLength := json.Node.PodList[0].LbMap

	fmt.Println(len(arrayLength))

	var newArray []ExcludeMap
	for i := 0; i < len(arrayLength); i++ {
		key := json.Node.PodList[0].LbMap[i].Key
		value := json.Node.PodList[0].LbMap[i].Value

		newArray = append(newArray, ExcludeMap{key, value})
	}

	fmt.Println("TestJsonValueSave = ", newArray, len(newArray))
}

func TestFindByName(t *testing.T) {
	worker := "WAS1"
	json := utils.GetJsonToTest()

	length := len(json.Node.PodList)
	var newArray []ExcludeMap

	for i := 0; i < length; i++ {
		pod := json.Node.PodList[i]
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

type ExcludeMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
