package main

import (
	"github.com/zlounes/fizzbuzz/config"
	"github.com/zlounes/fizzbuzz/server"
)

func main() {
	fizzbuzzServer := server.NewServer(config.RetrieveServerConfigFromCLI())
	fizzbuzzServer.Run()
}
