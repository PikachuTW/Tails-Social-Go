package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"tails-social-go/internal/event"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	if env := os.Getenv("ENV"); env != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	discordToken := os.Getenv("DISCORD_TOKEN")
	discord, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatalf("Error creating Discord client: %v", err)
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	discord.AddHandler(event.OnReady)
	discord.AddHandler(event.OnMessageCreate)

	if err := discord.Open(); err != nil {
		log.Fatalf("Error opening connection to Discord: %v", err)
	}
	log.Println("Connected to Discord")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	err = discord.Close()
	if err != nil {
		log.Fatalf("Error closing Discord: %v", err)
		return
	}
}
