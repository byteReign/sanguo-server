package errors

import (
	stderrors "errors"
	"fmt"
)

// BizError 业务错误，Service 层返回，Handler 层统一转响应
type BizError struct {
	code    int
	message string
	err     error
}

// New 创建业务错误
func New(code int, message string) *BizError {
	return &BizError{code: code, message: message}
}

// Wrap 包装底层错误，便于日志追踪
func Wrap(code int, message string, err error) *BizError {
	return &BizError{code: code, message: message, err: err}
}

func (e *BizError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("code=%d message=%s: %v", e.code, e.message, e.err)
	}
	return fmt.Sprintf("code=%d message=%s", e.code, e.message)
}

func (e *BizError) Unwrap() error {
	return e.err
}

func (e *BizError) Code() int {
	return e.code
}

func (e *BizError) Message() string {
	return e.message
}

// WithMessage 基于已有错误复制并覆盖提示文案
func (e *BizError) WithMessage(message string) *BizError {
	return &BizError{code: e.code, message: message, err: e.err}
}

// AsBizError 从 error 链中提取业务错误
func AsBizError(err error) (*BizError, bool) {
	var biz *BizError
	if stderrors.As(err, &biz) {
		return biz, true
	}
	return nil, false
}

// IsCode 判断是否为指定业务码
func IsCode(err error, code int) bool {
	biz, ok := AsBizError(err)
	return ok && biz.code == code
}
