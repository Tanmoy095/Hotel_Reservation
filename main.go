package main

import (
	"context"
	"flag"
	"log"

	"github.com/Tanmoy095/Hotel_Reservation.git/api"
	"github.com/Tanmoy095/Hotel_Reservation.git/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBURI = "mongodb://localhost:27017"

var config = fiber.Config{

	ErrorHandler: func(c *fiber.Ctx, err error) error {

		// Return from handler
		return c.JSON(map[string]string{"error": err.Error()})

	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the api server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal(err)
	}

	//with bson.M we are going to specify query
	// handler initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	app := fiber.New(config)
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user", userHandler.HandleGetUser)
	apiV1.Get("/user", userHandler.HandlePostUser)
	apiV1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}
