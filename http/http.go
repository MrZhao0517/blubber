package http

import (
	"bytes"
	"encoding/json"
	"github.com/gojek/heimdall/v7/httpclient"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type Client struct {
	ctx *httpclient.Client
}

const (
	X_WWW_FORM_URLENCODED = "application/x-www-form-urlencoded"
	JSON                  = "application/json"
	FORM_DATA             = "multipart/form-data"
)

func NewClient(opts ...httpclient.Option) *Client {
	//10分钟超时
	timeout := 10 * time.Minute
	opts = append(opts, httpclient.WithHTTPTimeout(timeout))
	ctx := httpclient.NewClient(opts...)
	return &Client{ctx: ctx}
}

func processingBody(body interface{}, contentType string) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	switch contentType {
	case JSON:
		if t := reflect.TypeOf(body); t.Kind() == reflect.String {
			return strings.NewReader(body.(string)), nil
		}
		if b, err := json.Marshal(body); err != nil {
			return nil, err
		} else {
			return bytes.NewReader(b), nil
		}
	case X_WWW_FORM_URLENCODED:
		return strings.NewReader(body.(string)), nil
	default:
		return body.(io.Reader), nil
	}
}

func (c Client) Post(url string, body interface{}, header http.Header, contentType string) (*http.Response, error) {
	reader, err := processingBody(body, contentType)
	if err != nil {
		return nil, err
	}
	return c.ctx.Post(url, reader, header)
}
