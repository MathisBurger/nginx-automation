package main

import (
	"github.com/MathisBurger/nginx-automation/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowedOrigins,
	}))

	app.Use(logger.New())
	InitReverseProxy(app)
	app.Listen(":" + cfg.ApplicationPort)

}

func InitReverseProxy(app *fiber.App) {

}
