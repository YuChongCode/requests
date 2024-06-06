package requests

import "net/http"

// Session 继承自http.Client，用于管理HTTP请求的会话。
type Session struct {
	http.Client
}

// NewSession 创建并返回一个新的Session实例。
func NewSession() *Session {
	return &Session{
		Client: http.Client{},
	}
}

// Request 发起一个HTTP请求并返回响应。
// args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
// 返回HTTP响应和可能的错误。
func (s *Session) Request(args ...interface{}) (*Response, error) {
	parseReq, err := ParseArgs(args...)
	if err != nil {
		return nil, err
	}
	req, err := parseReq.BuildRequest()
	if err != nil {
		return nil, err
	}
	resp, err := s.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{resp}, nil
}

// Get 发起一个GET请求并返回响应。
// args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
// 返回HTTP响应和可能的错误。
func (s *Session) Get(args ...interface{}) (*Response, error) {
	return s.Request(append(args, GET)...)
}

// Head 发起一个HEAD请求并返回响应。
// args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
// 返回HTTP响应和可能的错误。
func (s *Session) Head(args ...interface{}) (*Response, error) {
	return s.Request(append(args, HEAD)...)
}

// Post 发起一个POST请求并返回响应。
// args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
// 返回HTTP响应和可能的错误。
func (s *Session) Post(args ...interface{}) (*Response, error) {
	return s.Request(append(args, POST)...)
}

// Put 发起一个PUT请求并返回响应。
// args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
// 返回HTTP响应和可能的错误。
func (s *Session) Put(args ...interface{}) (*Response, error) {
	return s.Request(append(args, PUT)...)
}

// Patch 发起一个PATCH请求并返回响应。
// args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
// 返回HTTP响应和可能的错误。
func (s *Session) Patch(args ...interface{}) (*Response, error) {
	return s.Request(append(args, PATCH)...)
}

// Delete 发起一个DELETE请求并返回响应。
// args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
// 返回HTTP响应和可能的错误。
func (s *Session) Delete(args ...interface{}) (*Response, error) {
	return s.Request(append(args, DELETE)...)
}

// Options 发起一个OPTIONS请求并返回响应。
// args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
// 返回HTTP响应和可能的错误。
func (s *Session) Options(args ...interface{}) (*Response, error) {
	return s.Request(append(args, OPTIONS)...)
}
