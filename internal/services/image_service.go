package services

import (
  "bytes"
  "fmt"
  "image"
  "os"
  "path/filepath"

  cmd "github.com/frog-engine/frog-sdk/internal/command"
  "github.com/frog-engine/frog-sdk/internal/models"
)

// ImageService 接口定义，处理图片相关的业务逻辑
type ImageService struct {
}

// NewImageService 创建新的图片服务实例
func NewImageService() *ImageService {
  return &ImageService{}
}

// GetImageInfo 获取图片基本信息
func (s *ImageService) GetImageInfo(req *models.ImageInfoRequest) ([]*models.ImageInfoResponse, error) {
  var responses []*models.ImageInfoResponse

  for _, filePath := range req.FilePaths {
    if _, err := os.Stat(filePath); err != nil {
      return nil, fmt.Errorf("failed to get file info for %s: %w", filePath, err)
    }

    // 使用提取的命令获取图片格式和尺寸
    width, height, err := cmd.GetImageDimensions(filePath) // 正确调用新函数
    if err != nil {
      return nil, err
    }

    responses = append(responses, &models.ImageInfoResponse{
      FilePath: filePath,
      Format:   filepath.Ext(filePath)[1:], // 获取扩展名作为格式
      Width:    width,
      Height:   height,
    })
  }

  return responses, nil
}

// Convert 转换图片格式
func (s *ImageService) ImageTranscode(req *models.ConvertRequest) (*models.ConvertResponse, error) {
  if _, err := os.Stat(req.SourcePath); err != nil {
    return nil, fmt.Errorf("source file not found: %w", err)
  }

  outputPath := fmt.Sprintf("%s.%s", req.SourcePath[:len(req.SourcePath)-len(filepath.Ext(req.SourcePath))], req.TargetFormat)

  // 使用提取的命令进行格式转换
  if err := cmd.ConvertImageFormat(req.SourcePath, outputPath); err != nil { // 调用新函数
    return &models.ConvertResponse{}, fmt.Errorf("failed to convert image: %w", err)
  }

  return &models.ConvertResponse{OutputPath: outputPath}, nil
}

// ReadImageBlob 用于读取图像数据并进行解码处理
func (s *ImageService) ReadImageBlob(imageData []byte) (image.Image, error) {
  // 使用 image.Decode 函数来解码图像数据
  img, _, err := image.Decode(bytes.NewReader(imageData))
  if err != nil {
    return nil, fmt.Errorf("failed to decode image data: %w", err)
  }

  // 返回解码后的图像对象
  return img, nil
}
