package mock

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func RunTcpServer(host string) (stop chan<- struct{}, err error) {
	s := make(chan struct{})

	listen, err := net.Listen("tcp", host)
	if err != nil {
		return nil, err
	}
	fmt.Printf("server->Server is listening on %v...\n", host)
	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				log.Printf("server->Error accepting connection: %v", err)
				return
			}
			defer conn.Close()
			fmt.Println("server->Client connected:", conn.RemoteAddr())
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				if err.Error() == "EOF" {
					fmt.Println("server->Client disconnected")
				} else {
					fmt.Printf("server->Error reading from client: %v", err)
				}
			}
			fmt.Printf("server->Received from client: %s\n", string(buf[:n]))
		}
	}()

	go func() {
		<-s
		fmt.Println("server->Server is close")
		listen.Close()
	}()

	return s, nil
}

func RunTcpClient(host string, number uint8) error {
	for i := 0; i < int(number); i++ {
		conn, err := net.Dial("tcp", host)
		if err != nil {
			return err
		}
		message := "Hello from client!:" + strconv.Itoa(i)
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("client->Error sending message to server:", err)
			return err
		}
		fmt.Printf("client->successful write:%v\n", i)
		time.Sleep(time.Second)
		conn.Close()
		fmt.Printf("client->client Closing connection %v\n", i)

	}
	return nil
}
