package requests

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Method 定义HTTP方法的类型，以提供类型安全的方法名称。
type Method string

// Params 类型定义请求的参数。
type Params map[string]string

// Headers 类型封装http.Header，用于请求的头部信息管理。
type Headers map[string]string

// Form 类型定义表单数据。
type Form map[string]any

// Json 类型用于承载JSON格式的数据。
type Json interface{}

// BaseRequest 类型封装了HTTP请求的基础信息。
type BaseRequest struct {
	Method      Method
	Url         string
	Headers     Headers
	Params      Params
	ContentType string
	Body        io.Reader
	Form        Form
	Json        Json
	formSet     bool
	jsonSet     bool
}

// 定义了常见的HTTP方法常量。
const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
)

// Methods 是一个映射，用于将字符串方法名称映射到Method类型。
var Methods = map[string]Method{
	"GET":     GET,
	"POST":    POST,
	"PUT":     PUT,
	"DELETE":  DELETE,
	"PATCH":   PATCH,
	"HEAD":    HEAD,
	"OPTIONS": OPTIONS,
}

// ParseArgs 解析传入的参数，构建一个BaseRequest对象。
// 它支持多种类型的参数，包括HTTP方法、头部、参数、表单数据和JSON数据。
// 参数:
//
//	args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
//
// 返回值:
//
//	*BaseRequest: 根据传入的参数构建的BaseRequest对象。
//	error: 如果解析失败或参数不合法，则返回错误。
func ParseArgs(args ...any) (*BaseRequest, error) {
	req := &BaseRequest{}
	for _, arg := range args {
		switch d := arg.(type) {
		case Method:
			req.Method = d
		case Headers:
			req.Headers = d
		case Params:
			req.Params = d
		case string:
			if _, err := url.Parse(d); err == nil {
				req.Url = d
			} else {
				if method, ok := Methods[strings.ToUpper(d)]; req.Method == "" && ok {
					req.Method = method
				} else {
					return nil, errors.Wrapf(err, "invalid string %s", d)
				}
			}
		case Form:
			if req.formSet || req.jsonSet {
				return nil, errors.New("form and json data can not be set at the same time")
			}
			req.Form = d
			req.ContentType = "application/x-www-form-urlencoded"
			marshal, err := json.Marshal(d)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to marshal form data %v", d)
			}
			req.Body = bytes.NewReader(marshal)
			req.formSet = true
		case Json:
			if req.formSet || req.jsonSet {
				return nil, errors.New("form and json data can not be set at the same time")
			}
			req.Json = d
			req.ContentType = "application/json"
			marshal, err := json.Marshal(d)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to marshal json data %v", d)
			}
			req.Body = bytes.NewReader(marshal)
			req.jsonSet = true

		}
	}
	if req.Method == "" {
		req.Method = GET
	}
	if req.Url == "" {
		return nil, errors.New("url is required")
	}
	if req.ContentType != "" {
		if req.Headers == nil {
			req.Headers = make(Headers)
		}
		req.Headers["Content-Type"] = req.ContentType
	}
	return req, nil
}

func (r *BaseRequest) BuildRequest() (*http.Request, error) {
	req, err := http.NewRequest(string(r.Method), r.Url, r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create request for %s %s", r.Method, r.Url)
	}
	if r.Headers != nil {
		for k, v := range r.Headers {
			req.Header.Set(k, v)
		}
	}
	if r.Params != nil {
		q := req.URL.Query()
		for k, v := range r.Params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	return req, nil
}
