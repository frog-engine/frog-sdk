package test

import (
  "frog-sdk/internal/models"
  "frog-sdk/internal/services"
  "testing"
)

func TestImageService_GetImageInfo(t *testing.T) {
  service := services.NewImageService()
  req := models.ImageInfoRequest{
    FilePaths: []string{"image1.jpg"},
  }

  resp, err := service.GetImageInfo(req)
  if err != nil {
    t.Fatalf("Expected no error, got %v", err)
  }

  if len(resp) != 1 {
    t.Fatalf("Expected 1 response, got %d", len(resp))
  }
}
