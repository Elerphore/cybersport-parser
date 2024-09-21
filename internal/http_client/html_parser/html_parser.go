package htmlparser

import (
	"io"
	"log"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

func traverseChildren(n *html.Node, news *News) {
	if n.Data == "a" && n.Type == html.ElementNode && strings.Contains(n.Attr[1].Val, "link") {
		news.PostURL = n.Attr[0].Val
	}

	if n.Type == html.TextNode && n.Parent.Data == "h3" {
		news.Title = n.Data
	}

	if n.Data == "img" && n.Type == html.ElementNode {
		news.ImageURL = n.Attr[0].Val
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseChildren(c, news)
	}
}

func Traverse(doc *html.Node, newsList *[]News) {
	if doc.Data == "article" {
		var news = News{}

		traverseChildren(doc, &news)

		*newsList = slices.Insert(*newsList, len(*newsList), news)
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		Traverse(c, newsList)
	}
}

func ParseHTML(body io.ReadCloser) (parsedHTML *html.Node) {
	parsedHTML, errorHTML := html.Parse(body)

	if errorHTML != nil {
		log.Fatal(errorHTML)
	}

	return
}

type News struct {
	Title    string
	PostURL  string
	ImageURL string
}
