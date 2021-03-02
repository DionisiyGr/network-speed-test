package networkspeedcheck

import (
	"errors"

	"github.com/DionisiyGr/network-speed-test/netflix"
	"github.com/DionisiyGr/network-speed-test/speedtest"
)

// Possible types
const (
	TypeNetflix   = "fast"
	TypeSpeedtest = "speedtest"
)

// Response is response of tested speed
type Response struct {
	Download float64
	Upload   float64
}

// Options struct which keeps type of check
type Options struct {
	CheckType string
}

//NetworkSpeedCheck -
type NetworkSpeedCheck struct {
	nfTest *netflix.NetflixFast
	spTest *speedtest.SpeedTest
}

//New return new instance of NetworkSpeedCheck
func New() *NetworkSpeedCheck {
	return &NetworkSpeedCheck{
		spTest: speedtest.New(),
		nfTest: netflix.New(),
	}
}

// GetNetworkSpeed process checking for provided option which define type (Speedtest or Netflix)
func (nsc *NetworkSpeedCheck) GetNetworkSpeed(op *Options) (*Response, error) {
	switch op.CheckType {
	case TypeSpeedtest:
		res, err := nsc.spTest.Start()
		return &Response{Download: res.Download, Upload: res.Upload}, err
	case TypeNetflix:
		res, err := nsc.nfTest.Start()
		return &Response{Download: res.Download}, err
	}
	return nil, errors.New("undefined speed check type")
}
