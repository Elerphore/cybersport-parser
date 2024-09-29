package cybersport

import (
	htmlparser "elerphore/cybersport-parser/internal/html_parser"
	httpclient "elerphore/cybersport-parser/internal/http_client"
	"elerphore/cybersport-parser/internal/sqlite"
	"time"
)

func GetNews() {
	var resp = httpclient.DoGET()

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var proccessedNewsList = sqlite.GetNews()

	newArticlesExist := false

	var doc = htmlparser.ParseHTML(resp.Body)

	var lastCheckedLink string = ""

	if len(proccessedNewsList) > 0 {
		lastCheckedLink = proccessedNewsList[0].Link
		newArticlesExist = htmlparser.CheckForNewArticles(doc, lastCheckedLink)
	} else {
		newArticlesExist = true
	}

	if !newArticlesExist {
		return
	}

	var newsList = []htmlparser.News{}

	htmlparser.Traverse(doc, &newsList, lastCheckedLink)

	for i, j := 0, len(newsList)-1; i < j; i, j = i+1, j-1 {
		newsList[i], newsList[j] = newsList[j], newsList[i]
	}

	for _, item := range newsList {
		httpclient.DoDiscordRequest(item)

		sqlite.InsertNews(sqlite.News{Link: item.PostURL, NewsSourceId: 1})

		time.Sleep(5 * time.Second)
	}

	return
}
