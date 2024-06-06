package requests

// defaultSession 是一个全局变量，用于存储默认的会话对象。
// 它被初始化为一个新会话，用于简化HTTP请求的操作。
var defaultSession *Session

// Requests 是一个通用的HTTP请求方法，可以通过传入不同的参数来发起各种类型的HTTP请求。
// 它封装了defaultSession的Request方法，提供了更灵活的调用方式。
// 参数:
//
//	args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
//
// 返回值:
//
//	*http.Response: HTTP响应对象。
//	error: 错误对象，如果请求过程中发生错误，则非nil。
func Requests(args ...any) (*Response, error) {
	return defaultSession.Request(args...)
}

// Get 是一个简化GET请求的函数。
// 它封装了defaultSession的Get方法，用于发起GET请求。
// 参数:
//
//	args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
//
// 返回值:
//
//	*http.Response: HTTP响应对象。
//	error: 错误对象，如果请求过程中发生错误，则非nil。
func Get(args ...any) (*Response, error) {
	return defaultSession.Get(args...)
}

// Post 是一个简化POST请求的函数。
// 它封装了defaultSession的Post方法，用于发起POST请求。
// 参数:
//
//	args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
//
// 返回值:
//
//	*http.Response: HTTP响应对象。
//	error: 错误对象，如果请求过程中发生错误，则非nil。
func Post(args ...any) (*Response, error) {
	return defaultSession.Post(args...)
}

// Put 是一个简化PUT请求的函数。
// 它封装了defaultSession的Put方法，用于发起PUT请求。
// 参数:
//
//	args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
//
// 返回值:
//
//	*http.Response: HTTP响应对象。
//	error: 错误对象，如果请求过程中发生错误，则非nil。
func Put(args ...any) (*Response, error) {
	return defaultSession.Put(args...)
}

// Delete 是一个简化DELETE请求的函数。
// 它封装了defaultSession的Delete方法，用于发起DELETE请求。
// 参数:
//
//	args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
//
// 返回值:
//
//	*http.Response: HTTP响应对象。
//	error: 错误对象，如果请求过程中发生错误，则非nil。
func Delete(args ...any) (*Response, error) {
	return defaultSession.Delete(args...)
}

// Head 是一个简化HEAD请求的函数。
// 它封装了defaultSession的Head方法，用于发起HEAD请求。
// 参数:
//
//	args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
//
// 返回值:
//
//	*http.Response: HTTP响应对象。
//	error: 错误对象，如果请求过程中发生错误，则非nil。
func Head(args ...any) (*Response, error) {
	return defaultSession.Head(args...)
}

// Options 是一个简化OPTIONS请求的函数。
// 它封装了defaultSession的Options方法，用于发起OPTIONS请求。
// 参数:
//
//	args ...any: 可以是Method、Headers、Params、Form、Json、URL字符串或它们的组合。
//
// 返回值:
//
//	*http.Response: HTTP响应对象。
//	error: 错误对象，如果请求过程中发生错误，则非nil。
func Options(args ...any) (*Response, error) {
	return defaultSession.Options(args...)
}

// init 函数在包首次加载时执行，用于初始化defaultSession变量。
// 它调用NewSession函数来创建一个新的会话对象，并将其赋值给defaultSession变量。
func init() {
	defaultSession = NewSession()
}
