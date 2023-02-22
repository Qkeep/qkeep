package request

import "github.com/go-resty/resty/v2"

type RequestBuilder struct{}

func (rb *RequestBuilder) Build(req *resty.Request, data map[string]interface{}) *resty.Request {
	return req
}
