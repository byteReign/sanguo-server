package errors

import "net/http"

// 通用错误码（与 HTTP 语义对齐）
const (
	CodeBadRequest   = 400
	CodeUnauthorized = 401
	CodeForbidden    = 403
	CodeNotFound     = 404
	CodeConflict     = 409
	CodeServer       = 500
)

// 业务错误码按模块分段，便于扩展：
// 10000-19999 user
// 20000-29999 hero
// 30000-39999 battle/stage
const (
	CodeUserNotFound = 10001
)

// CodeDef 错误码定义
type CodeDef struct {
	Message    string
	HTTPStatus int
}

var codeRegistry = map[int]CodeDef{
	CodeBadRequest:   {Message: "bad request", HTTPStatus: http.StatusBadRequest},
	CodeUnauthorized: {Message: "unauthorized", HTTPStatus: http.StatusUnauthorized},
	CodeForbidden:    {Message: "forbidden", HTTPStatus: http.StatusForbidden},
	CodeNotFound:     {Message: "not found", HTTPStatus: http.StatusNotFound},
	CodeConflict:     {Message: "conflict", HTTPStatus: http.StatusConflict},
	CodeServer:       {Message: "internal server error", HTTPStatus: http.StatusInternalServerError},

	CodeUserNotFound: {Message: "用户不存在", HTTPStatus: http.StatusOK},
}

// RegisterCode 注册或覆盖错误码，各模块在 init 中扩展
func RegisterCode(code int, def CodeDef) {
	codeRegistry[code] = def
}

// LookupCode 查询错误码定义
func LookupCode(code int) (CodeDef, bool) {
	def, ok := codeRegistry[code]
	return def, ok
}

// 便捷构造方法，Service 层直接使用

func BadRequest(message string) *BizError {
	return New(CodeBadRequest, message)
}

func Unauthorized(message string) *BizError {
	return New(CodeUnauthorized, message)
}

func Forbidden(message string) *BizError {
	return New(CodeForbidden, message)
}

func NotFound(message string) *BizError {
	return New(CodeNotFound, message)
}

func Conflict(message string) *BizError {
	return New(CodeConflict, message)
}

func Internal(message string, err error) *BizError {
	return Wrap(CodeServer, message, err)
}
