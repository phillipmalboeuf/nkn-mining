package rpcApiResponse

type Base struct {
	Id      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type NullData struct{}
