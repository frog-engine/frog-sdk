package utils

import (
  "log"
  "os"
)

// Logger 日志工具
var Logger *log.Logger

// InitLogger 初始化日志
func InitLogger() {
  Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
