package http

func NewClient(config *Config) *Client {
	client := &Client{
		headers: map[string]string{},
	}
	if nil == config {
		client.config = &Config{
			ConnectTimeout: ConnectTimeout,
			RequestTimeout: RequestTimeout,
			KeepAliveTime:  KeepAliveTime,
			AcceptContent:  AcceptContent,
			AcceptEncoding: AcceptEncoding,
		}
	} else {
		client.config = config
	}

	client.headers["Accept-Encoding"] = client.config.AcceptEncoding
	client.headers["Accept"] = client.config.AcceptContent
	if 0 < client.config.KeepAliveTime {
		client.headers["Connection"] = "Keep-Alive"
	}
	client.headers["Content-Type"] = "application/json;charset=utf-8"

	return client
}
