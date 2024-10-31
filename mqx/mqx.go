package mqx

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/syzhang42/go-fire/stringsx"
)

type ProducerOption struct {
	Key                string // 生产者根据 key 区分消息来源
	DisableReplication bool   // Disable the replication for this message
}

type Message interface {
	Payload() []byte // 消息内容
	Topic() string   // 消息的 topic
	Key() string     // 对应 ProducerOption 的 Key
}

type Producer interface {
	// for PulsarMQ, interface{} -> MessageID
	// for RabbitMQ, interface{} is invalid
	Send(context.Context, []byte, *ProducerOption) (interface{}, error)
	Close()
}

type Consumer interface {
	Recv(context.Context) (Message, error)
	Ack(Message) error
	Nack(Message) error
	Close()
}

type Client interface {
	NewProducer(topic string) (Producer, error)
	NewConsumer(topic string) (Consumer, error)
	Close()
}

type MQManager struct {
	addr      string
	clients   map[string][]Client
	consumers map[string][]Consumer
	producers map[string][]Producer
}

// 每个client 一个消费者或者一个生产者    每个topic 每个worker对应一个client
func NewMQManager(addr string, topicx string, worker int) (*MQManager, error) {
	topics := stringsx.Split(topicx, ",")
	if len(topics) == 0 {
		return nil, ErrInvalidTopic
	}
	if worker == 0 {
		return nil, ErrInvalidWorker
	}
	mm := &MQManager{
		addr:      addr,
		clients:   make(map[string][]Client),
		consumers: make(map[string][]Consumer),
		producers: make(map[string][]Producer),
	}
	//初始化 worker * len(topic)个客户端
	if strings.HasPrefix(addr, "pulsar") {
		for _, topic := range topics {
			for i := 0; i < worker; i++ {
				if c, err := NewPulsarClient(addr); err != nil {
					return nil, err
				} else {
					if mm.clients[topic] == nil {
						mm.clients[topic] = make([]Client, 0, worker)
					}
					mm.clients[topic] = append(mm.clients[topic], c)
				}
			}
		}
		return mm, nil
	}
	return nil, ErrInvalidUrl
}

// 每个topic创建 worker个生产者，所有的topic必须包含在NewMQManager 的topicx 之内
func (mm *MQManager) NewProducers(topicx string) error {
	topics := stringsx.Split(topicx, ",")
	if len(topics) == 0 {
		return ErrInvalidTopic
	}
	//每个client创建一个producer
	for _, topic := range topics {
		if topicClients, ok := mm.clients[topic]; ok {
			for num, topicClient := range topicClients {
				producer, err := topicClient.NewProducer(topic)
				if err != nil {
					fmt.Println("create producer failed:", topic, "num:", num, "err:", err)
					return err
				}
				if mm.producers[topic] == nil {
					mm.producers[topic] = make([]Producer, 0, len(topicClients))
				}
				mm.producers[topic] = append(mm.producers[topic], producer)
				fmt.Println("create producer successful:", topic, "num:", num)
			}
		} else {
			fmt.Println("create producer failed: not new client:", topic)
			return ErrInvalidTopic
		}
	}
	return nil
}

// 每个topic创建 worker个消费者，所有的topic必须包含在NewMQManager 的topicx 之内
func (mm *MQManager) NewConsumers(topicx string) error {
	topics := stringsx.Split(topicx, ",")
	if len(topics) == 0 {
		return ErrInvalidTopic
	}
	//每个client创建一个producer
	for _, topic := range topics {
		if topicClients, ok := mm.clients[topic]; ok {
			for num, topicClient := range topicClients {
				consumer, err := topicClient.NewConsumer(topic)
				if err != nil {
					fmt.Println("create consumer failed:", topic, "num:", num, "err:", err)
					return err
				}
				if mm.consumers[topic] == nil {
					mm.consumers[topic] = make([]Consumer, 0, len(topicClients))
				}
				mm.consumers[topic] = append(mm.consumers[topic], consumer)
				fmt.Println("create consumer successful:", topic, "num:", num)
			}
		} else {
			fmt.Println("create consumer failed: not new client:", topic)
			return ErrInvalidTopic
		}
	}
	return nil
}

func (mm *MQManager) Send(topic string, ctx context.Context, data []byte, opt *ProducerOption) (err error) {
	if producers, ok := mm.producers[topic]; ok && len(producers) > 0 {
		for _, producer := range producers {
			if _, err = producer.Send(ctx, data, opt); err == nil {
				break
			} else {
				continue
			}
		}
		return
	}
	//足够的连接，但是没有足够的producer,创建producer重试
	if clients, ok := mm.clients[topic]; ok && len(clients) > 0 {
		err = mm.NewProducers(topic)
		if err != nil {
			return err
		}
		return mm.Send(topic, ctx, data, opt)
	}
	return ErrInvalidTopic
}

// exitCtx：优雅退出上下文，在你想要结束监听的地方cancel
// Millisecond
func (mm *MQManager) Recv(exitCtx context.Context, timeOut int64, topic string, msgcb func(Message)) {
	var wg sync.WaitGroup
	if consumers, ok := mm.consumers[topic]; ok && len(consumers) > 0 {
		for index, consumer := range consumers {
			wg.Add(1)
			consumer := consumer
			go func(index int) {
				defer wg.Done()
				for {
					select {
					case <-exitCtx.Done():
						fmt.Println(topic, index, ": recv go exit...")
						return
					default:
						ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Millisecond)
						// defer cancel()
						msg, err := consumer.Recv(ctx)
						if err != nil {
							if strings.Contains(err.Error(), "No more data") {
								time.Sleep(time.Millisecond * 100)
							}
							continue
						} else {
							resp, _ := msg.(pulsar.Message)
							if resp != nil && resp.Payload() != nil {
								msgcb(msg)
								consumer.Ack(msg)
							} else {
								continue
							}
						}
					}
				}
			}(index)
		}
		//足够的连接，但是没有足够的consumer,创建consume重试
	} else if clients, ok := mm.clients[topic]; ok && len(clients) > 0 {
		err := mm.NewConsumers(topic)
		if err != nil {
			fmt.Println(topic, ":recv failed,err:", err)
			return
		}
		mm.Recv(exitCtx, timeOut, topic, msgcb)
	} else {
		fmt.Println(topic, ":recv failed,err: len(consumer)==0")
	}

	wg.Wait()
}
