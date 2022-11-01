package fetch

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"github.com/gin-gonic/gin"
	"is-deploy-agent/utils"
	"os/exec"
)

// FetchJson
// setting.json을 직접 서버에 접속에 vi에 수정하기 번거로워 console에서 수정 후 console에서 다운 받아 업데이트
func FetchJson(*gin.Context) {
	json := utils.GetJson()
	consoleAddress := json.ConsoleInfo

	_, err := grab.Get(".", consoleAddress)
	if err != nil {
		fmt.Println(err)
		//todo log로 변경
	}

	cmd := exec.Command("./sync.sh")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		//todo log로 변경
	}

	fmt.Println("setting.json Update", string(output))
	//todo log로 변경
}
