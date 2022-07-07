package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateHTTPServer(port string) *fiber.App {
	connection := fiber.New()

	connection.Use(logger.New(logger.ConfigDefault))

	return connection
}

func handleRequests(app *fiber.App) {

}

func Connect(port string, db mongo.Database) {
	conn := CreateHTTPServer(port)

	// Here is the place where you can add your Repositories, Services and Controllers.

	handleRequests(conn)

	if err := conn.Listen(":" + port); err != nil {
		panic(err)
	}
}
