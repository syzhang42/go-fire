package testx

import (
	"container/heap"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	heapx "github.com/syzhang42/go-fire/fire/heap"
)

type option func(t *Testx)

// 默认开启
func SetUseP50(open bool) option {
	return func(t *Testx) {
		t.useP50 = open
	}
}

// 默认开启
func SetUseQps(open bool) option {
	return func(t *Testx) {
		t.useQps = open
	}
}

// 默认开启
func SetUseReq(open bool) option {
	return func(t *Testx) {
		t.useReq = open

	}
}

// 默认开启
func SetUseAvg(open bool) option {
	return func(t *Testx) {
		t.useAvg = open
	}
}

// 默认开启
func SetUseP99(open bool) option {
	return func(t *Testx) {
		t.useP99 = open
	}
}

// 默认开启
func SetUseP95(open bool) option {
	return func(t *Testx) {
		t.useP95 = open
	}
}

// 默认开启
func SetUseP90(open bool) option {
	return func(t *Testx) {
		t.useP90 = open
	}
}

// 默认4 线程
func SetThreadNum(i int) option {
	return func(t *Testx) {
		t.threadNum = i
	}
}

// 默认 60 单位s
func SetdurTime(i int) option {
	return func(t *Testx) {
		t.durTime = i
	}
}
func SetPrintTime(i int) option {
	return func(t *Testx) {
		t.printTime = i
	}
}
func SetTodo(f func()) option {
	return func(t *Testx) {
		t.todo = f
	}
}

type Testx struct {
	useP50 bool //延迟50线
	useP90 bool //延迟90线
	useP95 bool //延迟95线
	useP99 bool //延迟99线
	useAvg bool //平均延迟
	useReq bool //总请求数
	useQps bool //qps

	mu         sync.Mutex
	percentile heapx.Heapx[int] //延迟堆
	requests   int              //总请求

	threadNum int //并发数
	durTime   int //持续时间  s
	printTime int

	todo func()

	stop chan struct{}
}

var testx *Testx

func Begin(opts ...option) (res map[string]string, err error) {
	res = make(map[string]string)
	if testx == nil {
		testx = &Testx{
			useP50:     true,
			useP90:     true,
			useP99:     true,
			useP95:     true,
			useAvg:     true,
			useReq:     true,
			useQps:     true,
			percentile: *heapx.NewMinHeap([]int{}),

			threadNum: 4,
			durTime:   60,
			printTime: 60,
			stop:      make(chan struct{}, 1),
		}
		for _, opt := range opts {
			opt(testx)
		}
		if testx.todo == nil {
			return res, errors.New("func is nil")
		}
		if testx.threadNum <= 0 {
			return res, errors.New("threadNum must > 0")
		}
		if testx.durTime <= 0 {
			return res, errors.New("durTime must > 0")
		}
		if !testx.useP50 && !testx.useP90 && !testx.useP95 && !testx.useP99 && !testx.useAvg && !testx.useReq && !testx.useQps {
			return res, errors.New("all indicators are disabled")
		}
	}
	fmt.Println("begin:", "threadNum:", testx.threadNum, "durtime", testx.durTime, "......")

	var wg sync.WaitGroup
	for i := 0; i < testx.threadNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			endTime := time.Now().Add(time.Duration(testx.durTime) * time.Second)
			for time.Now().Before(endTime) {
				todoBase := time.Now()
				testx.todo()
				todoElapsedFlt := time.Since(todoBase).Milliseconds()
				testx.mu.Lock()
				testx.requests = testx.requests + 1
				heap.Push(&testx.percentile, int(todoElapsedFlt))
				testx.mu.Unlock()
			}
		}()
	}
	go func() {
		ticker := time.NewTicker(time.Duration(testx.printTime) * time.Second)
		defer ticker.Stop()
	Loop:
		for {
			select {
			case t := <-ticker.C:
				// 打印当前时间
				fmt.Println("当前时间:", t.Format("2006-01-02 15:04:05"), "......")
				//优雅下线
			case <-testx.stop:
				break Loop
			}
		}
	}()
	wg.Wait()
	testx.stop <- struct{}{}
	var tempPercentile = make([]int, 0)
	for {
		if testx.percentile.Len() == 0 {
			break
		}
		tempPercentile = append(tempPercentile, heap.Pop(&testx.percentile).(int))
	}
	if testx.useP50 {
		// 第50th百分位数位置
		p50_index := int(0.5*float64(testx.requests)) - 1
		if p50_index < 0 {
			p50_index = 0
		}
		p50_response_rate := tempPercentile[p50_index]
		res["50th percentile response rate:"] = strconv.Itoa(int(p50_response_rate))
	}
	if testx.useP90 {
		// 第90th百分位数位置
		p90_index := int(0.90*float64(testx.requests)) - 1
		if p90_index < 0 {
			p90_index = 0
		}
		p90_response_rate := tempPercentile[p90_index]
		res["90th percentile response rate:"] = strconv.Itoa(int(p90_response_rate))
	}

	if testx.useP95 {
		// 第95th百分位数位置
		p95_index := int(0.95*float64(testx.requests)) - 1
		if p95_index < 0 {
			p95_index = 0
		}
		p95_response_rate := tempPercentile[p95_index]
		res["95th percentile response rate:"] = strconv.Itoa(int(p95_response_rate))
	}
	if testx.useP99 {
		// 第99th百分位数位置
		p99_index := int(0.99*float64(testx.requests)) - 1
		if p99_index < 0 {
			p99_index = 0
		}
		p99_response_rate := tempPercentile[p99_index]
		res["99th percentile response rate:"] = strconv.Itoa(int(p99_response_rate))
	}
	if testx.useAvg {
		// 计算平均速率
		sum := 0
		for _, rate := range tempPercentile {
			sum += int(rate)
		}
		average_response_rate := float64(sum) / float64(testx.requests)
		res["avg percentile response rate:"] = strconv.Itoa(int(average_response_rate))
	}
	//fmt.Println(tempPercentile)
	if testx.useReq {
		res["Requests:"] = strconv.Itoa(int(testx.requests))
	}
	if testx.useQps {
		res["QPS:"] = strconv.FormatFloat((float64(testx.requests) / float64(testx.durTime)), 'f', -1, 64)
	}
	fmt.Println("end!")
	return
}
