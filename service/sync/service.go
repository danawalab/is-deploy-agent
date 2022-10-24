package sync

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"github.com/gin-gonic/gin"
	"log"
)

func FetchJson(*gin.Context) {
	response, err := grab.Get(".", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("setting.json Update Complete", response)
}
