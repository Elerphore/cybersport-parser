package httpclient

import (
	"bytes"
	"elerphore/cybersport-parser/internal/discord"
	htmlparser "elerphore/cybersport-parser/internal/html_parser"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	cybersportUrl string = "https://www.cybersport.ru"
	httpClient           = http.Client{}
)

func DoGET() (resp *http.Response) {
	request, errRequest := http.NewRequest("GET", cybersportUrl, nil)

	if errRequest != nil {
		log.Fatal(errRequest)
	}

	resp, errDo := httpClient.Do(request)

	if errDo != nil {
		log.Fatal(errDo)
	}

	if resp.StatusCode != 200 {
		log.Fatal("StatusCode is:", resp.StatusCode)
	}

	return
}

func DoDiscordRequest(news htmlparser.News) {
	var discordWebhookMessage discord.DiscordWebhookMessage = discord.PrepareWebhookMessage(news)

	var image = discord.DiscordWebhookMessageEmbedImage{
		URL: news.ImageURL,
	}

	var embed = &discordWebhookMessage.Embeds[0]

	embed.Title = news.Title
	embed.URL = cybersportUrl + news.PostURL
	embed.Image = image

	message_body_bytes, err_bytes := json.Marshal(discordWebhookMessage)

	fmt.Println(string(message_body_bytes))

	if err_bytes != nil {
		log.Fatal(err_bytes)
	}

	request, errRequest := http.NewRequest("POST", os.Getenv("WEBHOOK"), bytes.NewBuffer(message_body_bytes))
	request.Header.Add("Content-Type", "application/json")

	if errRequest != nil {
		log.Fatal(errRequest)
	}

	resp, errDo := httpClient.Do(request)

	if errDo != nil {
		log.Fatal(errDo)
	}

	if resp.StatusCode > 299 {
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(bodyBytes))

		log.Fatal("StatusCode is:", resp.StatusCode)
	}
}
