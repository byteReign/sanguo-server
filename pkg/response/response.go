package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sanguo-server/pkg/errors"
)

const CodeSuccess = 0

// Body 统一 API 响应结构
type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// CodeDef 业务码定义
type CodeDef = errors.CodeDef

// PageMeta 分页元信息
type PageMeta struct {
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
}

type renderContext struct {
	httpStatus int
}

// Option 响应可选配置，便于后续扩展字段或 HTTP 状态码
type Option func(*Body, *renderContext)

func init() {
	errors.RegisterCode(CodeSuccess, CodeDef{Message: "success", HTTPStatus: http.StatusOK})
}

// Register 注册或覆盖业务码，便于各模块扩展自己的错误码
func Register(code int, def CodeDef) {
	errors.RegisterCode(code, def)
}

func lookupCode(code int) CodeDef {
	if def, ok := errors.LookupCode(code); ok {
		return def
	}
	return CodeDef{Message: "unknown error", HTTPStatus: http.StatusOK}
}

// WithMessage 覆盖默认提示文案
func WithMessage(message string) Option {
	return func(body *Body, _ *renderContext) {
		body.Message = message
	}
}

// WithMeta 附加扩展信息（分页、追踪 ID 等）
func WithMeta(meta interface{}) Option {
	return func(body *Body, _ *renderContext) {
		body.Meta = meta
	}
}

// WithHTTPStatus 覆盖默认 HTTP 状态码
func WithHTTPStatus(status int) Option {
	return func(_ *Body, ctx *renderContext) {
		ctx.httpStatus = status
	}
}

// WithData 覆盖 data 字段
func WithData(data interface{}) Option {
	return func(body *Body, _ *renderContext) {
		body.Data = data
	}
}

// JSON 底层输出方法，完全自定义时使用
func JSON(c *gin.Context, body Body, opts ...Option) {
	write(c, body, opts...)
}

// Success 成功响应
func Success(c *gin.Context, data interface{}, opts ...Option) {
	write(c, Body{
		Code: CodeSuccess,
		Data: data,
	}, opts...)
}

// Fail 失败响应，message 为空时使用已注册的业务码文案
func Fail(c *gin.Context, code int, message string, opts ...Option) {
	def := lookupCode(code)
	body := Body{Code: code}
	if message != "" {
		body.Message = message
	} else {
		body.Message = def.Message
	}

	ctx := &renderContext{httpStatus: def.HTTPStatus}
	for _, opt := range opts {
		opt(&body, ctx)
	}
	c.JSON(ctx.httpStatus, body)
}

// Page 分页成功响应
func Page(c *gin.Context, data interface{}, meta PageMeta, opts ...Option) {
	opts = append([]Option{WithMeta(meta)}, opts...)
	Success(c, data, opts...)
}

func write(c *gin.Context, body Body, opts ...Option) {
	def := lookupCode(body.Code)
	ctx := &renderContext{httpStatus: def.HTTPStatus}

	for _, opt := range opts {
		opt(&body, ctx)
	}
	if body.Message == "" {
		body.Message = def.Message
	}

	c.JSON(ctx.httpStatus, body)
}
