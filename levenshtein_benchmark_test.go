package levenshtein

import "testing"

var benchmarks = []string{"a", "aa", "ab", "aba", "abba", "abbba", "bba"}

const fromBenchmark = "abba"

func BenchmarkDist(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, bench := range benchmarks {
			Dist(fromBenchmark, bench)
		}
	}
}

func benchmarkDistMax(max int, b *testing.B) {
	b.StopTimer()
	bytesBenchmarks := [][]byte{}
	for _, bench := range benchmarks {
		bytesBenchmarks = append(bytesBenchmarks, []byte(bench))
	}
	b.StartTimer()
	from := FromBytes([]byte(fromBenchmark))
	for n := 0; n < b.N; n++ {
		for _, bench := range bytesBenchmarks {
			from.Dist(bench, max)
		}
	}
}

func BenchmarkDistMax2(b *testing.B) {
	benchmarkDistMax(2, b)
}

func BenchmarkDistMax3(b *testing.B) {
	benchmarkDistMax(3, b)
}

func BenchmarkDistMax4(b *testing.B) {
	benchmarkDistMax(4, b)
}
