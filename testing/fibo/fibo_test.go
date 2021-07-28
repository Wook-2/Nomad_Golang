package fibo_test

import (
	"testing"

	"github.com/j1mmyson/testing/fibo"
)

func TestFibo(t *testing.T) {
	ans := fibo.GetFibo(50)

	if ans != 12586269025 {
		t.Errorf("Fibo(50) = %d; want 1258626902\n", ans)
	}
}

func TestFibo2(t *testing.T) {
	ans := fibo.GetFibo2(50)

	if ans != 12586269025 {
		t.Errorf("Fibo2(50) = %d; want 12586269025\n", ans)
	}
}

func BenchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo.GetFibo(50)
	}
}

func BenchmarkFibo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo.GetFibo2(50)
	}
}
