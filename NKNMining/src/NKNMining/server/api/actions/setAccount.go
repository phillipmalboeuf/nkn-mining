package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"NKNMining/storage"
	"NKNMining/crypto"
)

var SetAccountAPI IRestfulAPIAction = &setAccount{}

type setAccountData struct {
	Account 	string
	Key 		string
}

type setAccount struct {
	restfulAPIBase
}

func (s *setAccount) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/set/account"
}

func (s *setAccount) Action(ctx *gin.Context) {
	s.response = apiServerResponse.New(ctx)
	if storage.SETUP_STEP_CREATE_ACCOUNT != storage.NKNSetupInfo.CurrentStep {
		s.response.BadRequest("invalid request!")
		return
	}

	inputJson, err := s.ExtractInput(ctx)
	if nil != err {
		s.response.BadRequest("invalid raw request!")
		return
	}

	basicData := &restfulAPIBaseRequest{}
	err = json.Unmarshal([]byte(inputJson), basicData)
	if nil != err {
		s.response.BadRequest("invalid request format!")
		return
	}

	accountInfoJsonStr, err := crypto.AesDecrypt(basicData.Data, storage.NKNSetupInfo.SerialNumber)
	if nil != err {
		s.response.BadRequest("invalid request data!")
		return
	}

	accountData := &setAccountData{}
	err = json.Unmarshal([]byte(accountInfoJsonStr), accountData)
	if nil != err {
		s.response.BadRequest("invalid account data!")
		return
	}

	storage.NKNSetupInfo.Account = accountData.Account
	storage.NKNSetupInfo.Key = accountData.Key
	storage.NKNSetupInfo.CurrentStep = storage.SETUP_STEP_GEN_WALLET
	storage.NKNSetupInfo.Save()

	s.response.Success(nil)
	return
}
