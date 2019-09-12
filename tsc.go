package tsc

//go:generate asmfmt -w tsc_amd64.s

import (
	"time"
)

// BenchmarkStart start a benchmark
func BenchmarkStart() uint64

// BenchmarkEnd end a benchmark
func BenchmarkEnd() uint64

// Rdtscp returns TSC and core ID if your OS support that (for instance, linux does, darwin dos not)
func Rdtscp() (ret uint64, cpu uint32)

// Ticks runs Rdtscp but only return TSC, so one less write (probably useless)
func Ticks() uint64

// Cpuid runs CPUID, mainly useful to benchmark it.
func Cpuid()

// Frequency returns your TSC frequency
// will take n*period to find out
func Frequency(n int, period time.Duration) float64 {
	var startTime, endTime time.Time
	var startCounter, endCounter uint64
	freq := 0.0
	for i := 0; i < n; i++ {
		startTime = time.Now()
		startCounter = BenchmarkStart()
		time.Sleep(period)
		endCounter = BenchmarkEnd()
		endTime = time.Now()
		elapsed := float64(endTime.UnixNano()-startTime.UnixNano()) / 1000000000
		iterfreq := float64(endCounter-startCounter) / elapsed
		if iterfreq > freq {
			freq = iterfreq
		}
	}
	return freq
}

// Since returns the ticks diff
func Since(start uint64) uint64 {
	return Ticks() - start
}
