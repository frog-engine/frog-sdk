package api

import (
  "image"

  "github.com/frog-engine/frog-sdk/config"
  "github.com/frog-engine/frog-sdk/internal/models"
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
func (s *ImageApi) GetImageInfo(req *models.ImageInfoRequest) ([]*models.ImageInfoResponse, error) {
  return s.imageService.GetImageInfo(req)
}

func (s *ImageApi) ReadImageBlob(imageData []byte) (image.Image, error) {
  return s.imageService.ReadImageBlob(imageData)
}

// ImageConvert 转换图片格式
func (s *ImageApi) ImageConvert(req *models.ConvertRequest) (*models.ConvertResponse, error) {
  return s.imageService.ImageTranscode(req)
}

// GetTaskInfo 获取任务信息
func (s *ImageApi) GetTaskInfo(taskID string) (*models.TaskResponse, error) {
  // 调用任务服务获取任务信息
  return s.taskService.GetTaskInfo(taskID)
}
