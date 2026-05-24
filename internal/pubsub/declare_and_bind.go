package pubsub

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type SimpleQueueType int

const (
	Durable   SimpleQueueType = iota
	Transient SimpleQueueType = iota
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // SimpleQueueType is an "enum" type I made to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {
	isDurable := queueType == Durable
	isTransient := queueType == Transient
	c, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	newQueue, err := c.QueueDeclare(
		queueName,   // name
		isDurable,   // durable
		isTransient, // delete when unused
		isTransient, // exclusive
		false,       // no-wait
		nil,         //args
	)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	err = c.QueueBind(queueName, key, exchange, false, nil)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	return c, newQueue, nil

}
