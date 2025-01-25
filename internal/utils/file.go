package utils

import (
  "fmt"
  "os"

  "github.com/frog-engine/frog-sdk/pkg/logger"
)

func GetHomeDir() string {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    logger.Error(fmt.Errorf("failed to get home directory: %w", err))
    return "~"
  }
  return homeDir
}
