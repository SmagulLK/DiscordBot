package handlers

import (
	"DiscordBot/internal/service"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Router(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	const (
		HelpCommand      = "!help"
		PollCommand      = "!poll"
		WeatherCommand   = "!weather"
		ReminderCommand  = "!reminder"
		PlayCommand      = "!play"
		TranslateCommand = "!translate"
	)

	commands := map[string]func(){
		HelpCommand:      func() { service.HelpFunction(s, m) },
		PollCommand:      func() { service.PollService(s, m) },
		WeatherCommand:   func() { service.WeatherCheck(s, m) },
		ReminderCommand:  func() { service.Reminder(s, m) },
		PlayCommand:      func() { service.Game(s, m) },
		TranslateCommand: func() { service.Translate(s, m) },
	}

	for command, _ := range commands {
		if strings.Contains(m.Content, command) {
			commands[command]()
		}
	}
}
