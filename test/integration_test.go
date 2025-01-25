package test

import (
  "fmt"
  "log"
  "testing"

  "github.com/frog-engine/frog-sdk/internal/api"
  "github.com/frog-engine/frog-sdk/internal/models"
)

func TestIntegration(t *testing.T) {
  // 创建一个新的 SDKApi 实例
  sdkApi := api.NewImageApi()

  // 获取图片基本信息示例
  resp, err := sdkApi.GetImageInfo(&models.ImageInfoRequest{FilePaths: []string{"image1.jpg"}})
  if err != nil {
    log.Fatalf("Error getting image info: %v", err)
  }

  // 打印获取的图片信息
  for _, info := range resp {
    fmt.Printf("Image: %s, Format: %s, Size: %dx%d\n", info.FilePath, info.Format, info.Width, info.Height)
  }

}
