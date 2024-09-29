package discord

type DiscordWebhookMessage struct {
	Content         string                                   `json:"content"`
	Embeds          []DiscordWebhookMessageEmbed             `json:"embeds"`
	AllowedMentions DiscordWebhookMessageAllowedMentionRoles `json:"allowed_mentions"`
}

type DiscordWebhookMessageEmbed struct {
	Title  string                           `json:"title"`
	URL    string                           `json:"url"`
	Color  int                              `json:"color"`
	Author DiscordWebhookMessageEmbedAuthor `json:"author"`
	Image  DiscordWebhookMessageEmbedImage  `json:"image"`
}

type DiscordWebhookMessageEmbedAuthor struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type DiscordWebhookMessageEmbedImage struct {
	URL string `json:"url"`
}

type DiscordWebhookMessageAllowedMentionRoles struct {
	Roles []int `json:"roles"`
}

/*
 {
     "content": "<@&1287044269479297126>",
     "embeds": [
         {
             "title": "BetBoom Team выставила Saika на трансфер",
             "url": "https://www.cybersport.ru/tags/dota-2/betboom-team-vystavila-saika-na-transfer",
             "color": 39423,
             "author": {
                 "name": "cybersport.ru",
                 "url": "https://www.cybersport.ru",
                 "icon_url": "https://i.imgur.com/M70CrkO.png"
             },
             "image": {
                 "url": "https://virtus-img.cdnvideo.ru/images/material-card/plain/78/78b9c52b-7bfd-4a49-9a48-8dee30e23a1e.png"
             }
         }
     ],
     "allowed_mentions": {
         "roles": [
             1287044269479297126
         ]
     }
 }
*/
