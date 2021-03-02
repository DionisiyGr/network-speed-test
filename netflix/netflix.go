package netflix

import (
	"gopkg.in/ddo/go-fast.v0"
)

//NetflixFast -
type NetflixFast struct{}

//New is constructor to return NetflixFast instance
func New() *NetflixFast {
	return &NetflixFast{}
}

//Result will return result after benchmarking yourr speed.
//NOTE: Fast doesnt return upload speed
type Result struct {
	Download float64
}

//Start start function to benchmark your network speed via netflix`s Fast
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
	var res []float64
	go func() {
		for kbps := range KbpsChan {
			res = append(res, kbps)
		}
	}()

	err = fastCom.Measure(urls, KbpsChan)
	if err != nil {
		return nil, err
	}

	return &Result{Download: getAverage(res)}, nil
}

func getAverage(in []float64) float64 {
	if in != nil {
		var sum float64
		for _, v := range in {
			sum = sum + v
		}
		return sum / float64(len(in))
	}
	return 0
}
