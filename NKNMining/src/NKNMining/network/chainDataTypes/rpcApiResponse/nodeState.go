package rpcApiResponse

type NodeInfoResult struct {
	Addr              string `json:"addr"`   // The node's IP address
	Height            uint32 `json:"height"` // The node latest block height
	HttpProxyPort     uint16 `json:"httpProxyPort"`
	ID                string `json:"id"` // The nodes's id
	JsonRpcPort       uint16 `json:"jsonRpcPort"`
	PublicKey         string `json:"publicKey"`
	RelayMessageCount uint32 `json:"relayMessageCount"`
	SyncState         string `json:"syncState"` // node block sync status
	Version           string `json:"version"`   // The network protocol the node used
	WebsocketPort     uint16 `json:"websocketPort"`
}

type NodeInfo struct {
	Base
	Result NodeInfoResult `json:"result"`
}
