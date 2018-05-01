package link // import "github.com/Athulus/gophercise/link"

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

/*
Link represents an <a href=...></a> tag
*/
type Link struct {
	Href string
	Text string
}

/*
GetLinks takes in a reader with html data and retruns
a slice of all of the links in that reader
*/
func GetLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	fmt.Println("test")
	fmt.Println(doc)
	linkNodes := getLinkNodes(doc.FirstChild)
	links := make([]Link, len(linkNodes))
	for i, node := range linkNodes {
		links[i] = buildLink(node)
	}
	return links, nil
}

func getLinkNodes(n *html.Node) []*html.Node {
	var linkNodes []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	fmt.Println(n)
	//iterte through node depth first
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Println(c)
		linkNodes = append(linkNodes, getLinkNodes(c)...)
	}
	return linkNodes

}

func buildLink(n *html.Node) Link {
	var link Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.Text = text(n)
	return link
}

func text(n *html.Node) string {
	var linkText string
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		linkText += text(c) + " "
	}
	return strings.Join(strings.Fields(linkText), " ")
}
