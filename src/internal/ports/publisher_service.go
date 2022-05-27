package ports

import "Faceit/src/internal/model"

type PublisherService interface {
	Connect() error
	Reconnect() error
	Publish(model.Message) error
}
