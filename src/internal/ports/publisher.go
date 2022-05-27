package ports

import "Faceit/src/internal/model"

type Publisher interface {
	Connect() error
	Reconnect() error
	Publish(model.MessageBody) error
}
