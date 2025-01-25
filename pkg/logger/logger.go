/**
 * Logger Implementation
 *
 * @author jarryli@gmail.com
 * @date 2024-12-20
 */

package logger

import (
  "os"

  "github.com/sirupsen/logrus"
  "gopkg.in/natefinch/lumberjack.v2"
)

var log = logrus.New()

// Init 初始化日志配置
func Init() {
  log.SetOutput(os.Stdout)                  // 默认输出到标准输出
  log.SetFormatter(&logrus.JSONFormatter{}) // 使用JSON格式
  log.SetLevel(logrus.InfoLevel)            // 设置默认日志级别为 Info
}

// SetLogLevel 设置日志级别
func SetLogLevel(level logrus.Level) {
  log.SetLevel(level)
}

// Debug 输出调试级别日志
func Debug(msg ...interface{}) {
  log.Debug(msg...)
}

// Info 输出普通信息级别日志
func Info(msg ...interface{}) {
  log.Info(msg...)
}

// Warn 输出警告级别日志
func Warn(msg ...interface{}) {
  log.Warn(msg...)
}

// Error 输出错误级别日志
func Error(msg ...interface{}) {
  log.Error(msg...)
}

// Printf 输出普通日志
func Printf(format string, msg ...interface{}) {
  log.Printf(format, msg...)
}

// Println 输出普通日志
func Println(msg ...interface{}) {
  log.Println(msg...)
}

// Errorf 输出带格式的错误日志
func Errorf(format string, args ...interface{}) {
  log.Errorf(format, args...)
}

// Fatal 输出致命错误日志，并停止程序
func Fatal(msg ...interface{}) {
  log.Fatal(msg...)
}

// Fatalln 输出致命错误日志，并停止程序
func Fatalln(msg ...interface{}) {
  log.Fatalln(msg...)
}

// Fatalf 输出致命错误日志，并停止程序
func Fatalf(format string, msg ...interface{}) {
  log.Fatalf(format, msg...)
}

// Panic 输出严重错误日志，并触发panic
func Panic(msg ...interface{}) {
  log.Panic(msg...)
}

// WithFields 使用字段记录日志
func WithFields(fields logrus.Fields) *logrus.Entry {
  return log.WithFields(fields)
}

// Log 返回 logrus 实例
func Log() *logrus.Logger {
  return log
}

// SetLogToFile 配置将日志写入指定文件，支持日志轮转
func SetLogToFile(filename string, maxSize int, maxBackups int, maxAge int, compress bool) {
  log.SetOutput(&lumberjack.Logger{
    Filename:   filename,   // 日志文件名
    MaxSize:    maxSize,    // 文件大小上限，单位 MB
    MaxBackups: maxBackups, // 保留的最大备份数
    MaxAge:     maxAge,     // 文件保存的最大天数
    Compress:   compress,   // 是否压缩旧日志
  })
}
