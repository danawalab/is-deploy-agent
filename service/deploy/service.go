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

func Deploy(worker string) {
	pullWAR()
	removeWAR()
	copyWAR()
}

func copyWAR() {
	webappPath, fileName := getWebappPathAndFileName()

	origin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer origin.Close()

	copy, err := os.Create(webappPath + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer copy.Close()

	file, err := io.Copy(copy, origin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)
}

func removeWAR() {
	webappPath, fileName := getWebappPathAndFileName()

	err := os.Remove(webappPath + fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func getWebappPathAndFileName() (string, string) {
	models := readJson()

	webappPath := models[0].NodeList[0].PodList[0].WebappPath
	fileName := models[0].NodeList[0].PodList[0].FileName

	return webappPath, fileName
}

func pullWAR() {
	jenkinsURL := getJenkinsURL()
	_, fileName := getWebappPathAndFileName()

	response, err := grab.Get(".", jenkinsURL)
	if err != nil {
		log.Fatal(err)
	}

	response.Filename = fileName
	fmt.Println("Download Complete", response)
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
