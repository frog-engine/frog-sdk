package models

// ImageInfoRequest 请求结构体
type ImageInfoRequest struct {
  FilePaths []string // 支持单个或多个文件
}

// ImageInfoResponse 响应结构体
type ImageInfoResponse struct {
  FilePath string
  Format   string
  Width    int
  Height   int
}

// ConvertRequest 转换请求结构体
type ConvertRequest struct {
  SourcePath   string // 源文件路径
  TargetFormat string // 目标格式（如 png, jpeg）
}

// ConvertResponse 转换响应结构体
type ConvertResponse struct {
  OutputPath string // 输出文件路径
}

// TaskResponse 转换响应结构体
type TaskResponse struct {
  ID         string // 任务ID
  Status     string // 任务状态
  Error      string // 错误信息
  OutputPath string // 输出文件路径
}
