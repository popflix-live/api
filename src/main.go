package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/popflix-live/api/src/application"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	// Database connection will be added later
	// dbClient, err := db.ConnectDB()
	// helper.ErrorPanic(err)
	// defer dbClient.Prisma.Disconnect()

	app := application.New()

	err = app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
