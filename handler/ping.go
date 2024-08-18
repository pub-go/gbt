package handler

import (
	"code.gopub.tech/gbt/common/errs"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
	"github.com/youthlin/t/errors"
)

// Ping 接口测试
func Ping(ctx *gin.Context) (any, error) {
	if e := ctx.Query("e"); e != "" {
		return "", errs.ErrBadRequest.WithCause(errors.Errorf("return error: %v", e))
	}
	t := t.WithContext(ctx)
	return t.T("pong"), nil
}
