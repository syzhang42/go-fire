package httpx

import (
	"context"
	"net"
	"net/http"
	"time"
)

type Optionx func(*http.Transport)

/*
	 默认值：DialContext: defaultTransportDialContext(&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}),
*/
func SetDial(timeout int, keepAlive int) Optionx {
	return func(t *http.Transport) {
		t.DialContext = defaultTransportDialContext(&net.Dialer{
			Timeout:   time.Duration(timeout) * time.Second,
			KeepAlive: time.Duration(keepAlive) * time.Second,
		})
	}
}

/*
默认值：true 是否使用http2
*/
func SetForceAttemptHTTP2(open bool) Optionx {
	return func(t *http.Transport) {
		t.ForceAttemptHTTP2 = open
	}
}

/*
默认值：100 客户端可以保持的最大空闲连接数。
*/
func SetMaxIdleConns(num int) Optionx {
	return func(t *http.Transport) {
		t.MaxIdleConns = num
	}
}

/*
默认值：10 每个主机（Host）可以保持的最大空闲连接数
*/
func SetMaxIdleConnsPerHost(num int) Optionx {
	return func(t *http.Transport) {
		t.MaxIdleConnsPerHost = num
	}
}

/*
默认值：90  单位：s 空闲连接的超时时间
*/
func SetIdleConnTimeout(timeout int) Optionx {
	return func(t *http.Transport) {
		t.IdleConnTimeout = time.Duration(timeout) * time.Second
	}
}

/*
默认值：10  单位：s TLS握手的超时时间
*/
func SetTLSHandshakeTimeout(timeout int) Optionx {
	return func(t *http.Transport) {
		t.TLSHandshakeTimeout = time.Duration(timeout) * time.Second
	}
}

/*
默认值：1  单位：s 发送包含 Expect: 100-continue 标头的请求时，等待服务器发送 100 Continue 状态码的超时时间。如果在1秒内没有收到响应，客户端将发送请求的正文。
*/
func SetExpectContinueTimeout(timeout int) Optionx {
	return func(t *http.Transport) {
		t.ExpectContinueTimeout = time.Duration(timeout) * time.Second
	}
}
func defaultTransportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	return dialer.DialContext
}
func NewHttpClient(opts ...Optionx) *http.Client {
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: defaultTransportDialContext(&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}),

		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	for _, opt := range opts {
		opt(tr)
	}
	return &http.Client{
		Transport: tr,
	}
}

func NewHttpClientWithTimeOut(timeout time.Duration, opts ...Optionx) *http.Client {
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: defaultTransportDialContext(&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}),

		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	for _, opt := range opts {
		opt(tr)
	}
	return &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}
}
