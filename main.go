package main

import (
	"github/pull_request_webhook/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	infrastructure.Routes(engine)

	engine.Run(":4000")

}
