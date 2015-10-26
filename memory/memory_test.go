package memory

import (
	"testing"
)

func TestDefaultMemorySize(t *testing.T) {
	mem := <-Get
	if len(mem) != defaultMemorySize {
		t.Error("Default memory size should be 100")
	}
}

func TestDefaultQueueSize(t *testing.T) {
	for i := 0; i < channelSize; i++ {
		<-Get
	}

	select {
	case <-Get:
		t.Error("Should block on max channel size")
	default:
	}
}

func BenchmarkPerformance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mem := <-Get
		Give <- mem
	}
}

func BenchmarkPerformanceByteBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mem := <-GetBuffer
		GiveBuffer <- mem
	}
}

func BenchmarkAllocsSlice(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mem := <-Get
		copy(mem, "PONG")
		Give <- mem
	}
}

func BenchmarkAllocsByteBuffer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mem := <-GetBuffer
		mem.WriteString("PONG")
		GiveBuffer <- mem
	}
}
