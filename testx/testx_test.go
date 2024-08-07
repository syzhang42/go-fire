package testx

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	todo := func() {
		time.Sleep(time.Second)
	}
	res, err := Begin(SetThreadNum(2), SetdurTime(5), SetTodo(todo), SetPrintTime(1))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
