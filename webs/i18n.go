package webs

import (
	"embed"

	"code.gopub.tech/gbt/common/conf"
	"code.gopub.tech/logs"
	"code.gopub.tech/logs/pkg/kv"

	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
)

// InitI18n 初始化翻译
func InitI18n(defaultLangFs embed.FS) {
	if path := conf.AppConf.LangPath(); path != "" {
		t.Load(path)
		logs.Info(ctx, "set languange path: %v", path)
	} else {
		t.LoadFS(defaultLangFs)
		logs.Info(ctx, "use internal language translations")
	}
	t.SetLocale(conf.AppConf.Lang)
	logs.Info(ctx, "used locale: %s", t.UsedLocale())
}

// I18n 为每个请求决定使用哪种语言
func I18n(c *gin.Context) {
	ctx := GetContext(c)
	lang := t.GetUserLang(c.Request) // 获取浏览器偏好语言
	ctx = kv.Add(ctx, KeyLang, lang) // 在日志中打印
	ctx = t.SetCtxLocale(ctx, lang)  // 存在 ctx 里
	SetContext(c, ctx)               // 设置 ctx
	c.Next()
}
