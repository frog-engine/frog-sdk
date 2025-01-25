package test

import (
  "testing"

  "github.com/frog-engine/frog-sdk/internal/models"
)

func TestContext(t *testing.T) {
  ctx := models.Context{
    TaskID: "task_1",
  }

  if ctx.TaskID != "task_1" {
    t.Fatalf("Expected TaskID to be 'task_1', got %s", ctx.TaskID)
  }
}
