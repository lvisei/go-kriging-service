package config

import (
	"github.com/lvisei/go-kriging-service/pkg/util/json"
	"os"
	"strings"
	"sync"

	"github.com/koding/multiconfig"
)

var (
	// C 全局配置(需要先执行MustLoad，否则拿不到配置)
	C    = new(Config)
	once sync.Once
)

// MustLoad 加载配置
func MustLoad(fpaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, fpath := range fpaths {
			if strings.HasSuffix(fpath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fpath})
			}
		}

		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(C)
	})
}

// PrintWithJSON 基于JSON格式输出配置
func PrintWithJSON() {
	if C.PrintConfig {
		b, err := json.MarshalIndent(C, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}

// Config 配置参数
type Config struct {
	RunMode     string
	Swagger     bool
	PrintConfig bool
	HTTP        HTTP
	Log         Log
	Monitor     Monitor
	RateLimiter RateLimiter
	CORS        CORS
	GZIP        GZIP
}

// IsDebugMode 是否是debug模式
func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

// Log 日志配置参数
type Log struct {
	Level      int
	Format     string
	Output     string
	OutputFile string
}

// HTTP http配置参数
type HTTP struct {
	Host             string
	Port             int
	CertFile         string
	KeyFile          string
	ShutdownTimeout  int
	MaxContentLength int64
	MaxLoggerLength  int `default:"4096"`
}

// Monitor 监控配置参数
type Monitor struct {
	Enable    bool
	Addr      string
	ConfigDir string
}

// RateLimiter 请求频率限制配置参数
type RateLimiter struct {
	Enable bool
	Count  int64
}

// CORS 跨域请求配置参数
type CORS struct {
	Enable           bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           int
}

// GZIP gzip压缩
type GZIP struct {
	Enable             bool
	ExcludedExtentions []string
	ExcludedPaths      []string
}
