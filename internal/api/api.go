package api

import (
  "image"

  "github.com/frog-engine/frog-sdk/config"
  "github.com/frog-engine/frog-sdk/internal/models"
)

// Api 通用接口定义，处理图片相关的业务逻辑
type Api interface {
  Name() string
  GetConfig() *config.Config
  ReadImageBlob(imageData []byte) (image.Image, error)
  GetImageInfo(req *models.ImageInfoRequest) ([]*models.ImageInfoResponse, error)
  ImageConvert(req *models.ConvertRequest) (*models.ConvertResponse, error)
  GetTaskInfo(taskID string) (*models.TaskResponse, error)
}
