package main

import (
    "fmt"
    "flag"
    "os"
    "github.com/benjaminmestdagh/httprequester"
)

var (
    getBody bool
    threads int
    requests int
    sleep int
)

func init() {
    flag.BoolVar(&getBody, "get-body", false, "Whether to do a GET request or not")
    flag.IntVar(&threads, "threads", 1, "The number of threads to use")
    flag.IntVar(&requests, "requests", 1, "The number of requests to do")
    flag.IntVar(&sleep, "sleep", 0, "The number of milliseconds to sleep between requests")
}

func main() {
    flag.Parse()
    switch {
    case len(flag.Args()) == 0:
        fmt.Println("Error: no host given.")
        os.Exit(1000)
    case requests <= 0:
        fmt.Printf("Error: cannot send %v requests.\n", requests)
        os.Exit(1001)
    case threads <= 0:
        fmt.Printf("Error: cannot work with %v threads.\n", requests)
        os.Exit(1002)
    }

    host := flag.Args()[0]
    c := make(chan httprequester.Message)
    requester := &httprequester.HttpRequester{ host, getBody, requests, threads, sleep, c }
    fmt.Printf("Started HttpRequester version %v\n", httprequester.Version)
    go requester.Start()
    stop := false
    for !stop {
        message := <-c
        stop = message.Type == httprequester.STOP
        fmt.Println(message.Payload)
    }
}
