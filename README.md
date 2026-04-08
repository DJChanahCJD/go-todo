# Go Todo 任务管理系统

一个基于 Go 语言开发的轻量级任务管理 Web 应用，使用 Gin 框架和 GORM ORM，数据库采用 MySQL。

## 技术栈

- **后端**：Go 1.24 + Gin
- **数据库**：MySQL + GORM
- **前端**：Bootstrap 5 + HTML/JS

## 项目结构

```
go-todo/
├── config/          # 配置文件
│   ├── config.go   # 配置结构体
│   └── config.yaml # 配置文件
├── database/       # 数据库初始化
│   └── database.go
├── handlers/       # HTTP 处理器
│   └── todo_handler.go
├── models/         # 数据模型
│   └── todo.go
├── routes/         # 路由配置
│   └── routes.go
├── services/       # 业务逻辑
│   └── todo_service.go
├── templates/      # HTML 模板
│   ├── index.html
│   └── layout.html
├── main.go         # 入口文件
└── go.mod          # 依赖管理
```

## 功能特性

- ✅ 创建任务
- ✅ 查看任务列表
- ✅ 编辑任务
- ✅ 删除任务
- ✅ 标记任务完成
- ✅ RESTful API
- ✅ Web 界面

## 快速开始

### 1. 准备 MySQL 数据库

```sql
CREATE DATABASE go_todo;
```

### 2. 配置数据库连接

编辑 `config/config.yaml` 文件，修改 MySQL 连接配置：

```yaml
mysql:
  host: localhost
  port: 3306
  database: go_todo
  username: root
  password: your_password

server:
  port: 8080
```

### 3. 运行项目

```bash
# 直接运行
go run main.go

# 或编译后运行
go build
./go-todo.exe
```

### 4. 访问应用

- Web 界面：http://localhost:8080
- API 端点：http://localhost:8080/api/todos

## API 接口

| 方法   | 路径                  | 描述           |
|--------|---------------------|----------------|
| GET    | /api/todos          | 获取所有任务   |
| GET    | /api/todos/:id      | 获取指定任务   |
| POST   | /api/todos          | 创建新任务     |
| PUT    | /api/todos/:id      | 更新任务       |
| PUT    | /api/todos/:id/complete | 标记完成   |
| DELETE | /api/todos/:id      | 删除任务       |

### 请求示例

**创建任务**
```json
POST /api/todos
{
  "title": "完成任务",
  "description": "这是一个示例任务",
  "due_date": "2024-12-31T00:00:00Z"
}
```

## 环境要求

- Go 1.24+
- MySQL 5.7+
