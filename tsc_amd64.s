// https://www.intel.com/content/www/us/en/embedded/training/ia-32-ia-64-benchmark-code-execution-paper.html for the CPUID method.

// the calls to CPUID are very slow (2k-3k cycles), although it 
// is the proper way to do to benchmark code between 
// BenchmarkStart and BenchmarkEnd.
// Using RDTSCP (with Ticks()) is enough if you don't want to 
// slowdown as hell the rest of your code.

#include "textflag.h"

// func BenchmarkStart() uint64
TEXT ·BenchmarkStart(SB), NOSPLIT, $0
	CPUID
	RDTSC
	SHLQ $32, DX
	ORQ  DX, AX
	MOVQ AX, ret+0(FP)
	RET

// func BenchmarkEnd() uint64
TEXT ·BenchmarkEnd(SB), NOSPLIT, $0
	RDTSCP
	SHLQ $32, DX
	ORQ  DX, AX
	MOVQ AX, ret+0(FP)
	CPUID
	RET

// func Rdtscp() (ret uint64, cpu uint32)
TEXT ·Rdtscp(SB), NOSPLIT, $0
	RDTSCP
	SHLQ $32, DX
	ORQ  DX, AX
	MOVQ AX, ret+0(FP)
	MOVL CX, cpu+8(FP)
	RET

// func Ticks() uint64
TEXT ·Ticks(SB), NOSPLIT, $0
	RDTSCP
	SHLQ $32, DX
	ORQ  DX, AX
	MOVQ AX, ret+0(FP)
	RET

// func Cpuid()
// use to test CPUID overhead
TEXT ·Cpuid(SB), NOSPLIT, $0
    CPUID
	RET

