# appsflyer
  
AppsFlyer API client library for Go.
  
https://support.appsflyer.com/hc/en-us/articles/207034346-Pull-APIs-Pulling-AppsFlyer-Reports-by-APIs
  
## Installation

```bash
$ go get github.com/F0urTw0/appsflyer
```

## RawData
  
```go
import (
	"fmt"
	"github.com/bcjti/appsflyer"
)

func main() {

	var (
		apiToken = "xxx-xxx-xxx"
		appID    = "xxx-xxx-xxx"
		fromDate = "2016-08-20"
		toDate   = "2016-08-21"
	)

	client := NewClientWithParam(apiToken, appID, fromDate, toDate)

	// export APPSFLYER_API_TOKEN=""
	// client := NewClient(appID, fromDate, toDate)
	
	reports, err := client.GetRawData(appsflyer.InstallsReport)
	if err != nil {
		panic(err)
	}

	fmt.Println(reports)

}
```

## Maps

```go
import (
	"fmt"
	"github.com/F0urTw0/appsflyer"
)

func main() {

	var (
		apiToken = "xxx-xxx-xxx"
		appID    = "xxx-xxx-xxx"
		fromDate = "2016-08-20"
		toDate   = "2016-08-21"
	)

	client := NewClientWithParam(apiToken, appID, fromDate, toDate)

	// export APPSFLYER_API_TOKEN=""
	// client := NewClient(appID, fromDate, toDate)

	reports, err :=  client.GetMaps(appsflyer.InstallsReport)
	if err != nil {
        panic(err)
    }

	fmt.Println(reports)

}
```
