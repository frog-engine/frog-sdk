package config

import (
	"log"
	"os"
)

// Config 结构体定义 SDK 的配置选项
type Config struct {
	ToolsRoot       string
	FFMpegPath      string
	ImageMagickPath string
	DefaultQuality  int
	TempFilePath    string
	Environment     string
}

// GlobalConfig 全局配置实例
var GlobalConfig *Config

// InitConfig 初始化配置，根据环境变量设置不同的配置
func InitConfig(env string) (conf *Config) {
	// 如果空值则从环境变量中获取
	if env == "" {
		env = os.Getenv("FROG_SDK_ENV")
	}
	// 如果都为空则默认为开发环境
	if env == "" {
		env = "development"
	}

	switch env {
	case "production":
		GlobalConfig = &Config{
			ToolsRoot:       "/usr/local/bin",
			FFMpegPath:      "/usr/local/bin/ffmpeg",
			ImageMagickPath: "/usr/local/bin/convert",
			DefaultQuality:  85,
			TempFilePath:    "./tmp",
			Environment:     "production",
		}
	case "development":
		GlobalConfig = &Config{
			ToolsRoot:       "/usr/local/bin",
			FFMpegPath:      "/usr/bin/ffmpeg",
			ImageMagickPath: "/usr/bin/convert",
			DefaultQuality:  75,
			TempFilePath:    "./tmp",
			Environment:     "development",
		}
	default:
		log.Fatalf("Unknown environment: %s", env)
	}

	log.Printf("Running frogsdk in %s environment", GlobalConfig.Environment)
	return GlobalConfig
}
