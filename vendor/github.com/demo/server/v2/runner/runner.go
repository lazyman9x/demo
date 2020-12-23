package runner

import (
	"fmt"
	"log"
	"net/http"

	"github.com/demo/server/v2/config"
	"github.com/gin-gonic/gin"
)

func Run(engine *gin.Engine, conf *config.Configuration) {
	var httpHandler http.Handler = engine
	addr := fmt.Sprintf("%s:%d", conf.Server.ListenAddr, conf.Server.Port)
	fmt.Println("Started Listening for plain HTTP connection on " + addr)
	log.Fatal(http.ListenAndServe(addr, httpHandler))
}
