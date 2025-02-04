package routes

import (
	"todo-list-api/controllers"
	"todo-list-api/middlewares"
	"todo-list-api/repository"
	"todo-list-api/services"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes sets up all API routes.
func RegisterRoutes(r *gin.Engine) {
	// Initialize repositories.
	userRepo := repository.NewUserRepository()
	todoRepo := repository.NewTodoRepository()

	// Initialize services.
	authService := services.NewAuthService(userRepo)
	todoService := services.NewTodoService(todoRepo)

	// Initialize controllers.
	authController := controllers.NewAuthController(authService)
	todoController := controllers.NewTodoController(todoService)

	// Public routes.
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	// Protected routes (require JWT).
	authRoutes := r.Group("/")
	authRoutes.Use(middlewares.JWTAuthMiddleware())
	{
		authRoutes.POST("/todos", todoController.CreateTodo)
		authRoutes.PUT("/todos/:id", todoController.UpdateTodo)
		authRoutes.DELETE("/todos/:id", todoController.DeleteTodo)
		authRoutes.GET("/todos", todoController.GetTodos)
	}

	// Uncomment to serve Swagger docs.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
