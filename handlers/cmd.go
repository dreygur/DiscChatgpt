package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Register slash commands to discord
func InitCommands(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	log.Println("Adding commands...")
	RegisteredCommands = make([]*discordgo.ApplicationCommand, len(Commands))

	for i, v := range Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", cmd.Name, err)
		}
		RegisteredCommands[i] = cmd
		log.Println("Added command: " + cmd.Name)
	}
	log.Println("Added commands...")
}

// Remove already registered commands from discord
// while shutting down gracefully.
func RemoveCommands(s *discordgo.Session) {
	log.Println("Removing commands...")
	for _, v := range RegisteredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
		log.Println("Removed command: " + v.Name)
	}
}

// Commands list
var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"help":   Help,
	"voldus": Voldus,
}
