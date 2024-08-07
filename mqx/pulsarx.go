package mqx

import (
	"context"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var (
	defaultLogger *logrus.Logger
)

func init() {
	os.MkdirAll("/log/server", os.ModePerm)
	defaultLogger = logrus.New()
	rl, _ := rotatelogs.New( // 默认 rotate 时间1天，最长暴露时间 7 day
		"/log/server/pulsar.log.%Y%m%d%H%M",
		rotatelogs.WithLinkName("/log/server/pulsar"),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithMaxAge(time.Hour*24*7),
	)
	defaultLogger.SetOutput(rl)
}

type PulsarProducer struct {
	producer pulsar.Producer
}

func (pp *PulsarProducer) Send(ctx context.Context, payload []byte, opt *ProducerOption) (interface{}, error) {
	msg := &pulsar.ProducerMessage{
		Payload: payload,
	}

	if opt != nil {
		msg.Key = opt.Key
		msg.DisableReplication = opt.DisableReplication
	}

	return pp.producer.Send(ctx, msg)
}

func (pp *PulsarProducer) Close() {
	pp.producer.Close()
}

type PulsarConsumer struct {
	consumer pulsar.Consumer
}

func (pc *PulsarConsumer) Recv(ctx context.Context) (Message, error) {
	msg, err := pc.consumer.Receive(ctx)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (pc *PulsarConsumer) Ack(msg Message) error {
	pc.consumer.Ack(msg.(pulsar.Message))
	return nil
}

func (pc *PulsarConsumer) Nack(msg Message) error {
	pc.consumer.Nack(msg.(pulsar.Message))
	return nil
}

func (pc *PulsarConsumer) Close() {
	pc.consumer.Close()
}

type PulsarClient struct {
	url    string
	client pulsar.Client
}

func (pc *PulsarClient) NewProducer(topic string) (Producer, error) {

	producer, err := pc.client.CreateProducer(pulsar.ProducerOptions{
		Topic:                   topic,
		SendTimeout:             5 * time.Second,
		DisableBlockIfQueueFull: true,
		MaxPendingMessages:      10000,
		CompressionType:         pulsar.LZ4,
		CompressionLevel:        pulsar.Default,
	})
	if err != nil {
		return nil, err
	}

	return &PulsarProducer{
		producer: producer,
	}, err
}

func (pc *PulsarClient) NewConsumer(topic string) (Consumer, error) {

	consumer, err := pc.client.Subscribe(pulsar.ConsumerOptions{
		Topic:                      topic,
		SubscriptionName:           "sub-" + topic,
		Type:                       pulsar.Shared,
		ReplicateSubscriptionState: true,
	})
	if err != nil {
		return nil, err
	}

	return &PulsarConsumer{consumer}, nil
}

func NewPulsarClient(mqurl string) (Client, error) {
	c := &PulsarClient{
		url: mqurl,
	}

	if client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:                     mqurl,
		MaxConnectionsPerBroker: 10,
		Logger:                  log.NewLoggerWithLogrus(defaultLogger),
	}); err != nil {
		return nil, err
	} else {
		c.client = client
		return c, nil
	}
}

func (pc *PulsarClient) Close() {
	pc.client.Close()
}
