package benchmark

import "testing"

func BenchmarkPlusAppend(b *testing.B) {
	names := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	ret := ""
	for _, name := range names {
		ret += name
	}
	b.StopTimer()
}
