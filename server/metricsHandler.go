package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zlounes/fizzbuzz/metrics"
)

//Http handler returning the FissBuzz config the more used  and the count
type metricHandler struct {
	http.Handler
	channelCheckStat chan<- chan metrics.BestHint
}

//on Rest GET : return json response for metrics.BestHint content
func (handler *metricHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if err := handler.processMetrics(w, req); err != nil {
			returnHttpError(w, http.StatusInternalServerError, fmt.Sprintf("Error while building response : %v", err))
		}
		setStatusOk(w, "application/json; charset=utf-8")
	default:
		methodNotAllowed(w, "GET")
	}
}

func (handler *metricHandler) processMetrics(w http.ResponseWriter, req *http.Request) error {
	statResult := retrieveStat(handler.channelCheckStat)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(statResult); err != nil {
		return err
	}
	return nil
}

func retrieveStat(channelCheckStat chan<- chan metrics.BestHint) metrics.BestHint {
	chanDone := make(chan metrics.BestHint)
	channelCheckStat <- chanDone
	statResult := <-chanDone
	return statResult
}
