package netflix

import (
	"fmt"

	"gopkg.in/ddo/go-fast.v0"
)

type NetflixFast struct {
}

func New() *NetflixFast {
	return &NetflixFast{}
}

type Result struct {
	Latency  float64
	Upload   float64
	Download float64
}

func (nf *NetflixFast) Start() (*Result, error) {
	fastCom := fast.New()

	// init
	err := fastCom.Init()
	if err != nil {
		return nil, err
	}

	// get urls
	urls, err := fastCom.GetUrls()
	if err != nil {
		return nil, err
	}

	// measure
	KbpsChan := make(chan float64)

	go func() {
		for Kbps := range KbpsChan {
			fmt.Printf("%.2f Kbps %.2f Mbps\n", Kbps, Kbps/1000)
		}

		fmt.Println("done")
	}()

	err = fastCom.Measure(urls, KbpsChan)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
