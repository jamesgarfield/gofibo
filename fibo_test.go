package main

import (
	"runtime"
	"testing"
)

func TestGetFibo(t *testing.T) {
	n := GetFib(0)
	if n != 0 {
		t.Errorf("Should return 0, was %d", n)
	}

	n = GetFib(1)
	if n != 1 {
		t.Errorf("Should return 1, was %d", n)
	}

	n = GetFib(2)
	if n != 1 {
		t.Errorf("Should return 1, was %d", n)
	}

	n = GetFib(10)
	if n != 55 {
		t.Errorf("Should return 55, was %d", n)
	}

}

func BenchmarkGetFib(b *testing.B) {
	GetFib(25)
}

func BenchmarkGetFibWithProcs(b *testing.B) {
	runtime.GOMAXPROCS(4)
	GetFib(25)
}
