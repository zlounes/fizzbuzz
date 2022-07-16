package config

import (
	"flag"
	"strconv"
)

type FizzBuzzInput struct {
	Int1    int
	Int2    int
	Limit   int
	String1 string
	String2 string
}

type ServerConfig struct {
	Port int
}

func NewInputData(int1 int, int2, limit int, string1 string, string2 string) FizzBuzzInput {
	return FizzBuzzInput{Int1: int1, Int2: int2, Limit: limit, String1: string1, String2: string2}
}
func RetrieveServerConfigFromCLI() ServerConfig {
	var serverPort int
	flag.IntVar(&serverPort, "port", 8080, "define the server listen port")
	flag.Parse()
	return ServerConfig{Port: serverPort}
}

func (serverConfig *ServerConfig) GetServerPort() string {
	return ":" + strconv.Itoa(serverConfig.Port)
}
