package conf

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"code.gopub.tech/errors"
	"code.gopub.tech/logs"
	"code.gopub.tech/logs/pkg/arg"
	"github.com/gin-gonic/gin"
)

var AppConf *Config
var ctx = context.Background()

const (
	ConfigFileName = "conf/app.json"
)

type Config struct {
	dir      string // 配置文件所在的文件夹
	path     string // 配置文件全路径
	Addr     string // 监听地址
	Debug    bool   // 开启调试
	Lang     string // app 使用的默认语言
	Resource string // 资源文件夹(相对路径) 为空表示使用内置资源 (可包含 lang static, views 文件夹)
}

func newDefault() *Config {
	return &Config{
		Addr: ":8182",
	}
}

func (c *Config) AbsDir() string    { return c.dir }
func (c *Config) DebugPath() string { return c.resource("debug") }
func (c *Config) LangPath() string  { return c.resource("lang") }
func (c *Config) resource(sub string) string {
	resource := c.Resource
	if resource != "" { // 指定了资源文件夹
		path := filepath.Join(c.dir, resource, sub)
		if _, err := os.Stat(path); err == nil {
			return path // 资源文件夹中存在指定的子文件夹
		} else {
			logs.Warn(ctx, "conf path=%v err=%+v", path, err)
		}
	}
	return ""
}

// ReadConfig 读取配置文件
func ReadConfig(dir string) error {
	// 如果目录下没有配置文件，会创建一个默认的
	f, err := readOrCreateFile(dir)
	if err != nil {
		return errors.Wrapf(err, "readOrCreateFile")
	}
	// 配置文件的全路径
	path, err := filepath.Abs(f.Name())
	if err != nil {
		return errors.Wrapf(err, "getConfigFileAbsPath")
	}
	logs.Info(ctx, "config file abs path: %s", path)
	// 读取配置文件
	b, err := io.ReadAll(f)
	if err != nil {
		return errors.Wrapf(err, "failed to read config file: %s", ConfigFileName)
	}
	// 反序列出来
	var conf Config
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal config: %s", b)
	}
	conf.dir = filepath.Dir(path)
	conf.path = path
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	AppConf = &conf
	logs.Info(ctx, "read config file ok: %s", arg.JSON(AppConf))
	return nil
}

func readOrCreateFile(dir string) (*os.File, error) {
	fileName := filepath.Join(dir, ConfigFileName)
	logs.Info(ctx, "read config file: %s", fileName)
	f, err := os.Open(fileName)
	if err != nil {
		logs.Notice(ctx, "config file not exist, create: %s", fileName)
		configDir := filepath.Dir(fileName)
		if err := os.MkdirAll(configDir, 0755); err != nil { // drwxr-xr-x
			return nil, errors.Wrapf(err, "can not create dir: %s", configDir)
		}
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644) // -rw-r--r--
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create config file: %s", ConfigFileName)
		}
		b, _ := json.MarshalIndent(newDefault(), "", "\t")
		logs.Info(ctx, "write defaut config: %s", b)
		_, err = f.Write(b)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to write config file: %s", ConfigFileName)
		}
		f, err = os.Open(fileName)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read config file: %s", ConfigFileName)
	}
	return f, nil
}
