package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
	"time"
)

func Reminder(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!reminder") {

		timeStr := strings.TrimSpace(strings.TrimPrefix(m.Content, "!reminder"))

		duration, err := strconv.Atoi(timeStr)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Invalid time format. Please provide the reminder time in seconds.")
			return
		}

		go func() {
			time.Sleep(time.Duration(duration) * time.Second)
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Reminder: Time's up after %d seconds!", duration))
		}()

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Reminder set for %d seconds!", duration))
	}
}
