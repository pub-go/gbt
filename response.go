package main

import (
	"fmt"
	"net/http"

	"code.gopub.tech/gbt/common/errs"
	"github.com/gin-gonic/gin"
)

func Wrap(handler func(ctx *gin.Context) (any, error)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		result, err := handler(ctx)
		if err != nil {
			errcode := errs.Of(err)
			code := errcode.Code()
			var data string
			if gin.Mode() != gin.ReleaseMode {
				data = fmt.Sprintf("%+v", err)
			}
			ctx.JSON(code/100_000, Response{
				Code:    code,
				Message: err.Error(),
				Data:    data,
			})
			return
		}
		ctx.JSON(http.StatusOK, Response{Data: result})
	}
}

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
