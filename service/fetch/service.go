package fetch

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"github.com/gin-gonic/gin"
	"is-deploy-agent/utils"
	"log"
	"os/exec"
)

func FetchJson(*gin.Context) {
	models := utils.GetJson()
	consoleAddress := models[0].ConsoleInfo

	_, err := grab.Get(".", consoleAddress)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("./sync.sh")
	output, err := cmd.Output()

	fmt.Println("setting.json Update", string(output))
}
