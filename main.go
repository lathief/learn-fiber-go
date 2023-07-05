package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/lathief/learn-fiber-go/app/category"
	"github.com/lathief/learn-fiber-go/app/product"
	"github.com/lathief/learn-fiber-go/pkg/configs"
	"github.com/lathief/learn-fiber-go/pkg/middleware"
	"github.com/lathief/learn-fiber-go/pkg/utils"
	"github.com/lathief/learn-fiber-go/platform/database"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	config := configs.FiberConfig()
	db, err := database.OpenDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New(config)
	middleware.FiberMiddleware(app)
	productRouter := product.NewRouter(db)
	productRouter.Handle(app)
	categoryRouter := category.NewRouter(db)
	categoryRouter.Handle(app)
	// Build fiber connection URL
	fiberConnURL, _ := utils.ConnectionURLBuilder("fiber")
	// Run server
	if err = app.Listen(fiberConnURL); err != nil {
		log.Printf("Server is not running! Error: %v", err)
	}
}
