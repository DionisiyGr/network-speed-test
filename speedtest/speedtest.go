package speedtest

import (
	"fmt"

	"github.com/kylegrantlucas/speedtest"
)

//Result is struct for speedtest result
type Result struct {
	Latency  float64
	Upload   float64
	Download float64
}

//SpeedTest -
type SpeedTest struct{}

//New constructor to return new Speedtest instance
func New() *SpeedTest {
	return &SpeedTest{}
}

//Start start function to benchmark your network speed via speedtest
func (st *SpeedTest) Start() (*Result, error) {
	client, err := speedtest.NewDefaultClient()
	if err != nil {
		fmt.Printf("error creating client: %v", err)
		return nil, err
	}

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
