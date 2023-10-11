// Package main is the entry point of the application.
package main

// Importing necessary packages.
import (
	"log"
	"os"
	"rabbitx-client/bot"
	"rabbitx-client/client"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// marketID is the ID of the market to be monitored.
var marketID = "ETH-USD"

// main is the main function of the application.
// It loads environment variables, initializes the client and the bot, and runs the bot.
func main() {
	// Load environment variables from .env file.
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Parse API_KEY_EXPIRED environment variable.
	keyExpired, err := strconv.ParseInt(os.Getenv("API_KEY_EXPIRED"), 10, 32)
	if err != nil {
		log.Fatalf("Failed to convert API_KEY_EXPIRED to int: %s", err)
	}

	// Initialize the client with environment variables.
	rbClient := client.NewRbClient(
		os.Getenv("API_URL"),
		os.Getenv("WALLET"),
		os.Getenv("PRIVATE_KEY"),
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
		os.Getenv("REFRESH_TOKEN"),
		os.Getenv("PRIVATE_JWT"),
		keyExpired)

	logrus.Info("Client successfully created")

	// Update secrets.
	_, _, jwtPrivate, err := rbClient.GetSecrets()
	if err != nil {
		log.Fatalf("Failed to get secrets: %s", err)
	}

	// Save new secrets.
	if err := rbClient.SaveSecrets("./updated_secrets.json"); err != nil {
		log.Fatalf("Failed to save secrets: %s", err)
	}

	// Initialize and run the bot.
	rbBot := bot.NewBot(rbClient, os.Getenv("WS_URL"), jwtPrivate)
	if err := rbBot.Run(marketID); err != nil {
		log.Fatalf("Failed to run bot: %s", err)
	}

	logrus.Info("Bot successfully launched")

	// TODO: Launch webserver here to see the bot stats.

	// Exit on ctrl+c.
	select {}
}
