package speedtest

import "testing"

func BenchmarkSpeedtest(b *testing.B) {
	st := New()
	for i := 0; i < b.N; i++ {
		st.Start()
	}
}
