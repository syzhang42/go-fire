package mqx

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

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
		mm.Recv(ctxexit, utils.PulsarTopics[0], func(b Message) {
			if b != nil && b.Payload() != nil {
				fmt.Println("data:", string(b.Payload()), "idc:", b.Key())

			}
		})
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		go mm.Recv(ctxexit, utils.PulsarTopics[1], func(b Message) {
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
