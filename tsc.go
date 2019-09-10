package tsc

import (
	"time"
)

func BenchmarkStart() uint64
func BenchmarkEnd() uint64
func Rdtscp() (uint64, uint32)

// will take n*period to find out
func Frequency(n int, period time.Duration) float64 {
	var startTime, endTime time.Time
	var startCounter, endCounter uint64
	freq := 9.9e9
	for i := 0; i < n; i++ {
		startTime = time.Now()
		startCounter = BenchmarkStart()
		time.Sleep(period)
		endCounter = BenchmarkEnd()
		endTime = time.Now()
		elapsed := float64(endTime.UnixNano()-startTime.UnixNano()) / 1000000000
		iterfreq := float64(endCounter-startCounter) / elapsed
		if iterfreq < freq {
			freq = iterfreq
		}
	}
	return freq
}
