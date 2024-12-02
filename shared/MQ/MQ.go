package MQ

import (
	"github.com/streadway/amqp"
)

type MQ struct {
	conn *amqp.Connection
}

func NewMQ() *MQ {
	return &MQ{}
}

// 连接到 RabbitMQ
func (m *MQ) Connect() error {
	conn, err := amqp.Dial("amqp://rabbitmq:2214380963Wx!!@localhost:5672/")
	if err != nil {
		return err
	}
	m.conn = conn
	return nil
}

