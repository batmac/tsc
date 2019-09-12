package tsc

import (
	"testing"
	"time"

	"github.com/montanaflynn/stats"
)

var (
	k       uint64
	l       int
	maxIter = 1000
)

// Test if we are able to get something != 0 in Rdtsc's cpu
func cpuN() int {
	var cpu uint64
	for i := 0; i < 1000000; i++ {
		_, cpu := Rdtscp()
		if cpu != 0 {
			break
		}
	}
	return int(cpu)
}

func TestRdtscp(t *testing.T) {
	c := make(chan int)
	var n int
	iterations := 100
	// we pop a lot of goroutines, expecting at least one will be scheduled on another core
	for i := 0; i < iterations; i++ {
		go func() {
			c <- cpuN()
		}()
	}
	for i := 0; i < iterations; i++ {
		n += <-c
	}
	if n == 0 {
		t.Log("Warning: your OS doesn't seem to set IA32_TSC_AUX")
	}
}

// this "indirection" instead of inline code costs a few ops (<5 according to my tests)
func testGenericOverhead(t *testing.T, f func()) {
	s := make([]float64, maxIter)
	for i := 0; i < maxIter; i++ {
		m := BenchmarkStart()
		f()
		m = BenchmarkEnd() - m
		s[i] = float64(m)
	}
	printAndCheckSerieStats(t, s)
}
func TestBenchmarkInnerOverheadInline(t *testing.T) {
	s := make([]float64, maxIter)
	for i := 0; i < maxIter; i++ {
		m := BenchmarkStart()
		m = BenchmarkEnd() - m
		s[i] = float64(m)
	}
	printAndCheckSerieStats(t, s)
}
func printAndCheckSerieStats(t *testing.T, s []float64) {
	r, _ := stats.Mean(s)
	t.Logf("mean: %.2f\n", r)
	r, _ = stats.Median(s)
	t.Logf("median: %v\n", r)
	r, _ = stats.Midhinge(s)
	t.Logf("midhinge: %v\n", r)
	r, _ = stats.Trimean(s)
	t.Logf("trimean: %.2f\n", r)
	r, _ = stats.StdDevS(s)
	t.Logf("stdevs: %.2f\n", r)
	r, _ = stats.InterQuartileRange(s)
	t.Logf("interquartile range: %v\n", r)
	r, _ = stats.Percentile(s, 99)
	t.Logf("99%%: %v\n", r)
	r, _ = stats.Percentile(s, 95)
	t.Logf("95%%: %v\n", r)
	r, _ = stats.Percentile(s, 90)
	t.Logf("90%%: %v\n", r)
	r, _ = stats.Percentile(s, 75)
	t.Logf("75%%: %v\n", r)
	r, _ = stats.Max(s)
	t.Logf("max: %v\n", r)
	r, _ = stats.Min(s)
	t.Logf("min: %v\n", r)
	if r < 0 {
		t.Error("negative value found!")
	}
	// graph := asciigraph.Plot(s)
	// fmt.Println(graph)

	//h := thist.NewHist(s, "Example histogram", "auto", -1, true)
	//fmt.Println(h.Draw())
}
func TestBenchmarkInnerOverhead(t *testing.T) {
	testGenericOverhead(t, func() {})
}
func TestBenchmarkOuterOverhead(t *testing.T) {
	testGenericOverhead(t, func() {
		k = BenchmarkStart()
		k = BenchmarkEnd() - k
	})
}
func TestRdtscpOverhead(t *testing.T) {
	testGenericOverhead(t, func() {
		k, _ = Rdtscp()
	})
}
func TestTicksOverhead(t *testing.T) {
	testGenericOverhead(t, func() {
		k = Ticks()
	})
}
func TestSinceOverhead(t *testing.T) {
	testGenericOverhead(t, func() {
		k = Ticks()
		k = Since(k)
	})
}
func TestCPUIDOverhead(t *testing.T) {
	testGenericOverhead(t, func() {
		Cpuid()
	})
}
func TestFrequency(t *testing.T) {
	s := make([]float64, maxIter)
	for i := 0; i < maxIter; i++ {
		f := Frequency(1, 1*time.Millisecond)
		s[i] = f
		// t.Logf("TSC freq found: %.2e Hz", f)

		if f <= 0 {
			t.Fail()
		}
	}
	min, _ := stats.Min(s)
	max, _ := stats.Max(s)
	minmax := (min / max) * 100
	t.Logf("min/max: %.2f%%", minmax)
	if minmax < 90 {
		t.Log("Warning: your TSC doesn't look stable or sync")
	}
	printAndCheckSerieStats(t, s)
}
