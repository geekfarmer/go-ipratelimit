# Go rate limiter for HTTP requests based on IP address

This package provides a simple and easy to implement middleware for HTTP requests to rate limit the requests for each ip-address

```go
imports "github.com/geekfarmer/ipratelimit"
```

Create a ip address rate limiter with a maximum number of requests can be process per second and with a bucket size.

```go
var limiter = ipratelimit.New(1, 7)
```
Call IPRateLimitMiddleware() with all endpoints to process request. IPRateLimitMiddleware process request until it reached maximum requests

## How to use?

```go
import (
	"log"
  "net/http"

  "github.com/geekfarmer/ipratelimit"
)

// ipratelimit.New(r, b)
// r -> max events/second
// b -> max bucket size
var limiter = ipratelimit.New(1, 7)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	if err := http.ListenAndServe(":8888", limiter.IPRateLimitMiddleware(mux)); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	// any very expensive database call
	w.Write([]byte("alles gut"))
}
```


### Test HTTP load with [vegeta](https://github.com/tsenart/vegeta)

you can get all files in ./example folder

```
cd example
go get github.com/geekfarmer/ipratelimit
go build -o server .
./server
```

We can use vegaeta for HTTP load testing

```
brew install vegeta
```

we need to create a simple config file ( /example/request.config ) which have config for what requests do we want to produce

```
vegeta attack -duration=10s -rate=100 -targets=/PATH_TO/request.conf | vegeta report
```

