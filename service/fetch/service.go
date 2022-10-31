package fetch

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"github.com/gin-gonic/gin"
	"is-deploy-agent/utils"
	"os/exec"
)

func FetchJson(*gin.Context) {
	models := utils.GetJson()
	consoleAddress := models[0].ConsoleInfo

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
