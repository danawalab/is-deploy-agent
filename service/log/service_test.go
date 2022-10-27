package log

import (
	"bufio"
	"fmt"
	"is-deploy-agent/utils"
	"log"
	"os"
	"testing"
)

func ExcludeTestLogFileRead(t *testing.T) {
	models := utils.GetJsonToTest()

	logPath := models[0].NodeList[0].PodList[0].LogPath

	logs, err := os.Open(logPath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(logs)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}

func TestTailLog(t *testing.T) {
	models := utils.GetJsonToTest()

	logPath := models[0].NodeList[0].PodList[0].LogPath

	logs, err := os.Open(logPath)
	if err != nil {
		log.Fatal(err)
	}

	logs.Seek(0, 2)
	scanner := bufio.NewScanner(logs)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}
