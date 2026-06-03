package api

import (
  "image"

  "github.com/frog-engine/frog-sdk/config"
  internalmodels "github.com/frog-engine/frog-sdk/internal/models"
  pkgmodels "github.com/frog-engine/frog-sdk/pkg/models"
)

// Api 通用接口定义，处理图片相关的业务逻辑
type Api interface {
  Name() string
  GetConfig() *config.Config
  ReadImageBlob(imageData []byte) (image.Image, error)
  ReadImageInfo(imageData []byte) (map[string]interface{}, error)
  GetImageInfo(req *internalmodels.ImageInfoRequest) ([]*internalmodels.ImageInfoResponse, error)
	ImageConvert(req *internalmodels.ConvertRequest) (*internalmodels.ConvertResponse, error)
	ProcessImage(imageData []byte, req *pkgmodels.ProcessImageRequest) ([]byte, error)
	GetTaskInfo(taskID string) (*internalmodels.TaskResponse, error)
}
