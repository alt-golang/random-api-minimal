package gin

import (
	"fmt"
	"github.com/alt-golang/logger"
	"github.com/alt-golang/random-api-minimal/service"
	g "github.com/gin-gonic/gin"
)

type Server struct {
	Logger       logger.Logger
	Port         int
	Context      string
	Mode         string
	Engine       *g.Engine
	RandomNumber service.RandomNumber
}

func (Server Server) Init() {

	Server.Engine.Use(
		g.LoggerWithFormatter(func(param g.LogFormatterParams) string {
			log := ""
			consoleLogger := Server.Logger.(logger.ConsoleLogger)
			if consoleLogger.IsInfoEnabled() {
				log = consoleLogger.Formatter.Format(param.TimeStamp, consoleLogger.Config.Category, consoleLogger.Config.Levels.GetNameForValue(logger.INFO), "",
					&struct {
						ClientIP string
						Path     string
					}{
						ClientIP: param.ClientIP,
						Path:     param.Path,
					})
			}

			return log + "\n"
		}))
	Server.Engine.Use(g.Recovery())
	Server.Engine.Use(func(context *g.Context) {
		Server.Logger.Info("middleware /*:" + fmt.Sprint(context.Request.Body))
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		context.Header("Access-Control-Allow-Header", "Content-Type, Authorization")
		context.Header("Content-Type", "application/json")
		context.Next()
	})

	Server.Engine.OPTIONS("/", func(context *g.Context) {
		Server.Logger.Info("OPTIONS /*: " + fmt.Sprint(context.Request.Body))
		context.Status(200)
	})

	Server.Engine.GET("/", func(context *g.Context) {
		Server.Logger.Info("GET /:" + fmt.Sprint(context.Request.Body))
		context.IndentedJSON(200, Server.RandomNumber.Get())
	})
}

func (Server Server) Run() {
	Server.Engine.Run(fmt.Sprintf("127.0.0.1:%d", Server.Port))
}
