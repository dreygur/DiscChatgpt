package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func HelpMessageHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	forAdmin := []*discordgo.MessageEmbed{
		{
			Title:       "Help",
			Description: "This is a help message",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "`/help`",
					Value:  "Shows this help message",
					Inline: true,
				},
				{
					Name:   "`/voldus`",
					Value:  "ChatGPT",
					Inline: false,
				},
			},
		},
	}
	return forAdmin
}

// Handler for slash command help
func Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: HelpMessageHandler(s, i),
		},
	})
}
