package frogsdk

import (
  "sync"

  "github.com/frog-engine/frog-sdk/config"
  "github.com/frog-engine/frog-sdk/internal/api"
  "github.com/frog-engine/frog-sdk/pkg/logger"
)

var (
  globalAPI api.Api
  once      sync.Once
)

// Init 初始化SDK并返回ImageApi实例
func Init(newConf *config.Config, evn string) api.Api {
  globalConfig := config.InitConfig(evn)
  if newConf != nil {
    // 根据传入的conf更新theConf
    if newConf.ToolsRoot != "" {
      globalConfig.ToolsRoot = newConf.ToolsRoot
    }
    if newConf.FFMpegPath != "" {
      globalConfig.FFMpegPath = newConf.FFMpegPath
    }
    if newConf.ImageMagickPath != "" {
      globalConfig.ImageMagickPath = newConf.ImageMagickPath
    }
    if newConf.DefaultQuality != 0 {
      globalConfig.DefaultQuality = newConf.DefaultQuality
    }
    if newConf.TempFilePath != "" {
      globalConfig.TempFilePath = newConf.TempFilePath
    }
    if newConf.Environment != "" {
      globalConfig.Environment = newConf.Environment
    }
  }

  // 创建SDKApi实例
  globalAPI = api.InitImageApi(globalConfig)
  logger.Println("frogsdk initialized.")
  return globalAPI
}

func Name() string {
  return globalAPI.Name()
}

func GetAPI() api.Api {
  // 使用 sync.Once 确保 SDK 只初始化一次
  once.Do(func() {
    if globalAPI == nil {
      globalAPI = Init(nil, "")
    }
  })
  logger.Println("frogsdk has started.")
  return globalAPI
}

func Terminate() {
  logger.Info("stop frogsdk.")
}
