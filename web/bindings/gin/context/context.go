package context

import (
	"github.com/alt-golang/config"
	"github.com/alt-golang/logger"
	"github.com/alt-golang/random-api-minimal/service"
	gin "github.com/alt-golang/random-api-minimal/web/bindings/gin"
	g "github.com/gin-gonic/gin"
)

var cfg = config.GetConfigFromDir("config")
var maximum, _ = config.Get("random.number.maximum")
var RandomNumber = service.RandomNumber{
	Logger:  logger.GetLogger("github.com/alt-golang/random-api-minimal/service/RandomNumber"),
	Maximum: int(maximum.(float64)),
}

var port, _ = config.Get("server.port")
var mode, _ = config.Get("server.mode")

var Server = gin.Server{
	Logger:       logger.GetLogger("github.com/alt-golang/random-api-minimal/web/bindings/gin/Server"),
	Port:         int(port.(float64)),
	Context:      "",
	Mode:         mode.(string),
	RandomNumber: RandomNumber,
}

func Start() {
	g.SetMode(mode.(string))
	Server.Engine = g.New()
	Server.Init()
	Server.Run()
}
