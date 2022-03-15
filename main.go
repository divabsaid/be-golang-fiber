package main

import (
	"be-golang-fiber/app"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//init config
	config := app.InitConfig()

	// init database
	dbConn := app.InitDatabase(config)

	f := fiber.New()

	// Middleware
	f.Use(logger.New())
	f.Use(recover.New())
	// CORS
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "OPTIONS, GET, POST, PUT, DELETE",
	}))

	// init entity
	app.InitEntity(f, dbConn)

	// Init logger
	zerolog.TimestampFieldName = "timestamp"
	zerolog.MessageFieldName = "msg"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log.Fatal(f.Listen(":3000"))
}
