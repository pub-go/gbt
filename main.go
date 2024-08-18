package main

import (
	"context"
	"embed"
	"flag"
	"path/filepath"

	"code.gopub.tech/gbt/common/conf"
	"code.gopub.tech/gbt/webs"
	"code.gopub.tech/logs"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
)

//go:embed resource/lang
var language embed.FS
var ctx = context.Background()
var dir = flag.String("data", ".", "data dir")

func main() {
	MustInit() // 初始化
	startWeb() // 开启服务
}

func MustInit() {
	flag.Parse()
	dir := *dir

	abs, err := filepath.Abs(dir)
	if err != nil {
		logs.Panic(ctx, "failed to get abs path of %v: %+v", dir, err)
	}

	// logs 日志输出控制台、文件
	logs.SetDefault(logs.NewLogger(logs.CombineHandlers(
		logs.NewHandler(), // console
		logs.NewHandler(logs.WithFile(filepath.Join(abs, "logs", "app.log"))), // log file
	)))
	logs.Info(ctx, "use data dir %q. starting app...", abs)

	if err = conf.ReadConfig(dir); err != nil {
		logs.Panic(ctx, "failed to read/create config file: %+v", err)
	}

	webs.InitI18n(language)
	logs.Info(ctx, t.T("Hello, World"))
	logs.Debug(ctx, t.N("I have One appale.", "I've %v apples.", 2, 2))
}

func startWeb() {
	r := gin.Default()
	register(r)
	addr := conf.AppConf.Addr
	logs.Info(ctx, t.T("app run on %s", addr))
	logs.Info(ctx, "run: %+v", r.Run(addr))
}
