// GENERATES MOCK DATA (FROM LIVE) FOR BENCHMARKS
package main

import (
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

func main() {
	d := StdDl()
	shop := d.download("")
	ioutil.WriteFile("shop.html", []byte(shop), 0644)
	items := parseLinkNodes(shop)
	fetchItemInfo(&items, d)
}

const (
	url = "https://www.supremenewyork.com/shop"
)

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

func rLinkSearch(n *html.Node, l *[]Item) {
	if hasTag("a", n) {
		for _, attr := range n.Attr {
			if hasItemLink(attr) {
				s := strings.Split(attr.Val, "/")
				*l = append(*l, Item{Category: s[2], Slug: s[3]})
			}
		}
	}

	for cn := n.FirstChild; cn != nil; cn = cn.NextSibling {
		rLinkSearch(cn, l)
	}
}

func hasTag(t string, n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == t
}

func hasItemLink(a html.Attribute) bool {
	r, _ := regexp.Compile("/shop/.*?/.*")
	return a.Key == "href" && r.MatchString(a.Val)
}

func fetchItemInfo(l *[]Item, d *Downloader) {
	var wg sync.WaitGroup
	wg.Add(len(*l))

	for i := 0; i < len(*l); i++ {
		go func(i int, l *[]Item) {
			defer wg.Done()
			raw := d.download("/" + (*l)[i].Category + "/" + (*l)[i].Slug)
			ioutil.WriteFile((*l)[i].Category+"-"+(*l)[i].Slug+".html", []byte(raw), 0644)
		}(i, l)
	}

	wg.Wait()
}

type Item struct {
	Category string
	Slug     string
}

type Downloader struct {
	fetchPage PageFetcher
}

type PageFetcher func(url string) string

func (d *Downloader) download(slug string) string {
	return d.fetchPage(url + slug)
}

func StdDl() *Downloader {
	return NewDownloader(fetchHTML)
}
func NewDownloader(pf PageFetcher) *Downloader {
	return &Downloader{fetchPage: pf}
}
