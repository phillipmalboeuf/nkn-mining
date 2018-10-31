package rpcApiResponse

//registered asset will be assigned to contract address
type Asset struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Precision   uint32 `json:"Precision"`
	AssetType   byte   `json:"AssetType"`
	RecordType  byte   `json:"RecordType"`
}
