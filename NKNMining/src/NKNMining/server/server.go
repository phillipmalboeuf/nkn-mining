package server

import (
	"NKNMining/server/web"
	"NKNMining/config"
	"NKNMining/server/api"
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"NKNMining/common"
	"NKNMining/storage"
)

func startMutexServer() {
	router := gin.New()

	router.Any("/mutex", func(ctx *gin.Context) {

	})
	go router.Run(common.NS_MUTEX_PORT)
}

func Start() {
	webDir := "web"
	serviceBaseURI := "/"

	gin.SetMode(gin.ReleaseMode)
	f, _ := os.Create("http.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	webServer.InitRouters(router, serviceBaseURI, webDir)
	apiServer.InitRouters(router, serviceBaseURI)
	if storage.IsRemote {
		go router.Run(":" + config.ShellConfig.ServerPort)
	} else {
		go router.Run("127.0.0.1:" + config.ShellConfig.ServerPort)
	}

	startMutexServer()
}
