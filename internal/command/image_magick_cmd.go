package cmd

import (
  "fmt"
  "os/exec"

  "github.com/frog-engine/frog-sdk/pkg/logger"
)

// ImageMagickGetImageDimensions 使用 ImageMagick 获取图片的宽度和高度
func GetImageDimensions(filePath string) (int, int, error) {
  // cmd := exec.Command("identify", "-format", "%w %h", filePath)
  cmd := exec.Command("stat", filePath)

  output, err := cmd.Output()
  if err != nil {
    return 0, 0, fmt.Errorf("failed to identify image %s: %w", filePath, err)
  }

  var width, height int
  _, err = fmt.Sscanf(string(output), "%d %d", &width, &height)
  if err != nil {
    return 0, 0, fmt.Errorf("failed to parse dimensions: %w", err)
  }

  return width, height, nil
}

// ImageMagickConvertImageFormat 使用 ImageMagick 转换图片格式
func ConvertImageFormat(sourcePath, outputPath string) error {
  logger.Info("sourcePath", sourcePath, outputPath)
  return nil
  // cmd := exec.Command("convert", sourcePath, outputPath)
  // return cmd.Run()
}

// ImageMagickResizeImage 使用 ImageMagick 调整图片大小
func ResizeImage(sourcePath, outputPath string, width, height int) error {
  cmd := exec.Command("convert", sourcePath, "-resize", fmt.Sprintf("%dx%d", width, height), outputPath)
  return cmd.Run()
}

// ImageMagickRotateImage 使用 ImageMagick 旋转图片
func RotateImage(sourcePath, outputPath string, degrees int) error {
  cmd := exec.Command("convert", sourcePath, "-rotate", fmt.Sprintf("%d", degrees), outputPath)
  return cmd.Run()
}
