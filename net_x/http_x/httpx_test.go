package httpx

import (
	"fmt"
	"testing"
)

func TestHttpx(t *testing.T) {
	c := NewHttpClient(SetMaxIdleConns(10), SetMaxIdleConnsPerHost(5))
	req, err := NewRequest("GET", "https://baidu.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = c.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

}
