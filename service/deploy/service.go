package deploy

import (
	"encoding/json"
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"io"
	"is-deploy-agent/model"
	"log"
	"os"
)

func Deploy(node int, worker string) {
	fileName := pullWAR(node, worker)
	removeWAR(node, worker)
	copyWAR(node, worker, fileName)
}

func copyWAR(node int, worker string, fileName string) {
	webappPath, web := getWebappPathAndFileName(node, worker)

	origin, err := os.Open("./" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer origin.Close()

	copy, err := os.Create(webappPath + web)
	if err != nil {
		log.Fatal(err)
	}
	defer copy.Close()

	file, err := io.Copy(copy, origin)
	if err != nil {
		log.Fatal(err)
	}

	//os.Remove("./" + fileName) not work
	fmt.Println(file)
}

func removeWAR(node int, worker string) {
	webappPath, fileName := getWebappPathAndFileName(node, worker)

	err := os.Remove(webappPath + fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func pullWAR(node int, worker string) string {
	jenkinsURL := getJenkinsURL()

	response, err := grab.Get(".", jenkinsURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download Complete", response)
	return response.Filename
}

func getWebappPathAndFileName(node int, worker string) (string, string) {
	models := readJson()
	podLength := len(models[0].NodeList[node].PodList)

	var webappPath string
	var fileName string
	for pods := 0; pods < podLength; pods++ {
		pod := models[0].NodeList[node].PodList[pods]
		name := pod.Name

		if name == worker {
			webappPath = pod.WebappPath
			fileName = pod.FileName
		}
	}
	return webappPath, fileName
}

func getJenkinsURL() string {
	models := readJson()

	basicURL := models[0].NodeList[0].JenkinsURL.BasicURL
	middleURL := models[0].NodeList[0].JenkinsURL.MiddleURL
	jobName := models[0].NodeList[0].JenkinsURL.JobName
	groupId := models[0].NodeList[0].JenkinsURL.GroupId
	artifactID := models[0].NodeList[0].JenkinsURL.ArtifactId
	version := models[0].NodeList[0].JenkinsURL.Version

	return basicURL + jobName + "/" + groupId + "$" + artifactID +
		middleURL + groupId + "/" + artifactID + "/" + version + "/" + artifactID + "-" + version + ".war"
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
