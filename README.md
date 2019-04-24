The [Handshake Alliance](https://handshakealliance.org) API client for [hsd](https://github.com/handshake-org/hsd).

## Documentation

[https://godoc.org/github.com/HandshakeAlliance/hs-client-go/pkg](https://godoc.org/github.com/HandshakeAlliance/hs-client-go/pkg)

## Install

```bash
go get -u github.com/HandshakeAlliance/hs-client-go/pkg
```

## Getting started

```go
package main

import (
	"fmt"
	"log"

	hsc "github.com/HandshakeAlliance/hs-client-go/pkg/http"
)

func main() {

    //Your Handshake Node hostname
    hostname := "http://locahost:13037"
    hsClient := hsc.New(hostname)

    info, err := hsClient.Info()

	if err != nil {
		log.Fatal(err)
	}

    fmt.Printf("%v", info)

}

```

## Examples

Check out the [`./example`](./example) directory and documentation.

## License

MIT
