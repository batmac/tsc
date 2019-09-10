package tsc

import (
	"time"
)

func BenchmarkStart() uint64
func BenchmarkEnd() uint64
func Rdtscp() (uint64, uint32)

// will take n*period to find out
func Frequency(n int, period time.Duration) float64 {
	var startTime, endTime int64
	var startCounter, endCounter uint64
	freq := 9.9e9
	for i := 0; i < n; i++ {
		startTime = time.Now().UnixNano()
		startCounter = BenchmarkStart()
		time.Sleep(period)
		endCounter = BenchmarkEnd()
		endTime = time.Now().UnixNano()
		elapsed := float64(endTime-startTime) / 1000000000
		iterfreq := float64(endCounter-startCounter) / elapsed
		if iterfreq < freq {
			freq = iterfreq
		}
	}
	return freq
}
