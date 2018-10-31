package rpcApiResponse

type NodeAddrResult struct {
	Time     int64
	Services uint64
	IpAddr   [16]byte
	Port     uint16
	ID       uint64
}

type NodeAddr struct {
	Base
	Result []NodeAddrResult `json:"Result"`
}
