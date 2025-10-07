package event

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnReady(client *discordgo.Session, ready *discordgo.Ready) {
	log.Printf("[Ready] %s\nRunning for %d", ready.User.Username, len(ready.Guilds))
	err := client.UpdateGameStatus(0, "Tails' Social Media")
	if err != nil {
		log.Printf("Failed to update status: %v", err)
	}
}
