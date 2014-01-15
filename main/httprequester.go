package main

import (
    "fmt"
    "github.com/benjaminmestdagh/httprequester"
)

func main() {
    requester := &httprequester.HttpRequester{ "localhost", false }
    fmt.Printf("Started HttpRequester version %v\n", httprequester.Version)
    requester.Start()
}
