package main

import (
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/api/route"
	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the app
	app := bootstrap.App()

	// Get the environment variables
	env := app.Env

	// Connect to the database
	db := app.Mongo.Database(env.DBName)

	// Close the database connection when the main function is done
	defer app.CloseDBConnection()

	// Set the timeout for the context of the request
	timeout := time.Duration(env.ContextTimeout) * time.Second

	// Initialize the gin
	gin := gin.Default()

	// Setup the routes
	route.Setup(env, timeout, db, gin)

	// Run the server
	gin.Run(env.ServerAddress)
}
