package config

import (
	"flag"
	"strconv"
)

type InputData struct {
	Int1    int
	Int2    int
	Limit   int
	String1 string
	String2 string
}

type ServerConfig struct {
	port int
}

func RetrieveServerConfigFromCLI() ServerConfig {
	var serverPort int
	flag.IntVar(&serverPort, "port", 8080, "define the server listen port")
	flag.Parse()
	return ServerConfig{port: serverPort}
}

func (serverConfig *ServerConfig) GetServerPort() string {
	return ":" + strconv.Itoa(serverConfig.port)
}
