package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	config  *Config
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

// NewRabbitMQ 创建RabbitMQ实例
func NewRabbitMQ(config *Config) (*RabbitMQ, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
	)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
		config:  config,
	}, nil
}
// PublishMessage 发布消息
func (r *RabbitMQ) PublishMessage(exchange, routingKey string, body []byte) error {
	return r.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// ConsumeMessage 消费消息
func (r *RabbitMQ) ConsumeMessage(queue string, handler func([]byte) error) error {
	msgs, err := r.channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			if err := handler(msg.Body); err != nil {
				// 处理错误，可以记录日志
				log.Printf("处理消息失败: %v", err)
			}
		}
	}()

	return nil
}

// DeclareQueue 声明队列
func (r *RabbitMQ) DeclareQueue(name string) error {
	_, err := r.channel.QueueDeclare(
		name,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	return err
}

// Close 关闭连接
func (r *RabbitMQ) Close() {
	if r.channel != nil {
		_ = r.channel.Close()
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
}
