package webs

import (
	"fmt"
	"time"

	"code.gopub.tech/gbt/util"
	"code.gopub.tech/logs/pkg/kv"
	"github.com/gin-gonic/gin"
)

const (
	KeyReqStart = "reqStart"
	KeyTrace    = "trace"
	KeyLang     = "lang"
	KeyCtx      = "ctx"
	KeyErr      = "err"
	HeaderTrace = "X-Trace-ID"
)

// Trace 为每个请求设置一个唯一标记
func Trace(c *gin.Context) {
	// 请求开始时间
	now := time.Now()
	c.Set(KeyReqStart, now)

	// 生成一个唯一标记
	trace := GenTraceID()
	c.Set(KeyTrace, trace)
	c.Header(HeaderTrace, trace)

	ctx := GetContext(c)
	ctx = kv.Add(ctx,
		"method", c.Request.Method,
		"path", c.Request.URL.Path,
		KeyTrace, trace,
	)

	SetContext(c, ctx)
	c.Next()
}

// GenTraceID 生成 traceID
func GenTraceID() string {
	now := time.Now()
	return fmt.Sprintf("%14s%09d%07s",
		now.Format("20060102150405"), now.Nanosecond(), util.RandStr(7))
}
