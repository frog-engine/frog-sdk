# frog-sdk 介绍
Image Transcoder SDK 是一个用于图片和视频处理的Go语言SDK，支持多种开源工具的调用，提供丰富的API接口，满足多种业务需求。

## 服务分层
```bash
业务层（即时转码/异步转码） ->  服务层(同步/异步) -> 图片转码SDK -> 开源转码工具
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
$ go mod vendor
$ go mod verify
```

4. 启动服务
```bash
# SDK包供上层引用，此处启动仅是验证
$ go run cmd/frog-sdk/main.go
# 看见启动信息
# frog-sdk run successfully
```

## 使用示例
```go

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
```