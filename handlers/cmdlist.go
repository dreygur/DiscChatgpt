package handlers

import (
	"github.com/bwmarrin/discordgo"
)

// RegisteredCommands is a slice of ApplicationCommand
var RegisteredCommands []*discordgo.ApplicationCommand

// Commands
var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "help",
		Description: "Shows help message",
	},
	{
		Name:        "voldus",
		Description: "ChatGPT",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "prompt",
				Description: "Enter your prompt",
				Required:    true,
			},
		},
	},
}
