package service

import (
	"DiscordBot/internal/models"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

const (
	lenError       = "Is empty or too long"
	openWeatherURL = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s"
)

func WeatherCheck(s *discordgo.Session, m *discordgo.MessageCreate) {

	if err := godotenv.Load(); err != nil {
		slog.Error("error getting weather api key from environment file: %s", err.Error())
	}
	openWeatherAPIKey := os.Getenv("WEATHER_API_KEY")
	if openWeatherAPIKey == "" {
		slog.Error("error getting weather api key")
		return
	}
	text := strings.Split(m.Content, "|")
	if len(text) != 2 {
		s.ChannelMessageSend(m.ChannelID, lenError)
		return
	}

	location := strings.TrimSpace(text[1])

	url := fmt.Sprintf(openWeatherURL, location, openWeatherAPIKey)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != http.StatusOK {
		s.ChannelMessageSend(m.ChannelID, "Error fetching weather data or city not found.")
		return
	}
	defer resp.Body.Close()

	var weatherResp models.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error decoding weather data")
		return
	}

	temperatureCelsius := weatherResp.Main.Temperature - 273.15

	message := fmt.Sprintf("Current temperature in %s: %.2f°C", location, temperatureCelsius)
	s.ChannelMessageSend(m.ChannelID, message)
}
