package rpcApiResponse

type NodeAddrResult struct {
	Addr          string `json:"addr"`   // The node's IP address
	Height        uint32 `json:"height"` // The node latest block height
	HttpProxyPort uint16 `json:"httpProxyPort"`
	ID            string `json:"id"` // The nodes's id
	IsOutBound    bool   `json:"isOutBound"`
	JsonRpcPort   uint16 `json:"jsonRpcPort"`
	PublicKey     string `json:"publicKey"`
	SyncState     string `json:"syncState"` // node block sync status
	WebsocketPort uint16 `json:"websocketPort"`
}

type NodeAddr struct {
	Base
	Result []NodeAddrResult `json:"result"`
}
