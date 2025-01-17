package test

import (
  "frog-sdk/internal/models"
  "testing"
)

func TestContext(t *testing.T) {
  ctx := models.Context{
    PreviousResult: nil,
    TaskID:         "task_1",
  }

  if ctx.TaskID != "task_1" {
    t.Fatalf("Expected TaskID to be 'task_1', got %s", ctx.TaskID)
  }
}
