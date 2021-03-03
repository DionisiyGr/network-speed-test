package netflix

import (
	"testing"
)

func BenchmarkNetflix(b *testing.B) {
	nf := New()
	for i := 0; i < b.N; i++ {
		nf.Start()
	}
}
