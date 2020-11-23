package http

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	ConnectTimeout time.Duration
	RequestTimeout time.Duration
	KeepAliveTime  time.Duration
	AcceptContent  string
	AcceptEncoding string
}

type Client struct {
	baseUrl string
	headers map[string]string
	config  *Config
}

func (myself *Client) SetBaseUrl(baseUrl string) *Client {
	myself.baseUrl = baseUrl
	if "" != myself.baseUrl && strings.HasSuffix(myself.baseUrl, "/") {
		myself.baseUrl = myself.baseUrl[0 : len(myself.baseUrl)-1]
	}

	return myself
}

func (myself *Client) SetHeaders(headers map[string]string) *Client {
	if nil == headers {
		return myself
	}

	for key, value := range headers {
		myself.headers[key] = value
	}

	return myself
}

func (myself *Client) SetAcceptContent(acceptContent string) *Client {
	myself.config.AcceptContent = acceptContent
	myself.headers["Accept"] = myself.config.AcceptContent

	return myself
}

func (myself *Client) SetAcceptEncoding(acceptEncoding string) *Client {
	myself.config.AcceptEncoding = acceptEncoding
	myself.headers["Accept-Encoding"] = myself.config.AcceptEncoding

	return myself
}

func (myself *Client) SetTimeout(connectTimeout time.Duration, requestTimeout time.Duration, keepAliveTime time.Duration) *Client {
	myself.config.ConnectTimeout = connectTimeout
	myself.config.RequestTimeout = requestTimeout
	myself.config.KeepAliveTime = keepAliveTime

	return myself
}

func (myself *Client) Get(uri string, headers ...map[string]string) ([]byte, error) {
	return myself.request("GET", uri, nil, headers...)
}

func (myself *Client) Post(uri string, body interface{}, headers ...map[string]string) ([]byte, error) {
	return myself.request("POST", uri, body, headers...)
}

func (myself *Client) Put(uri string, body interface{}, headers ...map[string]string) ([]byte, error) {
	return myself.request("PUT", uri, body, headers...)
}

func (myself *Client) Delete(uri string, headers ...map[string]string) ([]byte, error) {
	return myself.request("DELETE", uri, nil, headers...)
}

func (myself *Client) request(method string, uri string, body interface{}, headers ...map[string]string) ([]byte, error) {
	var err error
	request := &http.Request{
		Method:     method,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	request.URL, err = myself.buildUrl(uri)
	if nil != err {
		return nil, err
	}
	myself.setHeaders(request, headers...)
	_ = myself.setBody(request, body)

	client := myself.buildClient()
	response, err := client.Do(request)
	if nil != err {
		return nil, err
	}
	if nil != response.Body {
		defer response.Body.Close()
	}

	return myself.getBytes(response)
}

func (myself *Client) buildClient() *http.Client {
	transport := &http.Transport{
		IdleConnTimeout:       45 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
		DisableKeepAlives:     false,
		DisableCompression:    false,
		TLSHandshakeTimeout:   10 * time.Second,
		MaxIdleConnsPerHost:   200,
		MaxIdleConns:          2000,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			dialer := &net.Dialer{
				Timeout:   myself.config.ConnectTimeout,
				KeepAlive: myself.config.KeepAliveTime,
				DualStack: true,
			}
			conn, err := dialer.DialContext(ctx, network, addr)
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
	}

	return &http.Client{
		Transport: transport,
		Timeout:   myself.config.RequestTimeout,
	}
}

func (myself *Client) buildUrl(uri string) (*url.URL, error) {
	if "" == myself.baseUrl {
		return url.Parse(uri)
	}

	if "" == uri || strings.HasPrefix(uri, "/") {
		return url.Parse(myself.baseUrl + uri)
	}

	return url.Parse(myself.baseUrl + "/" + uri)
}

func (myself *Client) setBody(request *http.Request, body interface{}) error {
	var err error
	var buffer []byte
	if nil != body {
		buffer, err = json.Marshal(body)
		if nil != err {
			return err
		}
	}
	request.Body = ioutil.NopCloser(bytes.NewReader(buffer))
	request.ContentLength = int64(len(buffer))

	return nil
}

func (myself *Client) setHeaders(request *http.Request, headers ...map[string]string) {
	for key, value := range myself.headers {
		request.Header.Set(key, value)
	}

	if nil == headers || 0 == len(headers) {
		return
	}

	keyValues := headers[0]
	for key, value := range keyValues {
		request.Header.Set(key, value)
	}
}

func (myself *Client) getBytes(response *http.Response) ([]byte, error) {
	if strings.EqualFold("gzip", response.Header.Get("Content-Encoding")) {
		reader, err := gzip.NewReader(response.Body)
		if nil != err {
			return nil, err
		}

		return ioutil.ReadAll(reader)
	}

	return ioutil.ReadAll(response.Body)
}
