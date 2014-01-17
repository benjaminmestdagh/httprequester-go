package httprequester

import (
    "fmt"
    "net/http"
)

const (
    Version = "0.1"
)

type HttpRequester struct {
    Address string
    GetBody bool
    Requests int
    Threads int
}

func (h HttpRequester) Start() {
    client := &http.Client{}

    request, err := http.NewRequest(h.getMethod(), "http://" + h.Address, nil)
    if err != nil {
        fmt.Println(err)
    }

    userAgent := "HttpRequester/" + Version
    request.Header.Set("User-Agent", userAgent)

    requestsPerThread := h.Requests / h.Threads
    remainingRequests := h.Requests % h.Threads

    c := make(chan string)
//    h.run(c, requestsPerThread, remainingRequests, client, request)
    for i := 0; i < h.Threads; i++ {
        if(remainingRequests > 0) {
            startRoutine(c, requestsPerThread + 1, client, request)
            remainingRequests--
        } else {
            startRoutine(c, requestsPerThread, client, request)
        }
    }

    for i := 0; i < h.Requests; i++ {
        message := <-c
        fmt.Println(message)
    }
}

func (h HttpRequester) getMethod() string {
    if h.GetBody {
        return "GET"
    }

    return "HEAD"
}

func startRoutine(c chan string, requests int, client *http.Client, request *http.Request) {
    go func() {
        for i := 0; i < requests; i++ {
            var message string
            response, err := client.Do(request)
            if err != nil {
                message = fmt.Sprintf("%v", err)
            } else {
                message = fmt.Sprintf("%v %v", response.Proto, response.Status)
                defer response.Body.Close()
            }
            c <- message
        }
    }()
}
