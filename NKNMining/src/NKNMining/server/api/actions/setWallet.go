package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"NKNMining/storage"
	"NKNMining/crypto"
)

var SetWalletAPI IRestfulAPIAction = &setWallet{}

type setWalletData struct {
	Wallet string
	Key string
}

type setWallet struct {
	restfulAPIBase
}

func (s *setWallet) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/set/wallet"
}

func (s *setWallet) Action(ctx *gin.Context) {
	s.response = apiServerResponse.New(ctx)

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


	walletInfoJsonStr, err := crypto.AesDecrypt(basicData.Data,
		storage.NKNSetupInfo.GetRequestKey())
	if nil != err {
		s.response.BadRequest("invalid request data!")
		return
	}

	walletInfoData := &setWalletData{}
	err = json.Unmarshal([]byte(walletInfoJsonStr), walletInfoData)
	if nil != err {
		s.response.BadRequest("invalid wallet setting data!")
		return
	}

	walletStorage := &storage.Wallet{}
	checked, err := walletStorage.Save(walletInfoData.Wallet)
	if nil != err {
		if checked {
			s.response.InternalServerError("save wallet failed!")
		} else {
			s.response.BadRequest("invalid wallet data!")
		}
	} else {
		storage.NKNSetupInfo.CurrentStep = storage.SETUP_STEP_SUCCESS
		storage.NKNSetupInfo.WKey = walletInfoData.Key
		storage.NKNSetupInfo.Save()
		s.response.Success(nil)
	}

	return
}
