package service

import (
	"Faceit/src/internal/model"
	"Faceit/src/log"
	"errors"

	"github.com/streadway/amqp"
)

type PublisherConnection struct {
	Conn     *amqp.Connection
	Channel  *amqp.Channel
	exchange string
	queue    string
	err      chan error
}

func NewPublisherConnection(exchange string, queue string) *PublisherConnection {
	return &PublisherConnection{
		exchange: exchange,
		queue:    queue,
		err:      make(chan error),
	}
}

func (c *PublisherConnection) Connect() error {
	var err error
	c.Conn, err = amqp.Dial("amqp://guest:guest@rabbitmq/")
	if err != nil {
		log.Logger.Error().Msgf("Error connecting to rabbitMQ. Error: %s", err)
		return err
	}

	go func() {
		<-c.Conn.NotifyClose(make(chan *amqp.Error))
		c.err <- errors.New("PublisherConnection Closed")
	}()

	c.Channel, err = c.Conn.Channel()
	if err != nil {
		log.Logger.Error().Msgf("Channel error: %s", err)
		return err
	}

	if err := c.Channel.ExchangeDeclare(
		c.exchange, // name
		"fanout",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	); err != nil {
		log.Logger.Error().Msgf("Exchange error: %s", err)
		return err
	}

	return nil
}

// Save bindQueue function if we want to use this as a library and use it in the consumer
// func (c *PublisherConnection) BindQueue() error {
// 	if _, err := c.channel.QueueDeclare(c.queue, true, false, false, false, nil); err != nil {
// 		log.Logger.Error().Msgf("Error declaring the queue. Error: %s", err)
// 		return err
// 	}

// 	if err := c.channel.QueueBind(c.queue, "", c.exchange, false, nil); err != nil {
// 		log.Logger.Error().Msgf("Error binding the queue. Error: %s", err)
// 		return err

// 	}

// 	return nil
// }

//Reconnect reconnects the PublisherConnection
func (c *PublisherConnection) Reconnect() error {
	if err := c.Connect(); err != nil {
		return err
	}
	// if err := c.BindQueue(); err != nil {
	// 	return err
	// }
	return nil
}

func (c *PublisherConnection) Publish(m model.Message) error {
	select {
	case err := <-c.err:
		if err != nil {
			c.Reconnect()
		}
	default:
	}

	p := amqp.Publishing{
		ContentType: m.ContentType,
		Body:        m.Data,
	}
	if err := c.Channel.Publish(c.exchange, m.Queue, false, false, p); err != nil {
		log.Logger.Error().Msgf("Error Publishing. Error: %s", err)
		return err
	}

	return nil
}
