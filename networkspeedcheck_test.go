package networkspeedcheck

import "testing"

func TestGetNetworkSpeed(t *testing.T) {
	nsc := New()
	tests := []struct {
		name  string
		args  string
		isErr bool
	}{
		{
			name:  "bad check type",
			args:  "test",
			isErr: true,
		},
		{
			name:  "empty check type",
			args:  `""`,
			isErr: true,
		},
		{
			name:  "speedtest type",
			args:  "speedtest",
			isErr: false,
		},
		{
			name:  "netflix type",
			args:  "fast",
			isErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := nsc.GetNetworkSpeed(&Options{tt.args})
			if (err != nil) != tt.isErr {
				t.Errorf("Client.GetServer() error = %v, wantErr %v", err, tt.isErr)
				return
			}

		})
	}
}
