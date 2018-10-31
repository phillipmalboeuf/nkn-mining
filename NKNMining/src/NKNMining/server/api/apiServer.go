package apiServer

import (
	. "NKNMining/server/api/actions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	apiList = []IRestfulAPIAction{
		GetStatusAPI,
		GetVersionAPI,
		SetAccountAPI,
		SetWalletAPI,
		StartNodeAPI,
		StopNodeAPI,
		ResetShellAPI,
		LoginAPI,
	}
)

func InitRouters(router *gin.Engine, baseURI string) {
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
				c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
			}
			c.AbortWithStatus(http.StatusOK)
		}
	})

	setApiRouters(router, baseURI)
}

func setApiRouters(router *gin.Engine, baseURI string) {
	for _, v := range apiList {
		router.Any(v.URI(baseURI), v.Action)
	}
}
