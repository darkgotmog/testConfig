package message

type Message struct {
	Id        int64
	Data      []byte
	LenShared int
}

type RequestMessage struct {
	Message    *Message
	ResponseCh chan Response
}

type Response struct {
	Err error
}
