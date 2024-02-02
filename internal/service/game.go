package service

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Game(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!play") {
		go playGame(s, m.ChannelID, m.Author.ID)
	}
}
func playGame(s *discordgo.Session, channelID string, userID string) {
	rand.Seed(time.Now().UnixNano())
	botNumber := rand.Intn(100) + 1

	s.ChannelMessageSend(channelID, "Guess a number between 1 and 100!")

	for {
		msg := <-waitForMessages(s, channelID, userID)

		userGuess, err := strconv.Atoi(msg.Content)
		if err != nil {
			s.ChannelMessageSend(channelID, "Please enter a valid number.")
			continue
		}

		if userGuess < botNumber {
			s.ChannelMessageSend(channelID, "Your number is smaller!")
		} else if userGuess > botNumber {
			s.ChannelMessageSend(channelID, "Your number is bigger!")
		} else {
			s.ChannelMessageSend(channelID, "Congratulations! You guessed the number.")
			return
		}
	}
}

func waitForMessages(s *discordgo.Session, channelID string, userID string) <-chan *discordgo.MessageCreate {
	ch := make(chan *discordgo.MessageCreate)

	messageHandler := func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if m.Author.ID == userID && m.ChannelID == channelID {
			ch <- m
		}
	}

	s.AddHandler(messageHandler)

	return ch
}
