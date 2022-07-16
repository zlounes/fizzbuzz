package metrics

import (
	"github.com/zlounes/fizzbuzz/config"
)

const (
	HASH_SEPARATOR = ";"
)

type BestHint struct {
	Entry   config.InputData
	NbCalls int
}

type FizzbuzzStat struct {
	mapStatItems     map[config.InputData]int
	channelInputStat <-chan config.InputData
	channelCheckStat <-chan chan BestHint
	channelClose     chan bool
	bestHint         BestHint
}

func NewFizzbuzzStat(channelInputStat <-chan config.InputData, channelCheckStat <-chan chan BestHint) *FizzbuzzStat {
	mapStatItems := make(map[config.InputData]int)
	channelClose := make(chan bool)
	result := &FizzbuzzStat{mapStatItems: mapStatItems, channelInputStat: channelInputStat,
		channelCheckStat: channelCheckStat, channelClose: channelClose}
	go result.run()
	return result
}

func (fizzbuzzStat *FizzbuzzStat) run() {
	for {
		select {
		case <-fizzbuzzStat.channelClose:
			return
		case inputData := <-fizzbuzzStat.channelInputStat:
			nbHints := fizzbuzzStat.addNewInput(inputData)
			fizzbuzzStat.checkBestHint(inputData, nbHints)
		case checkChannel := <-fizzbuzzStat.channelCheckStat:
			checkChannel <- fizzbuzzStat.bestHint
		}
	}
}

func (fizzbuzzStat *FizzbuzzStat) addNewInput(inputData config.InputData) int {
	if statItemCount, ok := fizzbuzzStat.mapStatItems[inputData]; !ok {
		fizzbuzzStat.mapStatItems[inputData] = 1
		return 1
	} else {
		statItemCount++
		fizzbuzzStat.mapStatItems[inputData] = statItemCount
		return statItemCount
	}
}

func (fizzbuzzStat *FizzbuzzStat) checkBestHint(inputData config.InputData, nbHints int) {
	if fizzbuzzStat.bestHint.NbCalls < nbHints {
		fizzbuzzStat.bestHint.NbCalls = nbHints
		fizzbuzzStat.bestHint.Entry = inputData
	}
}

func (fizzbuzzStat *FizzbuzzStat) Stop() {
	close(fizzbuzzStat.channelClose)
}
