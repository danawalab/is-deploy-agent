package log

import (
	"bufio"
	"fmt"
	"github.com/hpcloud/tail"
	"is-deploy-agent/utils"
	"log"
	"os"
	"os/exec"
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

func ExcludeTestTailLog(t *testing.T) {
	logPath := "../../sample/catalina.out"

	ta, _ := tail.TailFile(logPath, tail.Config{})

	for line := range ta.Lines {
		fmt.Println(line.Text)
	}

}

func ExcludeTailTypeA(t *testing.T) {
	cmd := exec.Command("tail", "-n 10", "../../sample/catalina.out")

	output, _ := cmd.Output()

	fmt.Println(string(output))
}
