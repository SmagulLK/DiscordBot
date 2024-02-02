package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	gtranslate "github.com/gilang-as/google-translate"
	"strings"
)

func Translate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!translate") {

		parts := strings.Split(m.Content, " | ")
		if len(parts) < 2 {
			s.ChannelMessageSend(m.ChannelID, "Invalid command format. Use `!translate | sentence`.")
			return
		}
		text := strings.TrimSpace(parts[1])

		translatedText, err := translateText(text)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error translating text.")
			fmt.Println("Error translating text:", err)
			return
		}

		s.ChannelMessageSend(m.ChannelID, translatedText)
	}
}

func translateText(text string) (string, error) {
	value := gtranslate.Translate{
		Text: text,

		To: "en",
	}
	translated, err := gtranslate.Translator(value)
	if err != nil {
		return "", err
	}

	return string(translated.Text), nil
}
