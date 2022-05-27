package model

//Struct for the body passed in the publisher message. The type will be set on the Request header
type MessageBody struct {
	Data []byte
	Type string
}
