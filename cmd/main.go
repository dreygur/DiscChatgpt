package main

import (
	"discchatgpt/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	dg.Client.Timeout = time.Minute * 20
	dg.MaxRestRetries = 100

	// Select the intents
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	// Start the bot
	err = dg.Open()
	if err != nil {
		panic(err)
	}
	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")

	// Register Commands
	handlers.InitCommands(dg)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sig

	// Stop the bot
	defer dg.Close()

	// Remove Commands
	// defer handlers.RemoveCommands(dg)
}
