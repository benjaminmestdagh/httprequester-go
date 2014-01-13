package httprequester

import (
    "fmt"
    "sync"
    "net/http"
)

const (
    Version = "0.1"
)

func Start() {
    client := &http.Client{}

    request, err := http.NewRequest("GET", "http://localhost", nil)
    if err != nil {
        fmt.Println(err)
    }

    userAgent := "HttpRequester/" + Version
    request.Header.Set("User-Agent", userAgent)

    waitGroup := sync.WaitGroup{}
    defer waitGroup.Wait()

    for i := 0; i < 10; i++ {
        waitGroup.Add(1)
        go func() {
            defer waitGroup.Done()
            response, err := client.Do(request)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Printf("%v %v\n", response.Proto, response.Status)
                defer response.Body.Close()
            }
        }()
    }

}
