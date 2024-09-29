package htmlparser

import (
	"io"
	"log"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

func Traverse(doc *html.Node, newsList *[]News, lastCheckedLink string) {
	var stopTraverse = false

	var traverseChildren func(n *html.Node, news *News)

	traverseChildren = func(n *html.Node, news *News) {
		if n.Data == "a" && n.Type == html.ElementNode && strings.Contains(n.Attr[1].Val, "link") {

			if n.Attr[0].Val == lastCheckedLink {
				stopTraverse = true
				return
			}

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

			if stopTraverse {
				break
			}
		}
	}

	var traverseArticles func(doc *html.Node, newsList *[]News)

	traverseArticles = func(doc *html.Node, newsList *[]News) {
		if doc.Data == "article" {
			var news = News{}

			traverseChildren(doc, &news)

			if len(news.PostURL) > 0 {
				*newsList = slices.Insert(*newsList, len(*newsList), news)
			}
		}

		for c := doc.FirstChild; c != nil; c = c.NextSibling {
			traverseArticles(c, newsList)

			if stopTraverse {
				break
			}

		}
	}

	traverseArticles(doc, newsList)

}

func ParseHTML(body io.ReadCloser) (parsedHTML *html.Node) {
	parsedHTML, errorHTML := html.Parse(body)

	if errorHTML != nil {
		log.Fatal(errorHTML)
	}

	return
}

func CheckForNewArticles(doc *html.Node, lastCheckedLink string) (isNewArticleExist bool) {
	isNewArticleExist = false

	var traverseChildren func(n *html.Node) (isNewArticleExist bool)

	traverseChildren = func(n *html.Node) bool {
		var isNewArticleExist bool
		if n.Data == "a" && n.Type == html.ElementNode && strings.Contains(n.Attr[1].Val, "link") {
			isNewArticleExist = n.Attr[0].Val != lastCheckedLink
			return isNewArticleExist
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			isNewArticleExist = traverseChildren(c)

			if isNewArticleExist {
				break
			}

		}

		return isNewArticleExist
	}

	var traverse func(doc *html.Node) (isNewLinExists bool)

	traverse = func(doc *html.Node) (isNewLinExists bool) {
		if doc.Data == "article" {
			traverseChildren(doc)
			return
		}

		for c := doc.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
		return
	}

	var findFirstArticle func(doc *html.Node) (article *html.Node)

	findFirstArticle = func(doc *html.Node) *html.Node {
		var article *html.Node

		if doc.Data == "article" {
			article = doc
			return article
		}

		for c := doc.FirstChild; c != nil; c = c.NextSibling {
			if article == nil {
				article = findFirstArticle(c)
			}
		}

		return article
	}

	var firstArticle = findFirstArticle(doc)
	isNewArticleExist = traverseChildren(firstArticle)

	return
}
