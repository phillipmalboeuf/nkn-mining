package config

type NodeShellConfig struct {
	LogFile			string
	ServerPort  	string
}

var ShellConfig = NodeShellConfig {
	LogFile: "NKNMining.log",
	ServerPort: "8181",
}

