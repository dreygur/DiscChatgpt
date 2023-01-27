package handlers

import (
	"discchatgpt/api"
	"log"

	"github.com/bwmarrin/discordgo"
)

func VoldusMessageHandler(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	prompt := i.ApplicationCommandData().Options[0].StringValue()
	message, err := api.RequestToChatGPT(prompt, "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")
	if err != nil {
		log.Println(err)
		return "Couldn't fetch ChatGPT"
	}
	return message.Choices[0].Text
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
