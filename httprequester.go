package httprequester

import (
    "fmt"
    "net/http"
)

const (
    Version = "0.5"
)

type HttpRequester struct {
    Host string
    GetBody bool
    Requests int
    Threads int
    Comchan chan Message
}

func (h HttpRequester) Start() {
    client := &http.Client{}
    request, err := http.NewRequest(h.getMethod(), "http://" + h.Host, nil)

    if err != nil {
        message := Message{ REQUEST_ERROR, fmt.Sprintf("%v", err) }
        h.Comchan <- message
    } else {
        requestsPerThread, remainingRequests := h.getThreadsAndRequests()
        c := make(chan Message)
        userAgent := "HttpRequester/" + Version
        request.Header.Set("User-Agent", userAgent)

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
            h.Comchan <- message
        }
    }


    message := Message{ STOP, "Stopped." }
    h.Comchan <- message
}

func (h HttpRequester) getMethod() string {
    if h.GetBody {
        return "GET"
    }

    return "HEAD"
}

func (h HttpRequester) getThreadsAndRequests() (int, int) {
    if h.Requests < h.Threads {
        h.Threads = h.Requests
        message := Message { INFO, fmt.Sprintf("Detected more threads than requests, presumed %v threads.", h.Threads) }
        h.Comchan <- message
    }

    requestsPerThread := h.Requests / h.Threads
    remainingRequests := h.Requests % h.Threads

    return requestsPerThread, remainingRequests
}

func startRoutine(c chan Message, requests int, client *http.Client, request *http.Request) {
    go func() {
        for i := 0; i < requests; i++ {
            response, err := client.Do(request)
            var message Message
            if err != nil {
                message = Message{ REQUEST_ERROR, fmt.Sprintf("%v", err) }
            } else {
                message = Message{ REQUEST_SUCCESS, fmt.Sprintf("%v %v", response.Proto, response.Status) }
                defer response.Body.Close()
            }
            c <- message
        }
    }()
}
