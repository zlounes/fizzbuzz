package metrics

import (
	"github.com/zlounes/fizzbuzz/config"
)

const (
	HASH_SEPARATOR = ";"
)

type BestHint struct {
	Entry   config.FizzBuzzInput
	NbCalls int
}

type FizzbuzzStat struct {
	mapStatItems     map[config.FizzBuzzInput]int
	channelInputStat <-chan config.FizzBuzzInput
	channelCheckStat <-chan chan BestHint
	channelClose     chan bool
	bestHint         BestHint
}

func NewFizzbuzzStat(channelInputStat <-chan config.FizzBuzzInput, channelCheckStat <-chan chan BestHint) *FizzbuzzStat {
	mapStatItems := make(map[config.FizzBuzzInput]int)
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

func (fizzbuzzStat *FizzbuzzStat) addNewInput(inputData config.FizzBuzzInput) int {
	if statItemCount, ok := fizzbuzzStat.mapStatItems[inputData]; !ok {
		fizzbuzzStat.mapStatItems[inputData] = 1
		return 1
	} else {
		statItemCount++
		fizzbuzzStat.mapStatItems[inputData] = statItemCount
		return statItemCount
	}
}

func (fizzbuzzStat *FizzbuzzStat) checkBestHint(inputData config.FizzBuzzInput, nbHints int) {
	if fizzbuzzStat.bestHint.NbCalls < nbHints {
		fizzbuzzStat.bestHint.NbCalls = nbHints
		fizzbuzzStat.bestHint.Entry = inputData
	}
}

func (fizzbuzzStat *FizzbuzzStat) Stop() {
	close(fizzbuzzStat.channelClose)
}
