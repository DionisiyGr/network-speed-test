## Installation

```sh
go get -u github.com/DionisiyGr/network-speed-test
```

## Usage
The main function is 
```
GetNetworkSpeed
```
Options need string parameter - "speedtest" or "fast"

## Examples
```
import (
	"fmt"

	nst "github.com/DionisiyGr/network-speed-test"
)

func main() {
	test := nst.New()
	res, _ := test.GetNetworkSpeed(&nst.Options{"speedtest"}) // will benchmark your network speed via speedtest
	fmt.Println(res)
}
```
or
```
import (
	"fmt"

	nst "github.com/DionisiyGr/network-speed-test"
)

func main() {
	test := nst.New()
	res, _ := test.GetNetworkSpeed(&nst.Options{"fast"}) // will benchmark your network speed via fast
	fmt.Println(res)
}
```


## Benchmark tests
To run benchmark tests for Netflix you can type
```
go test -bench=. github.com/DionisiyGr/network-speed-test/netflix
```
and next should be in yout console
```
goos: linux
goarch: amd64
pkg: github.com/DionisiyGr/network-speed-test/netflix
BenchmarkNetflix-6             1        11900437746 ns/op
PASS
ok      github.com/DionisiyGr/network-speed-test/netflix        11.905s
```
-benchmem flag also can be provided


To run benchmark test for Speedtest you need ty type
```
go test -bench=. github.com/DionisiyGr/network-speed-test/speedtest
```
and the next will be shown on console:
```
goos: linux
goarch: amd64
pkg: github.com/DionisiyGr/network-speed-test/speedtest
BenchmarkSpeedtest-6           1        7519135536 ns/op
PASS
ok      github.com/DionisiyGr/network-speed-test/speedtest      7.528s
```

## Note
Result for GetNetworkSpeed is struct which contains 2 fields - Download speed and Upload speed (Fast supports just download) in kbps.
"Fast" service returns download speed for different urls, so resut would bu just average number.