package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"code.gopub.tech/gbt/common/conf"
	"code.gopub.tech/gbt/handler"
	"code.gopub.tech/gbt/webs"
	"github.com/gin-gonic/gin"
)

//go:embed resource/debug
var debugResource embed.FS

func register(r *gin.Engine) {
	r.ContextWithFallback = true
	r.Use(webs.Trace, webs.I18n)

	api := r.Group("/api")
	api.GET("/ping", Wrap(handler.Ping))

	debug := r.Group("/debug")
	debug.StaticFS("/page/", getDebugPage())
	debugApi := debug.Group("/api")
	debugApi.GET("/ping", Wrap(handler.Ping))
	debugApi.POST("/upload", Wrap(handler.Upload))
}

func getDebugPage() http.FileSystem {
	var fsys fs.FS = os.DirFS("resource/debug") // 相对于执行时的工作目录
	if d := conf.AppConf.DebugPath(); d != "" {
		fsys = os.DirFS(d) // 配置文件指定的
	}
	if _, err := fsys.Open("index.html"); err != nil {
		// 如果不失败 就使用; 如果失败就用内置资源
		fsys, _ = fs.Sub(debugResource, "resource/debug")
	}
	return http.FS(fsys)
}
