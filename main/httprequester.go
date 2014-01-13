package main

import (
    "fmt"
    "github.com/benjaminmestdagh/httprequester"
)

func main() {
    fmt.Printf("Started HttpRequester version %v\n", httprequester.Version)
    httprequester.Start()
}
