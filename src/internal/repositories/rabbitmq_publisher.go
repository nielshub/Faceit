package repositories

import (
	"Faceit/src/internal/model"
	"Faceit/src/log"
	"errors"

	"github.com/streadway/amqp"
)

//Message is the amqp request to publish
type Message struct {
	Queue       string
	ContentType string
	Body        model.MessageBody
}

type Connection struct {
	name     string
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
	queue    string
	err      chan error
}

func NewConnection(name, exchange string, queue string) *Connection {
	return &Connection{
		name:     name,
		exchange: exchange,
		queue:    queue,
		err:      make(chan error),
	}
}

func (c *Connection) Connect() error {
	var err error
	c.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Logger.Error().Msgf("Error connecting to rabbitMQ. Error: %s", err)
		return err
	}
	defer c.conn.Close()

	go func() {
		<-c.conn.NotifyClose(make(chan *amqp.Error))
		c.err <- errors.New("Connection Closed")
	}()

	c.channel, err = c.conn.Channel()
	if err != nil {
		log.Logger.Error().Msgf("Channel error: %s", err)
		return err
	}
	defer c.channel.Close()

	if err := c.channel.ExchangeDeclare(
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
// func (c *Connection) BindQueue() error {
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

//Reconnect reconnects the connection
func (c *Connection) Reconnect() error {
	if err := c.Connect(); err != nil {
		return err
	}
	// if err := c.BindQueue(); err != nil {
	// 	return err
	// }
	return nil
}

func (c *Connection) Publish(m Message) error {
	select {
	case err := <-c.err:
		if err != nil {
			c.Reconnect()
		}
	default:
	}

	p := amqp.Publishing{
		Headers:     amqp.Table{"type": m.Body.Type},
		ContentType: m.ContentType,
		Body:        m.Body.Data,
	}
	if err := c.channel.Publish(c.exchange, m.Queue, false, false, p); err != nil {
		log.Logger.Error().Msgf("Error Publishing. Error: %s", err)
		return err
	}

	return nil
}
