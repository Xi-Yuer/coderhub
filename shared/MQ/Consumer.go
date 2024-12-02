package rabbitmq

import (
	"context"
	"log"
	"time"
)

type Consumer struct {
	mq       *RabbitMQ
	handlers map[string]MessageHandler
	ctx      context.Context
	cancel   context.CancelFunc
}

type MessageHandler func(message []byte) error

func NewConsumer(mq *RabbitMQ) *Consumer {
	ctx, cancel := context.WithCancel(context.Background())
	return &Consumer{
		mq:       mq,
		handlers: make(map[string]MessageHandler),
		ctx:      ctx,
		cancel:   cancel,
	}
}

// RegisterHandler 注册消息处理器
func (c *Consumer) RegisterHandler(queueName string, handler MessageHandler) {
	c.handlers[queueName] = handler
}

// Start 启动消费者服务
func (c *Consumer) Start() error {
	for queueName, handler := range c.handlers {
		// 为每个队列启动一个独立的消费者goroutine
		go c.consume(queueName, handler)
	}
	return nil
}

func (c *Consumer) consume(queueName string, handler MessageHandler) {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			msgs, err := c.mq.channel.Consume(
				queueName,
				"",    // consumer
				false, // auto-ack
				false, // exclusive
				false, // no-local
				false, // no-wait
				nil,   // args
			)
			if err != nil {
				log.Printf("开始消费失败: %v", err)
				time.Sleep(time.Second * 5) // 重试间隔
				continue
			}

			for msg := range msgs {
				// 处理消息
				err := handler(msg.Body)
				if err != nil {
					log.Printf("处理消息失败: %v", err)
					_ = msg.Nack(false, true) // 消息处理失败，重新入队
				} else {
					_ = msg.Ack(false) // 确认消息
				}
			}
		}
	}
}

// Stop 停止消费者服务
func (c *Consumer) Stop() {
	c.cancel()
}
