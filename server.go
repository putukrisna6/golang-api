package main

import (
	"github.com/gin-gonic/gin"
	"github.com/putukrisna6/golang-api/config"
	"github.com/putukrisna6/golang-api/controller"
	"github.com/putukrisna6/golang-api/middleware"
	"github.com/putukrisna6/golang-api/repository"
	"github.com/putukrisna6/golang-api/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)

	jwtService  service.JWTService  = service.NewJWTService()
	authService service.AuthService = service.NewAuthService(userRepository)
	userService service.UserService = service.NewUserService(userRepository)
	bookService service.BookService = service.NewBookService(bookRepository)

	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("profile", userController.Get)
		userRoutes.PUT("/", userController.Update)
	}

	bookRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.Get)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World",
		})
	})
	r.Run()
}
