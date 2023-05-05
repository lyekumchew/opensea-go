package opensea

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/pinealctx/restgo"
)

const (
	APIKey = "X-API-KEY"
)

type Client struct {
	*restgo.Client
	*option
}

func New(fnList ...OptionFn) *Client {
	var o = defaultOption
	for _, fn := range fnList {
		fn(o)
	}
	var header = make(http.Header)
	if o.apiKey != "" {
		header.Set(APIKey, o.apiKey)
	}
	header.Set("Accept", "application/json")

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.DialContext(ctx, network, addr)
		},
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	if o.host != "" {
		u, err := url.Parse(o.baseURL)
		if err != nil {
			panic(err)
		}
		hostname := u.Hostname() + ":443"
		host := o.host + ":443"
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			if addr == hostname {
				addr = host
			}
			return dialer.DialContext(ctx, network, addr)
		}
		transport.DisableKeepAlives = true
		transport.MaxIdleConnsPerHost = -1
	}

	return &Client{
		Client: restgo.New(restgo.WithBaseURL(o.baseURL), restgo.WithGlobalHeader(header), restgo.WithTransport(transport)),
		option: o,
	}
}

func (c *Client) GetApiKey() string {
	return c.apiKey
}

func (c *Client) get(ctx context.Context, resource string, params ...restgo.IParam) (restgo.IResponse, error) {
	var rsp, err = c.Execute(ctx, "GET", resource, params...)
	if err != nil {
		return nil, err
	}
	if c.retryWhenFreqLimit && rsp.StatusCode() == http.StatusTooManyRequests {
		var countValue = ctx.Value(countCtxKey)
		var count = c.retryCount
		if countValue != nil {
			count = countValue.(int)
		}
		if count != 0 {
			time.Sleep(c.retryInterval)
			return c.get(WrapperCountContext(ctx, count-1), resource, params...)
		}
	}
	return rsp, nil
}
