package main

import (
	"context"
	"embed"

	"code.gopub.tech/logs"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
)

//go:embed language
var language embed.FS
var ctx = context.Background()

func main() {
	logs.Debug(ctx, t.T("Hello, World"))
	i18n()
	startWeb()
}

func i18n() {
	t.LoadFS(language)
	logs.Debug(ctx, t.N("I have One appale.", "I've %v apples.", 2, 2))
}

func startWeb() {
	r := gin.Default()
	register(r)
	r.Run()
}
