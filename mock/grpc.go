package mock

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/syzhang42/go-fire/mock/grpc_helper"
	"google.golang.org/grpc"
)

// 实现 Greeter 服务
type server struct {
	grpc_helper.UnimplementedGreeterServer
}

func RunGrpcServer(host string) (stop chan<- struct{}, err error) {
	_stop := make(chan struct{})
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()

	grpc_helper.RegisterGreeterServer(s, &server{})
	fmt.Printf("server->Server is running at %v\n", host)

	go func() {
		<-_stop
		fmt.Printf("server->Server is stop\n")
		s.Stop()
	}()
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("server->failed to serve: %v", err)
		}
	}()
	time.Sleep(time.Second)
	return _stop, err
}

func (s *server) SayHello(ctx context.Context, req *grpc_helper.HelloRequest) (*grpc_helper.HelloReply, error) {
	// 创建并返回响应
	return &grpc_helper.HelloReply{
		Message: "Hello, " + req.GetName(),
	}, nil
}

func RunGrpcClient(host string, number uint8) error {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("client->error to conn:%v\n", err)
		return err
	}
	defer conn.Close()

	c := grpc_helper.NewGreeterClient(conn)

	// 调用 SayHello 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for i := 0; i < int(number); i++ {
		resp, err := c.SayHello(ctx, &grpc_helper.HelloRequest{Name: "World"})
		if err != nil {
			fmt.Printf("client->could not greet: %v\n", err)
		}
		// 打印服务器响应
		fmt.Println("client->Greeting:", resp.GetMessage())
	}
	return nil
}
