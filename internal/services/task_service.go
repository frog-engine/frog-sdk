package services

import (
  "frog-sdk/internal/models"
)

// TaskService，管理任务的调度和执行
type TaskService struct {
}

// NewTaskService 创建新的任务服务
func NewTaskService() *TaskService {
  return &TaskService{}
}

// GetImageInfo 获取图片信息
func (s *TaskService) GetTaskInfo(taskID string) (*models.TaskResponse, error) {
  // TODO: 实现获取任务信息的逻辑
  return nil, nil
}

// GetImageInfo 获取图片信息
func (s *TaskService) GetImageInfo(req models.ImageInfoRequest) ([]models.ImageInfoResponse, error) {
  // TODO: 实现获取图片信息的逻辑
  return nil, nil
}

// Convert 转换图片格式
func (s *TaskService) Convert(req models.ConvertRequest) (models.ConvertResponse, error) {
  // TODO: 实现图片格式转换的逻辑
  return models.ConvertResponse{OutputPath: "output.png"}, nil
}
