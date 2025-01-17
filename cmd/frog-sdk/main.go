package main

import (
  "fmt"
  "log"

  "frog-sdk/internal/api"
  "frog-sdk/internal/models"
)

func main() {
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

  // 示例：转换图片格式
  convertResp, err := sdkApi.ImageConvert(models.ConvertRequest{
    SourcePath:   "image1.jpg",
    TargetFormat: "png",
  })

  if err != nil {
    log.Fatalf("Error converting image: %v", err)
  }

  fmt.Printf("Converted image saved at: %s\n", convertResp.OutputPath)

  // 获取任务信息示例
  taskID := "task123"
  taskResp, err := sdkApi.GetTaskInfo(taskID)
  if err != nil {
    log.Fatalf("Error getting task info: %v", err)
  }

  fmt.Printf("Task ID: %s, Status: %s\n", taskResp.ID, taskResp.Status)
}
