package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"NKNMining/common"
	"NKNMining/container"
)

var GetVersionAPI IRestfulAPIAction = &getVersion{}

type getVersion struct {
	restfulAPIBase
}

func (g *getVersion) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/version"
}

func (g *getVersion) Action(ctx *gin.Context) {
	g.response = apiServerResponse.New(ctx)
	nodeVersion := "..."

	nodeVersion, err := container.NodeStatus.SyncRun([]string{"-version"}, "")
	if nil != err {
		nodeVersion = "UNKNOWN"
	}

	g.response.Success(map[string]string{
		"NodeVersion": nodeVersion,
		"ShellVersion": common.NS_VERSION,
		"GinVersion": gin.Version,
		})

	return
}
