package cmd

import (
  "fmt"
  "os/exec"
)

// FFprobeGetImageDimensions 获取图片的宽度和高度
func FFprobeGetImageDimensions(filePath string) (int, int, error) {
  cmd := exec.Command("ffprobe",
    "-v", "error",
    "-select_streams", "v:0",
    "-show_entries", "stream=width,height",
    "-of", "default=noprint_wrappers=1:nokey=1",
    filePath)

  output, err := cmd.Output()
  if err != nil {
    return 0, 0, fmt.Errorf("failed to probe image %s: %w", filePath, err)
  }

  var width, height int
  fmt.Sscanf(string(output), "%d\n%d", &width, &height)

  return width, height, nil
}
