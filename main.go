package main

import (
	httpclient "elerphore/cybersport-parser/internal/http_client"
	htmlparser "elerphore/cybersport-parser/internal/http_client/html_parser"
	"fmt"
)

func main() {
	var resp = httpclient.DoGET()

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var doc = htmlparser.ParseHTML(resp.Body)

	var newsList = []htmlparser.News{}
	htmlparser.Traverse(doc, &newsList)

	for _, item := range newsList {
		fmt.Println(item, "\n\n")
	}
}
