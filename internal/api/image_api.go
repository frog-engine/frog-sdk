package api

import (
  "image"

  "github.com/frog-engine/frog-sdk/config"
  internalmodels "github.com/frog-engine/frog-sdk/internal/models"
  pkgmodels "github.com/frog-engine/frog-sdk/pkg/models"
  "github.com/frog-engine/frog-sdk/internal/services"
)

// ImageApi 实现 Api 接口
type ImageApi struct {
  config       *config.Config
  imageService *services.ImageService
  taskService  *services.TaskService
}

func (a *ImageApi) Name() string {
  return "frog-sdk-image-api"
}

// InitImageApi 创建新的 ImageApi 实例，传入Config实例
func InitImageApi(config *config.Config) *ImageApi {
  return &ImageApi{
    config:       config,
    imageService: services.NewImageService(),
    taskService:  services.NewTaskService(),
  }
}

// NewImageApi 创建新的 ImageApi 实例，默认使用全局配置
func NewImageApi() *ImageApi {
  return &ImageApi{
    config:       config.InitConfig(""),
    imageService: services.NewImageService(),
    taskService:  services.NewTaskService(),
  }
}

func (a *ImageApi) GetConfig() *config.Config {
  return a.config
}

// GetImageInfo 获取图片基本信息
func (s *ImageApi) GetImageInfo(req *internalmodels.ImageInfoRequest) ([]*internalmodels.ImageInfoResponse, error) {
  return s.imageService.GetImageInfo(req)
}

func (s *ImageApi) ReadImageInfo(imageData []byte) (map[string]interface{}, error) {
  return s.imageService.ReadImageInfo(imageData)
}

func (s *ImageApi) ReadImageBlob(imageData []byte) (image.Image, error) {
  return s.imageService.ReadImageBlob(imageData)
}

// ImageConvert 转换图片格式
func (s *ImageApi) ImageConvert(req *internalmodels.ConvertRequest) (*internalmodels.ConvertResponse, error) {
  return s.imageService.ImageTranscode(req)
}

// ProcessImage 同步处理图片字节流
func (s *ImageApi) ProcessImage(imageData []byte, req *pkgmodels.ProcessImageRequest) ([]byte, error) {
  return s.imageService.ProcessImage(imageData, req)
}

// GetTaskInfo 获取任务信息
func (s *ImageApi) GetTaskInfo(taskID string) (*internalmodels.TaskResponse, error) {
  // 调用任务服务获取任务信息
  return s.taskService.GetTaskInfo(taskID)
}
