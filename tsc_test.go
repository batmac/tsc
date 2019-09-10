package tsc

import (
	"fmt"
	"testing"
	"time"
)

var k uint64

func cpuN() int {
	var sum int
	for i := 0; i < 1000000; i++ {
		_, cpu := Rdtscp()
		sum += int(cpu)
		if cpu > 0 {
			//fmt.Printf("%v: %v cycles, cpu %v\n", i, endCounter-startCounter, cpu)
			break
		}
	}
	return int(sum)
}
func TestFunction(t *testing.T) {
	var startTime, endTime int64
	var startCounter, endCounter uint64

	period := 500 * time.Millisecond
	startTime = time.Now().UnixNano()
	startCounter = BenchmarkStart()
	time.Sleep(time.Duration(period))
	endCounter = BenchmarkEnd()
	endTime = time.Now().UnixNano()
	fmt.Printf("%v %v %v %v\n", startTime, endTime, startCounter, endCounter)
	elapsed := float64(endTime-startTime) / 1000000000
	fmt.Printf("%v s elapsed\n", elapsed)
	fmt.Printf("%v cycles\n", endCounter-startCounter)
	fmt.Printf("%.2e Hz\n", float64(endCounter-startCounter)/elapsed)
	fmt.Printf("%.2e HZ!\n", Frequency(10, time.Duration(10*time.Millisecond)))

	fmt.Println("overhead:")
	startCounter = BenchmarkStart()
	endCounter = BenchmarkEnd()
	fmt.Printf("%v cycles\n", endCounter-startCounter)
}

func TestRdtscp(t *testing.T) {

	t.Log("rdtsc:")

	c := make(chan int)
	var n int
	iterations := 100

	for i := 0; i < iterations; i++ {
		go func() {
			c <- cpuN()
		}()
	}
	for i := 0; i < iterations; i++ {
		n += <-c
	}
	if n == 0 {
		t.Log("Warninng: your OS doesn't set IA32_TSC_AUX")
	}
	t.Logf("cpu %d\n", n)
}

func BenchmarkOverhead1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k += BenchmarkStart()
		k += BenchmarkEnd()
	}
}

func BenchmarkOverhead2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k += BenchmarkStart()
		k += BenchmarkEnd()
	}
}

func BenchmarkOverhead3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k += BenchmarkStart()
		k += BenchmarkEnd()
	}
}

func BenchmarkOverhead4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k += BenchmarkStart()
		k += BenchmarkEnd()
	}
}

func BenchmarkOverhead5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k += BenchmarkStart()
		k += BenchmarkEnd()
	}
}
