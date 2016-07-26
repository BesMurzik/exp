package echo

import "testing"

var s = []string{"one", "two", "three", "four", "five"}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(s)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(s)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(s)
	}
}
