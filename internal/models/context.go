package models

// Context 定义上下文管理的数据结构
type Context struct {
  PreviousResult interface{}
  TaskID         string
}
