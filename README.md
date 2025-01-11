# frog-sdk 介绍
Frog-SDK is an image transcoding SDK that encapsulates tools such as FFmpeg, ImageMagick, and Webpmux, providing an interface for upper-layer calls. It supports synchronous/asynchronous operations and multi-task processing.

## 服务分层
```bash
业务层（按需生产/预生产） ->  服务层(同步/异步) -> 图片统一转码SDK -> 转码工具
```
## 主要功能
1. 为上层提供通过SDK接口，封装工具命令
2. 调用转码工具ffpg/ImageMagick/Webpmux等进行转码
3. 提供开发上下文环境，提供任务回调，并能对一个对象连续操作

## 需求描述
// TODO
