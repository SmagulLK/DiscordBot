package main

import (
	"DiscordBot/internal/handlers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	godotenv.Load()
	botToken := os.Getenv("BOT_TOKEN")
	discordBot, err := discordgo.New("Bot " + botToken)
	if err != nil {
		slog.Error("error getting weather api key: %s", err.Error())
	}

	discordBot.Open()
	slog.Info("Bot is running")
	defer discordBot.Close()
	discordBot.AddHandler(handlers.MessageMiddleware)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
