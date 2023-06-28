package main

import (
	"website/server/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	routers.InitRouter(router)

	router.Run(":80")
	//router.RunTLS(":443", config.SSL.CertFile, config.SSL.CertKey)
}
