package speedtest

import (
	"fmt"

	"github.com/kylegrantlucas/speedtest"
)

type Result struct {
	Latency  float64
	Upload   float64
	Download float64
}

type SpeedTest struct{}

func New() *SpeedTest {
	return &SpeedTest{}
}

func (st *SpeedTest) Start() (*Result, error) {
	client, err := speedtest.NewDefaultClient()
	if err != nil {
		fmt.Printf("error creating client: %v", err)
		return nil, err
	}

	// Pass an empty string to select the fastest server
	server, err := client.GetServer("")
	if err != nil {
		fmt.Printf("error getting server: %v", err)
		return nil, err
	}

	dmbps, err := client.Download(server)
	if err != nil {
		fmt.Printf("error getting download: %v", err)
		return nil, err
	}

	umbps, err := client.Upload(server)
	if err != nil {
		fmt.Printf("error getting upload: %v", err)
	}

	return &Result{
		Latency:  server.Latency,
		Download: dmbps,
		Upload:   umbps,
	}, nil
}
