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

func TestGetAverage(t *testing.T) {
	tests := []struct {
		name  string
		isErr bool
		args  []float64
	}{
		{
			name:  "empty array",
			isErr: true,
			args:  nil,
		},
		{
			name:  "correct execution",
			isErr: false,
			args:  []float64{1.22, 17, 15.23, 165.2222},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getAverage(tt.args)
			if (res > 0) != !tt.isErr {
				t.Errorf("getAverage, res = %v wantErr %v", res, tt.isErr)
				return
			}
		})
	}
}
