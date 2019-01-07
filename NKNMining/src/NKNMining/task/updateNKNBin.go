package task

import (
	"time"
	"NKNMining/network/nknReleaseQuery"
	"NKNMining/common"
	"NKNMining/storage"
	"NKNMining/status"
	"NKNMining/container"
	"github.com/mholt/archiver"
	"NKNMining/crypto"
	"runtime"
	"os"
)

const (
	nkn_bin_file_path = "/bin"
)

var nknBinFirstUpdate = true

func nknBinNeedUpdate(version string) bool {
	if version != storage.NKNSetupInfo.BinVersion {
		return true
	}

	return  false
}

func mvNKNBin(from string, to string) error {
	return os.Rename(from, to)
}

func doBinUpdate(toVersion string, url string) {
	needRestart := false
	initialization := nknBinFirstUpdate
	currentStep := storage.NKNSetupInfo.CurrentStep


	if common.NknBinExists() {
		status.SetBinDownloaded()
	}

	if currentStep == storage.SETUP_NODE_UPDATE && !nknBinFirstUpdate {
		if !nknBinNeedUpdate(toVersion) {
			storage.NKNSetupInfo.CurrentStep = storage.SETUP_STEP_SUCCESS
			storage.NKNSetupInfo.Save()
		}
		return
	}

	defer func() {
		storage.NKNSetupInfo.CurrentStep = currentStep
		storage.NKNSetupInfo.Save()
	}()

	nknBinFirstUpdate = false
	var err error = nil

	if common.NknBinExists() {
		if nknBinFirstUpdate && storage.SETUP_NODE_UPDATE == storage.NKNSetupInfo.CurrentStep {
			currentStep = storage.SETUP_STEP_SUCCESS
		}

		binRunStatus, errInfo := status.GetServerStatus()

		if !nknBinNeedUpdate(toVersion) {
			if 	storage.NKNSetupInfo.CurrentStep == storage.SETUP_STEP_INIT ||
			 	storage.NKNSetupInfo.CurrentStep == storage.SETUP_STEP_GEN_WALLET ||
				storage.NKNSetupInfo.CurrentStep == storage.SETUP_STEP_CREATE_ACCOUNT {
				return
			}
			currentStep = storage.SETUP_STEP_SUCCESS
			return
		}

		if 	common.NS_STATUS_CTEATE_ACCOUNT == binRunStatus ||
			common.NS_STATUS_GEN_WALLET == binRunStatus ||
			"" != errInfo {
			return
		}

		if common.NS_STATUS_NODE_RUNNING == binRunStatus {
			needRestart = true
			container.Node.Stop()
		}
	}

	if 	storage.SETUP_STEP_CREATE_ACCOUNT != currentStep &&
		storage.SETUP_STEP_GEN_WALLET != currentStep {
		storage.NKNSetupInfo.CurrentStep = storage.SETUP_NODE_UPDATE
		storage.NKNSetupInfo.Save()
	}



	basicPath := common.GetCurrentDirectory() + nkn_bin_file_path
	unzippedBin := basicPath + "/" +runtime.GOOS + "-" + runtime.GOARCH + "/nknd"
	fileName := runtime.GOOS + "-" + runtime.GOARCH + "." + toVersion + ".zip"
	fullName := basicPath + "/" + fileName
	err = nknReleaseQuery.DownloadNKN(url, fullName, nil)
	if nil != err {
		return
	}

	err = archiver.Zip.Open(fullName, common.GetCurrentDirectory() + nkn_bin_file_path)

	if nil != err {
		common.Log.Error("unzip bin file failed: ", err)
		return
	}


	if common.IsWindowsOS() {
		err = mvNKNBin(unzippedBin, basicPath + "/nknd.exe")
	} else {
		err = mvNKNBin(unzippedBin, basicPath + "/nknd")
	}
	if nil != err {
		common.Log.Error("move bin file failed: ", err)
		return
	}

	if initialization {
		status.SetBinDownloaded()
	}

	storage.NKNSetupInfo.BinVersion = toVersion
	storage.NKNSetupInfo.Save()
	if needRestart {
		wKey, err := crypto.AesDecrypt(storage.NKNSetupInfo.WKey, storage.NKNSetupInfo.GetWalletKey())
		if nil != err {
			common.Log.Error(err)
			return
		}

		_, err = container.Node.AsyncRun([]string{"-p", wKey}, "")
		if nil != err {
			common.Log.Error(err)
			return
		}
	}

	return

}

func UpdateNKNBin()  {
	for {
		version, url, err := nknReleaseQuery.LastVersion()
		if nil != err {
			common.Log.Error(err)
		} else {
			doBinUpdate(version, url)
		}

		time.Sleep(120 * time.Second)
	}
}