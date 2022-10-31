package log

import (
	"bufio"
	"fmt"
	"github.com/hpcloud/tail"
	"is-deploy-agent/utils"
	"log"
	"os"
	"os/exec"
	"sync"
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
	mx := sync.RWMutex{}
	var ch = make(chan string)

	ta, err := tail.TailFile("../../sample/catalina.out", tail.Config{Follow: true, ReOpen: true, MustExist: true, Poll: true, Location: &tail.SeekInfo{Whence: 2}})
	if err != nil {
		fmt.Println(err)
	}

	var lgs string
	for line := range ta.Lines {
		go func() {
			mx.Lock()
			lg := line.Text
			ch <- lg
			mx.Unlock()
		}()
		lgs = <-ch
	}
	fmt.Println(lgs)
}

func ExcludeTailTypeA(t *testing.T) {
	cmd := exec.Command("tail", "-n 10", "../../sample/catalina.out")

	output, _ := cmd.Output()

	fmt.Println(string(output))
}
