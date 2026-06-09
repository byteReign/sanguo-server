package response

import (
	"github.com/gin-gonic/gin"

	"sanguo-server/pkg/errors"
	"sanguo-server/pkg/logger"
)

// HandleError 将 error 转为统一 API 响应
// - *errors.BizError → 按业务码返回
// - 其他 error → 记录日志并返回 500
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	if biz, ok := errors.AsBizError(err); ok {
		if biz.Unwrap() != nil {
			logger.WithError(biz.Unwrap()).Warn(biz.Message())
		}
		Fail(c, biz.Code(), biz.Message())
		return
	}

	logger.WithError(err).Error("unhandled error")
	Fail(c, errors.CodeServer, "服务器开小差了")
}
