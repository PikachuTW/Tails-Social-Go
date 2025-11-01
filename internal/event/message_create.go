package event

import (
	"fmt"
	"log"
	"sync"
	"tails-social-go/internal/scraper"

	"github.com/bwmarrin/discordgo"
)

var scrapers = []scraper.Scraper{
	scraper.NewFacebookScraper(),
	scraper.NewYoutubeScraper(),
	// scraper.NewThreadsScraper(),
}

func OnMessageCreate(client *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot || message.WebhookID != "" {
		return
	}

	var wg sync.WaitGroup

	for _, s := range scrapers {
		wg.Add(1)

		go func(s scraper.Scraper) {
			defer wg.Done()

			match := s.Match(message.Content)
			if match == "" {
				return
			}

			data, err := s.FetchData(match)
			if err != nil {
				log.Println(err)
				return
			}

			if data == nil {
				log.Printf("No data found for %s", match)
				return
			}

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
		}(s)
	}

	wg.Wait()
}
