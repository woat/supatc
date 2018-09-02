package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Item struct {
	name     string
	category string
	slug     string
}

const (
	url = "https://www.supremenewyork.com/shop"
)

func main() {
	raw := fetchHTML(url)
	items := parseLinkNodes(raw)
	fetchItemNames(items)
	findItem("black", items)
}

func fetchHTML(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(raw)
}

func parseLinkNodes(raw string) []Item {
	doc, err := html.Parse(strings.NewReader(raw))
	if err != nil {
		panic(err)
	}
	l := []Item{}
	rLinkSearch(doc, &l)
	return l
}

// Used in parseLinkNodes to find anchor tags that have a shop href.
// We must match for /shop/[item.category]/[item.slug].
func rLinkSearch(n *html.Node, l *[]Item) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			r, _ := regexp.Compile("/shop/.*?/.*")
			if a.Key == "href" && r.MatchString(a.Val) {
				s := strings.Split(a.Val, "/")
				*l = append(*l, Item{category: s[2], slug: s[3]})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		rLinkSearch(c, l)
	}
}

func fetchItemNames(l []Item) {
	for i := 0; i < len(l); i++ {
		raw := fetchHTML(url + "/" + l[i].category + "/" + l[i].slug)
		l[i].name = parseItemName(raw)
	}
}

func parseItemName(raw string) string {
	doc, err := html.Parse(strings.NewReader(raw))
	if err != nil {
		panic(err)
	}
	return rNameSearch(doc)
}

func rNameSearch(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s := rNameSearch(c)
		if s != "" {
			return s
		}
	}
	return ""
}

func findItem(sep string, items []Item) {
	for _, item := range items {
		in := strings.ToLower(item.name)
		sep = strings.ToLower(sep)
		if strings.Contains(in, sep) {
			fmt.Println(item)
		}
	}
}
