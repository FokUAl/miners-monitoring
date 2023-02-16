package pkg

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

func GetAttribute(n *html.Node, key string) (string, bool) {

	for _, attr := range n.Attr {

		if attr.Key == key {
			return attr.Val, true
		}
	}

	return "", false
}

func RenderNode(n *html.Node) string {

	var buf bytes.Buffer
	w := io.Writer(&buf)

	err := html.Render(w, n)

	if err != nil {
		return ""
	}

	return buf.String()
}

func CheckId(n *html.Node, id string) bool {

	if n.Type == html.ElementNode {

		s, ok := GetAttribute(n, "id")

		if ok && s == id {
			return true
		}
	}

	return false
}

func Traverse(n *html.Node, id string) *html.Node {

	if CheckId(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		res := Traverse(c, id)

		if res != nil {
			return res
		}
	}

	return nil
}

func GetElementById(n *html.Node, id string) *html.Node {

	return Traverse(n, id)
}
