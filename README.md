# Go-todo
## 项目介绍
这是一个基于Go语言的简单的任务管理系统，使用了Gin框架和Gorm框架，使用了MySQL数据库。

## 快速开始
先创建一个数据库，数据库名称为 `go_todo`
> 名称应与config/config.yaml中的database名称一致

```bash
go run main.go
# 或者
go build
./go-todo.exe
```
访问 `http://localhost:8080` 即可查看项目界面。
