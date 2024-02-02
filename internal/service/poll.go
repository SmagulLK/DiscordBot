package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func PollService(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!poll" {
		s.ChannelMessageSend(m.ChannelID, "Please provide a question and options for the poll in the format: `!poll | question | option1 | option2 | ...`")
		return
	}

	if strings.HasPrefix(m.Content, "!poll |") {
		args := strings.Split(m.Content, "|")

		for i := range args {
			args[i] = strings.TrimSpace(args[i])
		}

		if len(args) < 3 {
			s.ChannelMessageSend(m.ChannelID, "Invalid format. Please use `!poll | question | option1 | option2 | ...`")
			return
		}

		question := args[1]
		options := args[2:]

		pollMessage := "📊 **Poll:** " + question + "\n\n"

		for i, option := range options {
			pollMessage += fmt.Sprintf(":%d: %s\n", i+1, option)
		}

		message, err := s.ChannelMessageSend(m.ChannelID, pollMessage)
		if err != nil {
			fmt.Println("Error sending message: ", err)
			return
		}

		for i := range options {
			err := s.MessageReactionAdd(m.ChannelID, message.ID, fmt.Sprintf("%d⃣", i+1))
			if err != nil {
				fmt.Println("Error adding reaction: ", err)
				return
			}
		}
	}
}
