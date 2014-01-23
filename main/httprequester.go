package main

import (
    "fmt"
    "flag"
    "os"
    "github.com/benjaminmestdagh/httprequester"
)

var (
    host string
    getBody bool
    threads int
    requests int
    sleep int
)

func init() {
    flag.StringVar(&host, "host", "", "The host to send requests to")
    flag.BoolVar(&getBody, "get-body", false, "Whether to do a GET request or not")
    flag.IntVar(&threads, "threads", 1, "The number of threads to use")
    flag.IntVar(&requests, "requests", 1, "The number of requests to do")
    flag.IntVar(&sleep, "sleep", 0, "The number of seconds to sleep between requests")
}

func main() {
    flag.Parse()
    switch {
    case host == "":
        fmt.Println("Error: no host given.")
        os.Exit(1000)
    case requests == 0:
        fmt.Println("Error: cannot send zero requests.")
        os.Exit(1001)
    case threads == 0:
        fmt.Println("Error: cannot work with zero threads.")
        os.Exit(1002)
    }

    c := make(chan httprequester.Message)
    requester := &httprequester.HttpRequester{ host, getBody, requests, threads, c }
    fmt.Printf("Started HttpRequester version %v\n", httprequester.Version)
    go requester.Start()
    stop := false
    for !stop {
        message := <-c
        stop = message.Type == httprequester.STOP
        fmt.Println(message.Payload)
    }
}
