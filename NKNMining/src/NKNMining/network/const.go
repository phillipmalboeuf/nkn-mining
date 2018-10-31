package network

const (
	HTTP_METHOD_POST = "POST"
	HTTP_METHOD_GET  = "GET"

)

type HttpHeader struct {
	Name string
	Value string
}

const (
	HTTP_CONTENT_TYPE_JSON = "HTTP_CONTENT_TYPE_JSON"
	HTTP_ACCEPT_ENCODING = "HTTP_ACCEPT_ENCODING"
	HTTP_CONTENT_ENCODING = "HTTP_CONTENT_ENCODING"
)

func GetDefaultHeader(headerName string) (header *HttpHeader) {
	defaultHeaders := map[string] HttpHeader {
		HTTP_CONTENT_TYPE_JSON: {"Content-Type", "application/json"},
		HTTP_ACCEPT_ENCODING: {"Accept-Encoding", "gzip, deflate"},
		HTTP_CONTENT_ENCODING: {"Content-Encoding", "gzip"},
	}

	if def, ok := defaultHeaders[headerName]; ok {
		header = &HttpHeader{def.Name, def.Value}
	}

	return
}

const (
	RPC_API_BLOCK_HEIGHT 			= "BLOCK_HEIGHT"
	RPC_API_BLOCK_DETAIL_BY_HEIGHT  = "BLOCK_DETAIL"
	RPC_API_TX_DETAIL 				= "TX_DETAIL"
	RPC_API_NODE_STATE 				= "NODE_STATE"
	RPC_API_NODE_NEIGHBOR_ADDR		= "NODE_NEIGHBOR_ADDR"
)
