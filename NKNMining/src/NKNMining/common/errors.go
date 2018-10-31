package common

import "fmt"

type NodeShellError struct {
	Code     int
	UserData interface{}
}

func (n *NodeShellError) FmtOutput() string {
	var desc = errorDescription(n.Code)
	if nil != n.UserData {
		return fmt.Sprintf(desc+": %v", n.UserData)
	}
	return desc
}

func (n *NodeShellError) Error() string {
	return n.FmtOutput()
}

const (
	NS_SUCCESS = iota
	NS_ERR_FILE_NOT_FOUND
	NS_ERR_CAN_NOT_READ_FILE
	NS_ERR_JSON_UNMARSHAL
	NS_ERR_CONFIG
	NS_ERR_DATA_TYPE
	NS_ERR_NO_SUCH_DATA
	NS_ERR_NO_SUCH_METHOD
	NS_ERR_INDEX_OUT_OF_RANGE
	NS_ERR_END
)

var errorDesc = map[int]string{
	NS_SUCCESS:                "ok",
	NS_ERR_FILE_NOT_FOUND:     "file not found",
	NS_ERR_CAN_NOT_READ_FILE:  "can not read file",
	NS_ERR_JSON_UNMARSHAL:     "json unmarshal failed",
	NS_ERR_CONFIG:             "setup config failed",
	NS_ERR_DATA_TYPE:          "data type error",
	NS_ERR_NO_SUCH_DATA:       "data not found",
	NS_ERR_NO_SUCH_METHOD:     "no such method",
	NS_ERR_INDEX_OUT_OF_RANGE: "index out of range",
}

func errorDescription(errorCode int) string {
	if errorCode >= NS_ERR_END {
		return ""
	}

	return errorDesc[errorCode]
}
