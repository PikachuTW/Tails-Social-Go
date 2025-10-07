package event

import (
	"fmt"
	"log"
	"tails-social-go/internal/scraper"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(client *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot || message.WebhookID != "" {
		return
	}

	facebookScraper := scraper.NewFacebookScraper()
	youtubeScraper := scraper.NewYoutubeScraper()

	for _, s := range [2]scraper.Scraper{facebookScraper, youtubeScraper} {
		match := s.Match(message.Content)
		if match == "" {
			continue
		}

		data, err := s.FetchData(match)
		if err != nil {
			log.Println(err)
			continue
		}
		if data == nil {
			log.Printf("No data found for %s", match)
			continue
		}

		log.Println(data.Image)

		_, err = client.ChannelMessageSendEmbedReply(message.ChannelID, &discordgo.MessageEmbed{
			URL:         data.URL,
			Title:       data.Title,
			Description: data.Description,
			Color:       s.EmbedColor(),
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("%s | Sent by %s", s.SourceName(), message.Author.GlobalName),
			},
			Image: &discordgo.MessageEmbedImage{
				URL: data.Image,
			},
		}, message.Reference())

		if err != nil {
			log.Println(err)
			return
		}
	}
}
