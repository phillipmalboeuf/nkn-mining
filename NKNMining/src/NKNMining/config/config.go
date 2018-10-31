package config

const (
	NODE_RPC_SERVER_V2 = "2.0"
)

type NodeShellConfig struct {
	LogFile			string
	ServerPort  	string
}

var ShellConfig = NodeShellConfig {
	LogFile: "NKNMining.log",
	ServerPort: "8181",
}

