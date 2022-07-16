package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zlounes/fizzbuzz/config"
	. "github.com/zlounes/fizzbuzz/config"
	"github.com/zlounes/fizzbuzz/metrics"
)

//Manage the http server with the fizzbuzz handlers
//mapped to :
// Post /fizzbuzz for the fizzbuzz calculator
// Get /fizzbuzz/stat for the statistics
type FizzBuzzServer struct {
	handler          http.Handler
	serverConfig     ServerConfig
	server           *http.Server
	serverMux        *http.ServeMux
	stat             *metrics.FizzbuzzStat
	channelInputStat chan<- InputData
	channelCheckStat chan<- chan metrics.BestHint
}

//Accept a ServerConfig for defining the port
//dont start to listent on socket until Run() is called
func NewServer(serverConfig ServerConfig) *FizzBuzzServer {
	channelInputStat := make(chan config.InputData)
	channelCheckStat := make(chan chan metrics.BestHint)
	stat := metrics.NewFizzbuzzStat(channelInputStat, channelCheckStat)
	m := http.NewServeMux()
	m.Handle("/fizzbuzz", &fizzBuzzHandler{channelInputStat: channelInputStat})
	m.Handle("/fizzbuzz/stat", &metricHandler{channelCheckStat: channelCheckStat})
	httpServer := http.Server{Addr: serverConfig.GetServerPort(), Handler: m}
	result := &FizzBuzzServer{server: &httpServer, serverMux: m, serverConfig: serverConfig, stat: stat,
		channelInputStat: channelInputStat, channelCheckStat: channelCheckStat}
	return result
}

//start listening on SeverConfig.serverPort
func (fizzbuzzServer *FizzBuzzServer) Run() {
	fmt.Println("Fizzbuzz listen on ", fizzbuzzServer.serverConfig.GetServerPort(), ", Use Ctrl+C to Stop...")
	err := fizzbuzzServer.server.ListenAndServe()
	if err != nil {
		log.Panic(fmt.Errorf("Unable to listen: %s", err))
	}
}

//used by unit tests, avoiding listening socket
func (fizzbuzzServer *FizzBuzzServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fizzbuzzServer.serverMux.ServeHTTP(w, r)
}

//close the http server
func (fizzbuzzServer *FizzBuzzServer) Stop() {
	fizzbuzzServer.server.Close()
	fizzbuzzServer.stat.Stop()
}

func returnHttpError(w http.ResponseWriter, statusCode int, reasonStr string) {
	http.Error(w, reasonStr, statusCode)
}

func setStatusOk(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
}

func methodNotAllowed(w http.ResponseWriter, acceptedRest string) {
	returnHttpError(w, http.StatusMethodNotAllowed, fmt.Sprintf("Sorry, only %s methods is supported.", acceptedRest))
}
