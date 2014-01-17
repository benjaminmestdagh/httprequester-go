package main

import (
    "fmt"
    "github.com/benjaminmestdagh/httprequester"
)

func main() {
    requester := &httprequester.HttpRequester{ "localhost", false, 100, 10 }
    fmt.Printf("Started HttpRequester version %v\n", httprequester.Version)
    requester.Start()
}
