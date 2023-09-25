package main

import (
	"online-articles/client/restclient"
	"online-articles/configuration"
	"online-articles/controller"
	_ "online-articles/docs"
	"online-articles/exception"
	repository "online-articles/repository/impl"
	service "online-articles/service/impl"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Online Article
// @version 1.0.0
// @description Online Article
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description Authorization For JWT
func main() {
	//setup configuration
	config := configuration.New()
	database := configuration.NewDatabase(config)
	configuration.MinioBucket(config)

	//repository
	userRepository := repository.NewUserRepositoryImpl(database)
	articleRepository := repository.NewArticleRepositoryImpl(database)

	//rest client
	httpBinRestClient := restclient.NewHttpBinRestClient()

	//service
	userService := service.NewUserServiceImpl(&userRepository)
	articleService := service.NewArticleServiceImpl(&articleRepository)
	httpBinService := service.NewHttpBinServiceImpl(&httpBinRestClient)

	//controller
	userController := controller.NewUserController(&userService, config)
	articleController := controller.NewArticleController(&articleService, config)
	httpBinController := controller.NewHttpBinController(&httpBinService)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	//routing
	userController.Route(app)
	articleController.Route(app)
	httpBinController.Route(app)

	//swagger
	app.Get("/online-articles/*", swagger.HandlerDefault)

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
