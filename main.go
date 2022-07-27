package main

import (
	"github.com/rogeriofontes/api-go-gin/api/controller"
	"github.com/rogeriofontes/api-go-gin/api/repository"
	"github.com/rogeriofontes/api-go-gin/api/routes"
	"github.com/rogeriofontes/api-go-gin/api/service"
	"github.com/rogeriofontes/api-go-gin/database"
	"github.com/rogeriofontes/api-go-gin/infrastructure"
	"github.com/rogeriofontes/api-go-gin/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NewGinRouter()
	db := database.NewDatabase()

	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	postRoute := routes.NewPostRoute(postController, router)
	postRoute.Setup()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	db.DB.AutoMigrate(&models.Post{}, &models.User{})
	router.Gin.Run(":8000")
}
