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

## Note
Result for GetNetworkSpeed is struct which contains 2 fields - Download speed and Upload speed (Fast supports just download) in kbps.
"Fast" service returns download speed for different urls, so resut would bu just average number.