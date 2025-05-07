package command

import (
  "bytes"
  "encoding/json"
  "fmt"
  "os"
  "os/exec"
)

// ReadImageInfo 使用 ffprobe 获取图像的完整信息
func ReadImageInfo(imageData []byte) (map[string]interface{}, error) {
  // 创建临时文件
  tmpFile, err := os.CreateTemp("", "image-*.jpg") // 你可以根据实际情况改为 png、webp 等
  if err != nil {
    return nil, fmt.Errorf("failed to create temp file: %w", err)
  }
  defer os.Remove(tmpFile.Name()) // 清理临时文件

  // 写入数据
  if _, err := tmpFile.Write(imageData); err != nil {
    return nil, fmt.Errorf("failed to write image data to temp file: %w", err)
  }
  tmpFile.Close()

  // 准备执行 ffprobe
  cmd := exec.Command("ffprobe",
    "-v", "error",
    "-select_streams", "v:0",
    "-show_entries",
    "stream=width,height,codec_name,codec_type,pix_fmt,color_space,color_transfer,color_primaries,field_order,r_frame_rate",
    "-show_format",
    "-of", "json",
    tmpFile.Name(),
  )

  var out bytes.Buffer
  cmd.Stdout = &out
  if err := cmd.Run(); err != nil {
    return nil, fmt.Errorf("failed to run ffprobe: %w", err)
  }

  // 解析 JSON 输出
  var result map[string]interface{}
  if err := json.Unmarshal(out.Bytes(), &result); err != nil {
    return nil, fmt.Errorf("failed to parse ffprobe output: %w", err)
  }

  return result, nil
}

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
