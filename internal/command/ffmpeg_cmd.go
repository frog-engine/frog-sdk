package cmd

import (
  "os/exec"
)

// FFmpegConvertImageFormat 转换图片格式
func FFmpegConvertImageFormat(sourcePath, outputPath string) error {
  cmd := exec.Command("ffmpeg", "-i", sourcePath, outputPath)
  return cmd.Run()
}
