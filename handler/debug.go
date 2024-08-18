package handler

import (
	"code.gopub.tech/errors"
	"code.gopub.tech/gbt/common/errs"
	"code.gopub.tech/gbt/service"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
)

// Upload 上传一个 torrent 种子文件测试
func Upload(ctx *gin.Context) (any, error) {
	t := t.WithContext(ctx)
	header, err := ctx.FormFile("file")
	if err != nil {
		return nil, errs.ErrBadRequest.WithCause(errors.Wrapf(err, t.T("failed to get upload file")))
	}
	f, err := header.Open()
	if err != nil {
		return nil, errs.ErrBadRequest.WithCause(errors.Wrapf(err, t.T("failed to open upload file: %s", header.Filename)))
	}
	meta, err := service.ReadMeta(f)
	if err != nil {
		return nil, errs.ErrBadRequest.WithCause(errors.Wrapf(err, t.T("failed to read upload file: %s", header.Filename)))
	}
	return meta, nil
}
