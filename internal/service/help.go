package service

import "github.com/bwmarrin/discordgo"

const (
	helptext = "** Available commands: **\n- `!help`: Display available commands\n- `!translate`: Translate a sentence\n- `!reminder`: Set a reminder\n- `!play`: Play a game\n- `!weather`: Check the weather\n"
)

func HelpFunction(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, helptext)
}
