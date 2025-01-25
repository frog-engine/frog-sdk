package main

import (
  "fmt"
  "path/filepath"

  frogsdk "github.com/frog-engine/frog-sdk"

  "github.com/frog-engine/frog-sdk/config"
  "github.com/frog-engine/frog-sdk/internal/models"
  "github.com/frog-engine/frog-sdk/internal/utils"
  "github.com/frog-engine/frog-sdk/pkg/logger"
)

func main() {

  // 创建一个新的 SDKApi 实例
  sdkApi := frogsdk.Init(&config.Config{ToolsRoot: "/usr/bin/"}, "production")
  // sdkApi := api.NewImageApi()
  logger.Printf("api %v", sdkApi)

  var image1 = filepath.Join(utils.GetHomeDir(), "Downloads/tmp", "image1.jpg")
  // 获取图片基本信息示例
  resp, err := sdkApi.GetImageInfo(&models.ImageInfoRequest{FilePaths: []string{image1}})
  if err != nil {
    logger.Fatalf("Error getting image info: %v", err)
  }

  // 打印获取的图片信息
  for _, info := range resp {
    fmt.Printf("Image: %s, Format: %s, Size: %dx%d\n", info.FilePath, info.Format, info.Width, info.Height)
  }

  // 示例：转换图片格式
  convertResp, err := sdkApi.ImageConvert(&models.ConvertRequest{
    SourcePath:   image1,
    TargetFormat: "png",
  })

  if err != nil {
    logger.Fatalf("Error converting image: %v", err)
  }

  fmt.Printf("Converted image saved at: %s\n", convertResp.OutputPath)

  // 获取任务信息示例
  taskID := "task123"
  taskResp, err := sdkApi.GetTaskInfo(taskID)
  if err != nil {
    logger.Fatalf("Error getting task info: %v", err)
  }

  if taskResp != nil {
    fmt.Printf("Task ID: %s, Status: %s\n", taskResp.ID, taskResp.Status)
  }
}
