package status

import (
	"NKNMining/container"
	"NKNMining/storage"
	"NKNMining/common"
)

func statusFromNodeContainer(status *int) (ret bool) {
	ret = true
	switch container.Node.Status() {
	case common.NKN_NODE_EXITED:
		*status = common.NS_STATUS_NODE_EXITED
		break

	case common.NKN_NODE_RUNNING:
		*status = common.NS_STATUS_NODE_RUNNING
		break

	default:
		ret = false
		break
	}

	return
}

func GetServerStatus() (status int, errInfo string) {
	status = common.NS_STATUS_CTEATE_ACCOUNT
	errInfo = ""

	switch storage.NKNSetupInfo.CurrentStep {
	case storage.SETUP_STEP_CREATE_ACCOUNT:
		break

	case storage.SETUP_STEP_GEN_WALLET:
		status = common.NS_STATUS_GEN_WALLET
		break

	case storage.SETUP_NODE_UPDATE:
		status = common.NS_STATUS_NODE_UPDATE
		break

	case storage.SETUP_STEP_SUCCESS:
		if !statusFromNodeContainer(&status) {
			errInfo = "invalid node status!"
		}
		break

	default:
		errInfo = "invalid setup step!"
		break
	}

	return
}

func CanStartNode() bool {
	status, errInfo := GetServerStatus()
	return common.NS_STATUS_NODE_EXITED == status &&
		"" == errInfo
}