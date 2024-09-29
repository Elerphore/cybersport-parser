package discord

import (
	htmlparser "elerphore/cybersport-parser/internal/html_parser"
)

func PrepareWebhookMessage(news htmlparser.News) (discordWebhookMessage DiscordWebhookMessage) {
	discordWebhookMessage = DiscordWebhookMessage{
		Content: "<@&1287044269479297126>",
	}

	var allowedMentions = DiscordWebhookMessageAllowedMentionRoles{
		Roles: []int{1287044269479297126},
	}

	var author = DiscordWebhookMessageEmbedAuthor{
		Name:    "cybersport.ru",
		URL:     "https://www.cybersport.ru",
		IconURL: "https://i.imgur.com/M70CrkO.png",
	}

	var embed = DiscordWebhookMessageEmbed{
		Author: author,
		Color:  39423,
	}

	discordWebhookMessage.Embeds = []DiscordWebhookMessageEmbed{embed}
	discordWebhookMessage.AllowedMentions = allowedMentions

	return

}
