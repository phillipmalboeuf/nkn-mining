package webServer

import (
	"NKNMining/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"NKNMining/server/web/const"
)

func InitRouters(router *gin.Engine, baseURI string, webDir string) {
	if !common.FileExist(webDir) {
		common.Log.Fatalf("web directory [%s] not exist", webDir)
		os.Exit(-1)
	}

	if strings.Contains(webDir, ".") {
		common.Log.Fatalf("web directory contains '.' [webDir]", webDir)
		os.Exit(-1)
	}

	setRouters(router, baseURI, webDir)
}

func setRouters(router *gin.Engine, baseURI string, webDir string) {
	router.StaticFS(baseURI + webServerConsts.WEB_SERVER_URI_BASE, http.Dir(webDir))
}
