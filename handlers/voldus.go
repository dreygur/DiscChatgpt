package handlers

import (
	"discchatgpt/api"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func VoldusMessageHandler(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	prompt := i.ApplicationCommandData().Options[0].StringValue()
	message, err := api.RequestToChatGPT(prompt)
	if err != nil {
		log.Println(err)
		return "Couldn't fetch ChatGPT"
	}
	var text string

	for _, v := range message.Choices {
		text = fmt.Sprintf("%s %s", text, v.Text)
	}
	return text
}

func Voldus(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			// Embeds: VoldusMessageHandler(s, i),
			Content: VoldusMessageHandler(s, i),
		},
	})
}
