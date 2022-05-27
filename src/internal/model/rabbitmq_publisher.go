package model

type Message struct {
	Queue       string
	ContentType string
	Data        []byte
}
