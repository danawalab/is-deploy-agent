package tests

import (
	"fmt"
	"os/exec"
	"testing"
)

func ExcludeTestExecShell(t *testing.T) {
	cmd := exec.Command("./test.sh")
	output, _ := cmd.Output()

	fmt.Println(output)
}

// danawa Repository에서는 해당 기능 삭제
/*
func ExcludeTestDownloadWAR(t *testing.T) {
	models := getJson()

	basicURL := models[0].NodeList[0].JenkinsURL.BasicURL
	middleURL := models[0].NodeList[0].JenkinsURL.MiddleURL
	jobName := models[0].NodeList[0].JenkinsURL.JobName
	groupId := models[0].NodeList[0].JenkinsURL.GroupId
	artifactID := models[0].NodeList[0].JenkinsURL.ArtifactId
	version := models[0].NodeList[0].JenkinsURL.Version

	jenkinsURL := basicURL + jobName + "/" + groupId + "$" + artifactID +
		middleURL + groupId + "/" + artifactID + "/" + version + "/" + artifactID + "-" + version + ".war"

	fmt.Println("URL = ", jenkinsURL)
	resp, err := grab.Get(".", jenkinsURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download Test", resp.Filename)
}

func ExcludeTestDeploy(t *testing.T) {
	models := getJson()

	webappPath := models[0].NodeList[0].PodList[0].WebappPath
	err := os.Remove(webappPath + "web.war")
	if err != nil {
		log.Fatal(err)
	}

	war, err := os.Open("./web.war")
	if err != nil {
		log.Fatal(err)
	}

	defer war.Close()

	copy, err4 := os.Create(webappPath + "web.war")
	if err4 != nil {
		log.Fatal(err)
	}
	defer copy.Close()

	_, err5 := io.Copy(copy, war)
	if err5 != nil {
		log.Fatal(err)
	}
}
*/
