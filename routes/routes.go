package routes

import (
	"github.com/djchanahcjd/go-todo/handlers"
	"github.com/djchanahcjd/go-todo/services"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

// 测试拦截器
func TestInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("test", "test")
		c.Next()
	}
}

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	// 创建默认的gin路由引擎
	router := gin.Default()
	// favicon
	router.Use(favicon.New("./favicon.ico"))

	// 创建服务和处理器实例
	todoService := services.NewTodoService()
	todoHandler := handlers.NewTodoHandler(todoService)

	// API路由组
	api := router.Group("/api")
	{
		// Todo相关路由
		todos := api.Group("/todos")
		{
			todos.POST("/", todoHandler.CreateTodo)
			todos.GET("/", todoHandler.GetAllTodos)
			todos.GET("/:id", todoHandler.GetTodo)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
			todos.PUT("/:id/complete", todoHandler.CompleteTodo)
		}
	}
	// 健康检查
	router.GET("/healthz", TestInterceptor(), func(c *gin.Context) {
		test := c.MustGet("test").(string)
		c.JSON(200, gin.H{"status": "ok " + test + "\n"})
	})

	// 加载HTML模板
	router.LoadHTMLGlob("templates/*")

	// Web界面路由
	router.GET("/", todoHandler.TodoIndex)

	return router
}
