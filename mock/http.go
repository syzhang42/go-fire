package mock

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	httpx "github.com/syzhang42/go-fire/net_x/http_x"
)

func RunHttpServer(host string) (stop chan<- struct{}, err error) {
	s := make(chan struct{})

	g := gin.New()
	g.Use()
	g.GET("/", handler())
	srv := &http.Server{
		Addr:    host,
		Handler: g,
	}
	fmt.Printf("server->Server is running at http://%v\n", host)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("server->run http server error:%v\n", err)
		}
	}()

	go func() {
		<-s
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("server->Server Shutdown:", err)
		}
	}()
	time.Sleep(time.Second)
	return s, nil
}

func handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "success", "code": 0, "data": "hello word"})
	}
}

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)

func RunHttpCli(method Method, host string, number uint8) error {
	c := httpx.NewHttpClient(httpx.SetMaxIdleConns(100), httpx.SetMaxIdleConnsPerHost(10), httpx.SetForceAttemptHTTP2(false))
	req, err := httpx.NewRequest(string(method), host, nil)
	if err != nil {
		return err
	}
	for i := 0; i < int(number); i++ {
		_, err = c.Do(req)
		if err != nil {
			return err
		}
		fmt.Printf("client->do %v successful\n", i)
		time.Sleep(time.Second)
	}
	return nil
}
