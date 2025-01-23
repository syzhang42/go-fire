package mock

import (
	"testing"

	"github.com/syzhang42/go-fire/auth"
)

func TestHttp(t *testing.T) {
	stop, err := RunHttpServer(":9091")
	auth.Must(err)

	err = RunHttpCli(GET, "http://localhost:9091", 3)
	auth.Must(err)

	stop <- struct{}{}
}

func TestTcp(t *testing.T) {
	stop, err := RunTcpServer(":9091")
	auth.Must(err)

	err = RunTcpClient("localhost:9091", 3)
	auth.Must(err)

	stop <- struct{}{}
}

func TestGrpc(t *testing.T) {
	stop, err := RunGrpcServer(":9091")
	auth.Must(err)

	err = RunGrpcClient("localhost:9091", 3)
	auth.Must(err)

	stop <- struct{}{}
}
