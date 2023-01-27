package main

import (
	"discchatgpt/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// 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
func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot OTg3MDYwNDQxNzczNjU4MTcz.G5OcuP.2Z4mYV2aCvy_fT_jCyfg75---Qf-IgjsKnnDO4")
	if err != nil {
		panic(err)
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Select the intents
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	// message, err := api.RequestToChatGPT("Write a hello world program (javascript)", "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")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(message)

	// Start the bot
	err = dg.Open()
	if err != nil {
		panic(err)
	}
	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.", "info")

	// Register Commands
	handlers.InitCommands(dg)
	// Remove Commands
	// command.RemoveCommands(dg)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sig

	// Stop the bot
	defer dg.Close()

	// defer handlers.RemoveCommands(dg)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
