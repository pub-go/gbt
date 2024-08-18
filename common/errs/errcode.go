// errs packege 提供形如 xxx_yy_zzz 的统一错误码
// 其中,
// xxx 采用 HTTP 状态码表示
// yy  表示错误类型 00=普通错误 01=数据库错误 02=缓存错误 03=第三方错误
// zzz 表示细分编码
package errs

import (
	"fmt"
	"strings"

	"code.gopub.tech/errors"
)

var (
	ErrUnknown    = New(500_00_000, "unknown error")
	ErrBadRequest = New(400_00_000, "bad request")
	ErrNotFound   = New(404_00_000, "not found")
)

// ErrCode 错误码接口
type ErrCode interface {
	error
	Code() int
	Message() string
}

// Of 将 err 包装为错误码接口
// 如果 err 为 nil 直接返回 nil
// 否则将其转为 ErrCode 接口, 如果不能转为 ErrCode 则认为是 ErrUnknown
func Of(err error) ErrCode {
	if err == nil {
		return nil
	}
	if c, ok := err.(ErrCode); ok {
		return c
	}
	return ErrUnknown.WithCause(err)
}

// Or 如果 err 上没有带错误码, 则使用给定的错误码包装它
// 如果 err 为 nil 直接返回 nil
func Or(err error, errcode ErrCode) ErrCode {
	if err == nil {
		return nil
	}
	c := Of(err)
	if c.Code() == ErrUnknown.Code() {
		return &errCode{
			code:    errcode.Code(),
			message: errcode.Message(),
			cause:   err,
		}
	}
	return c
}

var (
	_ ErrCode       = (*errCode)(nil)
	_ fmt.Formatter = (*errCode)(nil)
)

// errCode 实现 ErrCode 接口的错误码结构体
type errCode struct {
	code    int
	message string
	cause   error
}

// New 新建一个错误码实例
func New(code int, message string) *errCode {
	return &errCode{code: code, message: message}
}

// WithCause 返回一个新的错误码实例, 带有详情错误原因
// 如果 err 为 nil 直接返回 nil
func (e *errCode) WithCause(err error) ErrCode {
	if err == nil {
		return nil
	}
	return &errCode{
		code:    e.code,
		message: e.message,
		cause:   err,
	}
}

// Code implements ErrCode
// 返回错误码
func (e *errCode) Code() int { return e.code }

// Message implements ErrCode
// 返回错误码描述
func (e *errCode) Message() string { return e.message }

// Error implements error
// 返回错误详情
func (e *errCode) Error() string {
	if e.cause != nil {
		cause := e.cause.Error()
		cause = strings.ReplaceAll(cause, fmt.Sprintf(": [%d] %s", e.code, e.message), "")
		return fmt.Sprintf("[%d] %s: %v", e.code, e.message, cause)
	}
	return fmt.Sprintf("[%d] %s", e.code, e.message)
}

// Format implements fmt.Formatter
// 额外支持 %+v, %q 等格式化动词
func (e *errCode) Format(s fmt.State, verb rune) {
	errors.FormatError(e, s, verb)
}

// Unwrap 返回底层错误原因. Format 格式化时会自动通过 Unwrap 获取错误链
func (e *errCode) Unwrap() error { return e.cause }

// Is 判断该实例是否能当作 target 错误.
// errors.Is 会用到本方法.
// 如果 target 是 ErrCode 类型, 那么比较错误码是否相等;
// 否则看底层错误是否是 target 错误.
func (e *errCode) Is(target error) bool {
	if c, ok := target.(ErrCode); ok {
		return c.Code() == e.code
	}
	return errors.Is(e.cause, target)
}
