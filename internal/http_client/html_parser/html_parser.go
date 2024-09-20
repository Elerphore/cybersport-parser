package htmlparser

import (
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func traverseChildren(n *html.Node) {

	if n.Data == "a" && n.Type == html.ElementNode && strings.Contains(n.Attr[1].Val, "link") {
		fmt.Println(n.Data, n.Attr[0].Val)
	}

	if n.Data == "h3" && n.Type == html.ElementNode {
		fmt.Println(n.Data, n.Attr)
	}

	if n.Type == html.TextNode && n.Parent.Data == "h3" {
		fmt.Println(n.Data, n.Attr)
	}

	if n.Data == "img" && n.Type == html.ElementNode {
		fmt.Println(n.Data, n.Attr[0].Val, "\n")
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseChildren(c)
	}
}

func Traverse(doc *html.Node) {
	if doc.Data == "article" {
		traverseChildren(doc)
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		Traverse(c)
	}

}

func ParseHTML(body io.ReadCloser) (parsedHTML *html.Node) {
	parsedHTML, errorHTML := html.Parse(body)

	if errorHTML != nil {
		log.Fatal(errorHTML)
	}

	return
}
