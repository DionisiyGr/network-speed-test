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

// Response -
type Response struct {
	Download float64
	Upload   float64
}

// Options -
type Options struct {
	CheckType string
}
type NetworkSpeedCheck struct {
	nfTest *netflix.NetflixFast
	spTest *speedtest.SpeedTest
}

func New() *NetworkSpeedCheck {
	return &NetworkSpeedCheck{
		spTest: speedtest.New(),
		nfTest: netflix.New(),
	}
}

// GetNetworkSpeed process checking for provided option which define type (Ookla or Netflix)
func (nsc *NetworkSpeedCheck) GetNetworkSpeed(op *Options) (*Response, error) {
	switch op.CheckType {
	case TypeSpeedtest:
		res, err := nsc.spTest.Start()
		return &Response{Download: res.Download, Upload: res.Upload}, err
	case TypeNetflix:
		_, err := nsc.nfTest.Start()
		return nil, err
		//return &Response{Download: fres.Download, Upload: fres.Upload}, err
	}
	return nil, errors.New("undefined speed check type")
}
