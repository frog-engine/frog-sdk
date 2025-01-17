# SDK 开发规范

## 1. 目录结构
- **cmd/**: 存放命令行工具的实现（可选）。
- **internal/**: 存放内部实现的包，供SDK内部使用，不对外暴露。
  - **api/**: 定义SDK的接口，包括图片处理和任务管理的功能。
  - **tools/**: 封装对多个开源工具的调用。
  - **models/**: 存放数据模型，定义与图片和任务相关的数据结构。
  - **services/**: 业务逻辑层，处理图片转码、任务管理等逻辑。
  - **utils/**: 工具函数，如日志记录和配置管理。
  - **errors/**: 自定义错误类型的定义。
- **pkg/**: 存放公共包，供外部使用。
- **examples/**: 存放示例代码，展示如何使用SDK。
- **test/**: 存放测试文件，包括单元测试和集成测试。
- **go.mod**: Go模块文件，管理依赖。
- **README.md**: 项目说明文档，提供使用说明和开发指南。

## 2. API设计
- **接口定义**: 在 `internal/api/api.go` 中定义接口，提供如获取信息、文件比较、转换、压缩、设置与调整等功能。
- **实现**: 在 `internal/api/api_impl.go` 中实现这些接口，处理具体的业务逻辑。
- **参数设计**: API接口的参数应清晰明了，支持单个和多个文件的处理。

## 3. 功能实现
- **获取信息**: 使用 `ffprobe`, `mediainfo`, `ImageMagick` 提取图片和视频的基本信息。
- **文件比较**: 使用 `ffprobe`, `ImageMagick` 计算图片间的相似性。
- **转换**: 使用 `ffmpeg`, `ImageMagick`, `gif2webp` 转码为不同格式。
- **压缩**: 使用 `ffmpeg`, `ImageMagick`, `gifsicle`, `oxipng` 进行有损/无损压缩。
- **设置与调整**: 使用 `ffmpeg`, `ImageMagick`, `jpgiccplus` 进行裁剪、缩放、旋转等操作。
- **抽帧**: 使用 `ffmpeg`, `ImageMagick` 从动图或视频中抽取静态帧。
- **文本与水印**: 使用 `ImageMagick`, `delogo` 添加或删除水印。
- **滤镜与特效**: 使用 `ImageMagick` 应用滤镜效果。
- **AI增强**: 使用 `ZoomAI` 进行图片增强。
- **高级功能**: 在 `context_service.go` 中实现任务管理，包括上下文管理、任务串行与并行执行、日志记录。

## 4. 错误处理
- **自定义错误类型**: 使用自定义错误类型，提供详细的错误信息。
- **错误信息**: 错误信息应包含错误码和描述，便于调试和定位问题。

## 5. 日志记录
- **日志功能**: 提供日志记录功能，记录SDK的运行状态和错误信息。
- **日志级别**: 支持日志级别（如INFO、ERROR等），便于过滤和查看。

## 6. 配置管理
- **配置功能**: 提供配置管理功能，支持读取和管理SDK的配置。
- **环境变量**: 配置应支持环境变量和配置文件，便于灵活配置。

## 7. 示例代码
- **示例代码**: 提供示例代码，帮助用户快速上手。
- **使用场景**: 示例代码应涵盖常见的使用场景，展示SDK的主要功能。

## 8. 测试
- **单元测试**: 为每个功能模块编写单元测试，确保功能的正确性。
- **集成测试**: 编写集成测试，验证不同模块之间的协作。
- **测试覆盖率**: 确保测试覆盖率达到一定标准，保证代码质量。

## 9. 版本管理
- **版本控制**: 使用语义化版本控制（SemVer）管理SDK版本。
- **变更日志**: 维护变更日志，记录每个版本的主要变更和修复。

## 10. 贡献
- **贡献指南**: 提供贡献指南，欢迎社区参与开发和改进。
- **代码审查**: 所有提交的代码应经过审查，确保代码质量和一致性。

## 11. 许可证
- **开源许可证**: 本项目采用开源许可证，详细信息请查看 LICENSE 文件。

## 12. Go语言编程基本规范

### 12.1 命名规范
- **包名**: 使用小写字母，尽量简短且具有描述性。避免使用下划线和大写字母。
  ```go
  package imageprocessing
  ```

- **变量名**: 使用驼峰命名法（camelCase），首字母小写。变量名应具有描述性。
  ```go
  imageWidth := 800
  ```

- **常量名**: 使用全大写字母，单词之间用下划线分隔。
  ```go
  const MAX_BUFFER_SIZE = 1024
  ```

- **函数名**: 使用驼峰命名法，首字母大写。函数名应清晰描述其功能。
  ```go
  func GetImageInfo() {}
  ```

- **结构体名**: 使用驼峰命名法，首字母大写，通常为名词。
  ```go
  type ImageInfo struct {
      Format string
      Width  int
      Height int
  }
  ```

- **接口名**: 以`-er`结尾，表示该接口的功能。
  ```go
  type Reader interface {
      Read() ([]byte, error)
  }
  ```

### 12.2 格式规范
- **缩进**: 使用制表符（Tab）进行缩进，通常为4个空格的宽度。
  ```go
  func Example() {
      fmt.Println("Hello, World!")
  }
  ```

- **行长度**: 每行代码的长度应不超过80个字符，避免横向滚动。
  ```go
  fmt.Println("This is a very long line of code that exceeds the recommended length limit.")
  ```

- **空行**: 在函数之间和逻辑块之间使用空行分隔，以提高可读性。
  ```go
  func FirstFunction() {}

  func SecondFunction() {}
  ```

- **注释**: 使用完整的句子进行注释，首字母大写。注释应清晰描述代码的功能和目的。
  ```go
  // GetImageInfo retrieves the information of the specified image.
  func GetImageInfo() {}
  ```

### 12.3 参数规范
- **参数命名**: 使用描述性名称，避免使用单字母变量名（除非在循环中）。
  ```go
  func ResizeImage(sourcePath string, targetWidth int) {}
  ```

- **参数顺序**: 将相关参数放在一起，通常将输入参数放在前面，输出参数放在后面。
  ```go
  func ConvertImage(sourcePath string, targetFormat string) (string, error) {}
  ```

- **返回值**: 函数应返回错误类型的值，表示操作是否成功。返回值应在函数文档中说明。
  ```go
  func SaveImage(image ImageInfo) error {
      // Implementation
  }
  ```

### 12.4 错误处理
- **错误检查**: 每次调用可能返回错误的函数后，必须检查错误并进行处理。
  ```go
  result, err := SomeFunction()
  if err != nil {
      log.Fatalf("Error occurred: %v", err)
  }
  ```

- **错误类型**: 使用自定义错误类型时，确保错误信息清晰且具有描述性。
  ```go
  type MyError struct {
      Message string
  }
  ```

### 12.5 代码组织
- **文件结构**: 每个文件应包含相关的功能，避免将不相关的功能放在同一文件中。
  ```go
  // image.go
  package models

  type Image struct {
      // Fields
  }
  ```

- **包划分**: 将功能相关的代码组织到同一包中，确保包的功能单一且明确。
  ```go
  package api
  ```

### 12.6 代码审查
- **代码审查流程**: 所有代码提交应经过代码审查，确保代码质量和一致性。
  ```markdown
  // 提交代码后，创建Pull Request并请求审查。
  ```

- **审查标准**: 审查时应关注代码的可读性、可维护性和性能。
  ```markdown
  // 代码审查时，检查命名规范和注释完整性。
  ```

### 12.7 文档
- **文档生成**: 使用Go的文档工具生成API文档，确保文档与代码保持同步。
  ```bash
  go doc ./...
  ```

- **使用示例**: 在文档中提供使用示例，帮助用户理解如何使用SDK。
  ```go
  // Example usage of the SDK
  ```

# 入口文件示例
```go
package main

import (
    "fmt"
    "log"

    "github.com/yourusername/image-transcoder-sdk/internal/api"
)

func main() {
    // 示例：获取图片信息
    req := api.ImageInfoRequest{
        FilePaths: []string{"image1.jpg", "image2.png"},
    }

    resp, err := api.GetImageInfo(req)
    if err != nil {
        log.Fatalf("Error getting image info: %v", err)
    }

    // 打印获取的图片信息
    for _, info := range resp {
        fmt.Printf("Image: %s, Format: %s, Size: %dx%d\n", info.FilePath, info.Format, info.Width, info.Height)
    }

    // 示例：转换图片格式
    convertReq := api.ConvertRequest{
        SourcePath: "image1.jpg",
        TargetFormat: "png",
    }

    convertResp, err := api.Convert(convertReq)
    if err != nil {
        log.Fatalf("Error converting image: %v", err)
    }

    fmt.Printf("Converted image saved at: %s\n", convertResp.OutputPath)
}
```