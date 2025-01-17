package test

import (
  "fmt"
  "log"
  "testing"

  "frog-sdk/internal/api"
  "frog-sdk/internal/models"
)

func TestGetImageInfo(t *testing.T) {

  // 创建一个新的 SDKApi 实例
  sdkApi := api.NewSDKApi()

  // 获取图片基本信息示例
  resp, err := sdkApi.GetImageInfo(models.ImageInfoRequest{FilePaths: []string{"image1.jpg"}})
  if err != nil {
    log.Fatalf("Error getting image info: %v", err)
  }

  // 打印获取的图片信息
  for _, info := range resp {
    fmt.Printf("Image: %s, Format: %s, Size: %dx%d\n", info.FilePath, info.Format, info.Width, info.Height)
  }

  if err != nil {
    t.Fatalf("Expected no error, got %v", err)
  }

  if len(resp) != 1 {
    t.Fatalf("Expected 1 response, got %d", len(resp))
  }
}
