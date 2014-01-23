package httprequester

const (
    START = iota
    STOP
    REQUEST_SUCCESS
    REQUEST_ERROR
    INFO
)

type Message struct {
    Type int
    Payload string
}
