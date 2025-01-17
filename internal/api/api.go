package api

import (
  "frog-sdk/internal/models"
  "frog-sdk/internal/services"
)

// SDKApi 通用接口定义，处理图片相关的业务逻辑
type SDKApi interface {
  GetImageInfo(req models.ImageInfoRequest) ([]models.ImageInfoResponse, error)
  Convert(req models.ConvertRequest) (models.ConvertResponse, error)
  GetTaskInfo(taskID string) (*models.TaskResponse, error)
}

// SDKApiImpl 实现 SDKApi 接口
type SDKApiImpl struct {
  imageService *services.ImageService
  taskService  *services.TaskService
}

// NewImageApi 创建新的 ImageApi 实例
func NewSDKApi() *SDKApiImpl {
  return &SDKApiImpl{
    imageService: services.NewImageService(),
    taskService:  services.NewTaskService(),
  }
}

// GetImageInfo 获取图片基本信息
func (s *SDKApiImpl) GetImageInfo(req models.ImageInfoRequest) ([]models.ImageInfoResponse, error) {
  return s.imageService.GetImageInfo(req)
}

// Convert 转换图片格式
func (s *SDKApiImpl) ImageConvert(req models.ConvertRequest) (models.ConvertResponse, error) {
  return s.imageService.ImageTranscode(req)
}

// GetTaskInfo 获取任务信息
func (s *SDKApiImpl) GetTaskInfo(taskID string) (*models.TaskResponse, error) {
  // 调用任务服务获取任务信息
  return s.taskService.GetTaskInfo(taskID)
}
