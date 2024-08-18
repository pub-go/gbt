package handler

import (
	"code.gopub.tech/gbt/common/errs"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t/errors"
)

func Ping(ctx *gin.Context) (any, error) {
	if e := ctx.Query("e"); e != "" {
		return "", errs.ErrBadRequest.WithCause(errors.Errorf("return error: %v", e))
	}
	return "pong", nil
}
