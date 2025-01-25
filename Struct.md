# SDK 目录结构
```md
frog-sdk/
├── init.go                     # 对外API入库函数
├── cmd/                        # 主应用程序目录（可选）
│   └── frog-sdk/               # SDK的命令行工具
│       └── main.go             # 入口文件
├── internal/                   # 内部实现包
│   ├── api/                    # API接口定义
│   │   ├── api.go              # API接口定义
│   │   └── api_impl.go         # API实现
│   ├── tools/                  # 工具调用封装
│   │   ├── ffmpeg.go           # FFmpeg工具调用封装
│   │   ├── imagemagick.go      # ImageMagick工具调用封装
│   │   ├── svt_av1.go          # SVT-AV1工具调用封装
│   │   ├── libvips.go          # libvips工具调用封装
│   │   ├── gifsicle.go         # Gifsicle工具调用封装
│   │   ├── oxipng.go           # Oxipng工具调用封装
│   │   ├── mozjpeg.go          # MozJPEG工具调用封装
│   │   └── waifu2x.go          # Waifu2x工具调用封装
│   ├── models/                 # 数据模型
│   │   ├── image.go            # 图片相关数据结构
│   │   ├── task.go             # 任务相关数据结构
│   │   └── context.go          # 上下文管理数据结构
│   ├── services/               # 业务逻辑层
│   │   ├── image_service.go     # 图片处理服务
│   │   ├── task_service.go      # 任务管理服务
│   │   ├── ai_service.go        # AI增强服务（可选）
│   │   └── context_service.go    # 上下文管理服务
│   ├── utils/                  # 工具函数
│   │   ├── logger.go           # 日志工具
│   │   └── config.go           # 配置管理
│   └── errors/                 # 错误处理
│       └── errors.go           # 自定义错误类型
├── pkg/                        # 公共包（供外部使用）
│   ├── client/                 # SDK客户端
│   │   └── client.go           # SDK客户端实现
│   ├── types/                  # 公共类型定义
│   │   └── types.go            # 公共类型
│   └── middleware/             # 中间件（如日志、监控等）
│       └── logging.go          # 日志中间件
├── examples/                   # 示例代码
│   └── basic_usage.go          # 基本用法示例
├── test/                       # 测试文件
│   ├── api_test.go             # API测试
│   ├── service_test.go         # 服务测试
│   ├── context_test.go         # 上下文管理测试
│   └── integration_test.go      # 集成测试
├── go.mod                      # Go模块文件
└── README.md                   # 项目说明文档
```

## 结构设计说明
- 该结构支持模块化设计，便于未来功能的扩展和维护。
- 可以根据需要在 internal/ 或 pkg/ 中添加新的功能模块。
- 通过 utils/ 和 middleware/ 可以轻松添加通用工具和中间件。
- 该结构支持模块化设计，便于未来功能的扩展和维护。
- 提供清晰的API接口和示例代码，SDK可以方便地被其他语言的绑定或封装使用。

## 功能实现说明
1. 获取信息: 使用 ffprobe, mediainfo, ImageMagick 提取图片和视频的基本信息。
2. 文件比较: 使用 ffprobe, ImageMagick 计算图片间的相似性。
3. 转换: 使用 ffmpeg, ImageMagick, gif2webp 转码为不同格式。
压缩: 使用 ffmpeg, ImageMagick, gifsicle, oxipng 进行有损/无损压缩。
5. 设置与调整: 使用 ffmpeg, ImageMagick, jpgiccplus 进行裁剪、缩放、旋转等操作。
6. 抽帧: 使用 ffmpeg, ImageMagick 从动图或视频中抽取静态帧。
7. 文本与水印: 使用 ImageMagick, delogo 添加或删除水印。
滤镜与特效: 使用 ImageMagick 应用滤镜效果。
9. AI增强: 使用 ZoomAI 进行图片增强。
10. 高级功能: 在 context_service.go 中实现任务管理，包括上下文管理、任务串行与并行执行、日志记录。
扩展性与适配性

## API接口示例
```go
package api

// ImageInfoRequest 请求结构体
type ImageInfoRequest struct {
    FilePaths []string // 支持单个或多个文件
}

// ImageInfoResponse 响应结构体
type ImageInfoResponse struct {
    // 图片信息
}

// GetImageInfo 获取图片基本信息
func GetImageInfo(req ImageInfoRequest) (ImageInfoResponse, error)

// CompareImagesRequest 请求结构体
type CompareImagesRequest struct {
    Image1 string
    Image2 string
}

// CompareImagesResponse 响应结构体
type CompareImagesResponse struct {
    Similarity float64 // SSIM或PSNR值
}

// CompareImages 比较两张图片
func CompareImages(req CompareImagesRequest) (CompareImagesResponse, error)

// ConvertRequest 请求结构体
type ConvertRequest struct {
    SourcePath string
    TargetFormat string
}

// ConvertResponse 响应结构体
type ConvertResponse struct {
    OutputPath string
}

// Convert 转码为不同格式
func Convert(req ConvertRequest) (ConvertResponse, error)

// CompressRequest 请求结构体
type CompressRequest struct {
    FilePath string
    Lossy bool // 是否有损压缩
}

// CompressResponse 响应结构体
type CompressResponse struct {
    OutputPath string
}

// Compress 压缩图片
func Compress(req CompressRequest) (CompressResponse, error)

// 其他裁切, 缩放, 去水印等函数
```
