package mqx

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
	"github.com/syzhang42/go-fire/auth"
	"github.com/syzhang42/go-fire/utils"
)

func TestPulsar(t *testing.T) {
	mm, err := NewMQManager(utils.PulsarAddr, utils.PulsarTopicx, utils.PulsarWorker)
	if err != nil {
		panic(err)
	}

	err = mm.NewProducers(utils.PulsarTopicx)
	if err != nil {
		panic(err)
	}

	err = mm.NewConsumers(utils.PulsarTopicx)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	ctxexit, cancelexit := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		mm.Recv(ctxexit, 1, utils.PulsarTopics[0], func(b Message) {
			if b != nil && b.Payload() != nil {
				fmt.Println("data:", string(b.Payload()), "idc:", b.Key())

			}
		})
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		go mm.Recv(ctxexit, 1, utils.PulsarTopics[1], func(b Message) {
			if b != nil && b.Payload() != nil {
				fmt.Println("data:", string(b.Payload()), "idc:", b.Key())

			}
		})
	}()

	for i := 0; i < 10; i++ {
		mm.Send(utils.PulsarTopics[0], context.Background(), []byte(fmt.Sprintf("%v:%v", utils.PulsarTopics[0], i)), &ProducerOption{
			Key:                "dx",
			DisableReplication: true,
		})
		mm.Send(utils.PulsarTopics[1], context.Background(), []byte(fmt.Sprintf("%v:%v", utils.PulsarTopics[1], i)), &ProducerOption{
			Key:                "dx",
			DisableReplication: true,
		})
	}
	time.Sleep(time.Second)
	cancelexit()
	wg.Wait()
}

func TestPulsarTxn(t *testing.T) {
	c, err := pulsar.NewClient(
		pulsar.ClientOptions{
			URL:                     utils.PulsarAddr,
			MaxConnectionsPerBroker: 10,
			Logger:                  log.NewLoggerWithLogrus(defaultLogger),
			EnableTransaction:       true,
		})
	auth.Must(err)

	consumer, err := c.Subscribe(pulsar.ConsumerOptions{
		Topic:                      utils.PulsarTopics[0],
		SubscriptionName:           "sub-" + utils.PulsarTopics[0],
		Type:                       pulsar.Shared,
		ReplicateSubscriptionState: true,
	})
	auth.Must(err)

	producer, err := c.CreateProducer(pulsar.ProducerOptions{
		Topic:                   utils.PulsarTopics[0],
		SendTimeout:             5 * time.Second,
		DisableBlockIfQueueFull: true,
		MaxPendingMessages:      10000,
		CompressionType:         pulsar.LZ4,
		CompressionLevel:        pulsar.Default,
	})
	auth.Must(err)

	txn, err := c.NewTransaction(time.Hour)
	auth.Must(err)

	for i := 0; i < 10; i++ {
		_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: make([]byte, 1024),
		})
		auth.Must(err)
	}
	for i := 0; i < 10; i++ {
		_, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
			Transaction: txn,
			Payload:     make([]byte, 1024),
		})
		auth.Must(err)
	}
	ctx, exitcancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("go done...")
				return
			default:
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
				defer cancel()
				msg, err := consumer.Receive(ctx)
				if err != nil {
					if strings.Contains(err.Error(), "No more data") {
						time.Sleep(time.Millisecond * 100)
					}
					continue
				} else {
					if msg != nil && msg.Payload() != nil {
						fmt.Println("recv", time.Now())
						consumer.Ack(msg)
					} else {
						continue
					}
				}
			}
		}
	}()
	time.Sleep(2 * time.Second)
	err = txn.Commit(context.Background())
	if err != nil {
		exitcancel()
		panic(err)
	}
	time.Sleep(2 * time.Second)
	exitcancel()
}
