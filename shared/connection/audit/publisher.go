package audit

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

var (
	ErrMessageBrokerNotInitialized = errors.New("message broker is not initialized")
	ErrConnectionNotInitialized    = errors.New("connection is not initialized")
	ErrRegisterExchange            = errors.New("failed to register an exchange")
	ErrOpenChannel                 = errors.New("failed to open a channel")
	ErrRegisterQueue               = errors.New("failed to register an queue")
	ErrBindingQueue                = errors.New("failed to binding queue")
	ErrRegisterConsumer            = errors.New("failed to register a consumer")
)

type EventHandler func(delivery amqp.Delivery)

type Publisher struct {
	conn *amqp.Connection
}

func NewPublisher(login, password, host string) (*Publisher, error) {
	var err error

	broker := &Publisher{}

	if broker.conn, err = amqp.Dial(fmt.Sprintf("amqp://"+login+":"+password+"@%s/", host)); err != nil {
		return nil, fmt.Errorf("failed to connect AMQP broker at: %s", host)
	}
	return broker, nil
}

func (m *Publisher) Produce(exchangeName, queueName, bindingKey string, body []byte) error {
	if m == nil {
		return ErrMessageBrokerNotInitialized
	}

	if m.conn == nil {
		return ErrConnectionNotInitialized
	}

	ch, err := m.conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		return ErrRegisterExchange
	}

	queue, err := ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)

	if err != nil {
		return ErrRegisterQueue
	}

	err = ch.QueueBind(
		queue.Name,   // name of the queue
		bindingKey,   // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)

	if err != nil {
		return ErrBindingQueue
	}

	err = ch.Publish( // Publishes a message onto the queue.
		exchangeName, // exchange
		bindingKey,   // routing key      q.Name
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body, // Our JSON body as []byte
		})

	return err
}

func (m *Publisher) Consume(queueName, consumerName string, eventHandler EventHandler) error {
	if m == nil {
		return ErrMessageBrokerNotInitialized
	}
	ch, err := m.conn.Channel()
	if err != nil {
		return ErrOpenChannel
	}

	msgs, err := ch.Consume(
		queueName,    // queue
		consumerName, // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		return ErrRegisterConsumer
	}

	go consumeLoop(msgs, eventHandler)

	return nil
}

func (m *Publisher) Close() {
	if m != nil {
		if m.conn != nil {
			m.conn.Close()
		}
	}
}

func consumeLoop(deliveries <-chan amqp.Delivery, eventHandler EventHandler) {
	for d := range deliveries {
		// Invoke the handlerFunc func we passed as parameter.
		eventHandler(d)
		//d.Ack(false)
	}
}
