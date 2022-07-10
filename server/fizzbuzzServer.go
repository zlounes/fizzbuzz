package server

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/zlounes/fizzbuzz/config"
	"github.com/zlounes/fizzbuzz/handlers"
)

type FizzBuzzServer struct {
	handler      http.Handler
	serverConfig ServerConfig
	server       http.Server
}

func NewServer(serverConfig ServerConfig) *FizzBuzzServer {

	m := http.NewServeMux()
	m.Handle("/fizzbuzz", &handlers.FizzBuzzHandler{})
	httpServer := http.Server{Addr: serverConfig.GetServerPort(), Handler: m}
	result := &FizzBuzzServer{server: httpServer, serverConfig: serverConfig}
	return result
}

func (fizzbuzzServer FizzBuzzServer) Run() {
	fmt.Println("Fizzbuzz listen on ", fizzbuzzServer.serverConfig.GetServerPort(), "...")
	err := fizzbuzzServer.server.ListenAndServe()
	if err != nil {
		log.Panic(fmt.Errorf("Unable to listen: %s", err))
	}
}
