// https://www.intel.com/content/www/us/en/embedded/training/ia-32-ia-64-benchmark-code-execution-paper.html

#include "textflag.h"

// BenchmarkStart() uint64
TEXT ·BenchmarkStart(SB), NOSPLIT, $0
	CPUID
	RDTSC
	SHLQ $32, DX
	ORQ  DX, AX
	MOVQ AX, r+0(FP)
	RET

// BenchmarkEnd() uint64
TEXT ·BenchmarkEnd(SB), NOSPLIT, $0
	RDTSCP
	SHLQ $32, DX
	ORQ  DX, AX
	MOVQ AX, r+0(FP)
	CPUID
	RET

// Rdtscp() uint32
TEXT ·Rdtscp(SB), NOSPLIT, $0
	RDTSCP
	SHLQ $32, DX
	ORQ  DX, AX
	MOVQ AX, r+0(FP)
	MOVQ CX, cpu+8(FP)
	RET

