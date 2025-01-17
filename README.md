# frog-sdk 介绍
Image Transcoder SDK 是一个用于图片和视频处理的Go语言SDK，支持多种开源工具的调用，提供丰富的API接口，满足多种业务需求。

## 服务分层
```bash
业务层（按需生产/预生产） ->  服务层(同步/异步) -> 图片统一转码SDK -> 转码工具
```

## 主要功能
1. 为上层提供通过SDK接口，封装工具命令
2. 调用转码工具ffpg/ImageMagick/Webpmux等进行转码
3. 提供开发上下文环境，提供任务回调，并能对一个对象连续操作

## 具体功能
- 获取图片和视频的基本信息及元数据
- 图片间的相似性比较
- 图片格式转换
- 图片压缩（有损/无损）
- 图片裁剪、缩放、旋转等设置与调整
- 从动图或视频中抽取静态帧
- 添加文字和水印
- 应用滤镜和特效
- AI增强功能
- 任务管理（上下文管理、任务串行与并行执行）

### 安装
1. 克隆项目
```bash
$ git clone https://github.com/frog-engine/frog-sdk.git
```
2. 进入项目目录
```bash
$ cd frog-sdk
```

3. 安装依赖
```bash
$ go mod download
$ go mod tidy
$ go mod verify
```

4. 启动服务
```bash
$ go run cmd/frog-go/main.go
# 看见启动信息
# frog-sdk run successfully
```

## 使用示例
```go
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
```