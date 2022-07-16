package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zlounes/fizzbuzz/config"
)

func TestStat(t *testing.T) {
	a := assert.New(t)
	channelInputStat := make(chan config.InputData)
	channelCheckStat := make(chan chan BestHint)
	stat := NewFizzbuzzStat(channelInputStat, channelCheckStat)
	channelInputStat <- config.NewInputData(1, 2, 3, "fizz", "buzz")
	channelInputStat <- config.NewInputData(4, 5, 6, "best", "hint")
	channelInputStat <- config.NewInputData(1, 2, 3, "fizz", "buzzy")
	channelInputStat <- config.NewInputData(1, 2, 3, "fizz", "buzzy2")
	channelInputStat <- config.NewInputData(4, 5, 6, "best", "hint")
	channelInputStat <- config.NewInputData(4, 5, 6, "best", "hint")
	channelInputStat <- config.NewInputData(1, 2, 3, "fizz", "buzz")
	chanDone := make(chan BestHint)
	channelCheckStat <- chanDone
	bestHint := <-chanDone
	a.Equal(config.NewInputData(4, 5, 6, "best", "hint"), bestHint.Entry, "not expected best hint retrieved")
	a.Equal(3, bestHint.NbCalls, "not expected number of hints")
	stat.Stop()
}
