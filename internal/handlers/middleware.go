package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log/slog"
	"strings"
)

func MessageMiddleware(s *discordgo.Session, m *discordgo.MessageCreate) {
	keywords := []string{"!help", "!poll", "!weather", "!play", "!translate", "!reminder"}

	for _, keyword := range keywords {
		if strings.HasPrefix(m.Content, keyword) {

			slog.Info("command:", keyword)
		}
	}

	Router(s, m)
}
