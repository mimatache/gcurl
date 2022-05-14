package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client interface {
	WithHeader(key, value string) Client
	WithPayload(v interface{}) Client
	Do(method string) (*Response, error)
}

type Response struct {
	*http.Response
}

func (r *Response) GetBody() string {
	if r.Response.Body == nil {
		return ""
	}
	defer r.Response.Body.Close()
	switch r.Response.Header.Get("Content-Type") {
	case "application/json":
		output, err := jsonDecode(r.Response.Body)
		if err != nil {
			return stringDecode(r.Response.Body)
		}
		return output
	default:
		return stringDecode(r.Response.Body)
	}
}

func jsonDecode(reader io.Reader) (string, error) {
	resp := map[string]interface{}{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&resp)
	if err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}
	output, err := json.MarshalIndent(resp, "", "\t")
	return string(output), err
}

func stringDecode(reader io.Reader) string {
	bodyB, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	bodyStr := string(bytes.ReplaceAll(bodyB, []byte("\r"), []byte("\r\n")))
	return bodyStr
}

type Option = func(c *http.Client) *http.Client

func New(address string, options ...Option) Client {
	clnt := http.DefaultClient
	for _, o := range options {
		clnt = o(clnt)
	}
	return &simpleClient{
		addr:   address,
		client: clnt,
	}
}

type simpleClient struct {
	addr    string
	payload interface{}
	headers map[string]string
	client  *http.Client
}

func (c *simpleClient) WithHeader(key, value string) Client {
	c2 := *c
	if c2.headers == nil {
		c2.headers = map[string]string{}
	}
	c2.headers[key] = value
	return &c2
}

func (c *simpleClient) WithPayload(v interface{}) Client {
	c2 := *c
	c2.payload = v
	return &c2
}

func (c *simpleClient) Do(method string) (*Response, error) {
	payload, err := json.Marshal(c.payload)
	if err != nil {
		return nil, err
	}
	path := c.addr
	req, err := http.NewRequest(method, path, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{
		Response: resp,
	}, nil
}
