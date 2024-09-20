package htmlparser

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

func Traverse(doc *html.Node) {
	var traverse func(*html.Node)

	traverse = func(n *html.Node) {

		if n.Data == "h3" && n.Type == html.ElementNode {
			fmt.Println(n.Data, n.Attr)
		}

		if n.Type == html.TextNode && n.Parent.Data == "h3" {
			fmt.Println(n.Data, n.Attr, "\n")
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)
}

func ParseHTML(body io.ReadCloser) (parsedHTML *html.Node) {
	parsedHTML, errorHTML := html.Parse(body)

	if errorHTML != nil {
		log.Fatal(errorHTML)
	}

	return
}
