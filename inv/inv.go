// Package inv is used to find, parse, and locate all items stocked. The amount
// of recursion found here does make me a bit sad.
package inv

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

const (
	url = "https://www.supremenewyork.com/shop"
)

type Item struct {
	Name     string
	Category string
	Slug     string
}

// Downloader will hold a PageFetcher for use in dependency injection.
// See fetchHTML.
type Downloader struct {
	fetchPage PageFetcher
}

type PageFetcher func(url string) string

func (d *Downloader) download(slug string) string {
	return d.fetchPage(url + slug)
}

// See inv_test.go for example.
func NewDownloader(pf PageFetcher) *Downloader {
	return &Downloader{fetchPage: pf}
}

// Default PageFetcher for this package.
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

// The anchor links contains all the product slugs available for current drop.
// Searching through each href and matching it toward the shop slug will
// return a list of all the products slug ready for further parsing.
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
	fmt.Println(n.Data)
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

// The usage of pointers in this function makes me feel sad.
func fetchItemInfo(l *[]Item, d *Downloader) {
	var is []Item
	mx := &sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(len(*l))

	for i := 0; i < len(*l); i++ {
		go func(i int, l *[]Item) {
			defer wg.Done()
			raw := d.download("/" + (*l)[i].Category + "/" + (*l)[i].Slug)
			if !outOfStock(raw) {
				(*l)[i].Name = parseItemName(raw)
				mx.Lock()
				is = append(is, (*l)[i])
				mx.Unlock()
			}
		}(i, l)
	}

	wg.Wait()
	*l = is
}

func outOfStock(raw string) bool {
	doc, err := html.Parse(strings.NewReader(raw))
	if err != nil {
		panic(err)
	}

	return rStockSearch(doc)
}

func rStockSearch(n *html.Node) bool {
	if hasTag("b", n) {
		for _, attr := range n.Attr {
			if attr.Key == "class" {
				return true
			}
		}
	}

	for cn := n.FirstChild; cn != nil; cn = cn.NextSibling {
		res := rStockSearch(cn)
		if res != false {
			return res
		}
	}

	return false
}

func parseItemName(raw string) string {
	doc, err := html.Parse(strings.NewReader(raw))
	if err != nil {
		panic(err)
	}

	return rNameSearch(doc)
}

func rNameSearch(n *html.Node) string {
	if hasTag("title", n) {
		return n.FirstChild.Data
	}

	for cn := n.FirstChild; cn != nil; cn = cn.NextSibling {
		title := rNameSearch(cn)
		if title != "" {
			return title
		}
	}

	return ""
}

// Standard Downloader available for usuage in cross pkg.
func StdDl() *Downloader {
	return NewDownloader(fetchHTML)
}

// Retrieves all inventory information available using a Downloader. I might
// have only written this just to test it. Open/Closed principle?
// See Downloader, PageFetcher, fetchHTML.
func Retrieve(d *Downloader) []Item {
	raw := d.download("")
	items := parseLinkNodes(raw)
	fetchItemInfo(&items, d)
	return items
}

// Takes a search query and item list returns a new item list of the matches.
// A feature for a future feature.
func Find(q string, l []Item) ([]Item, bool) {
	var found []Item

	for _, item := range l {
		r, _ := regexp.Compile("(?i)" + q)
		if r.MatchString(item.Name) {
			found = append(found, item)
		}
	}

	if len(found) != 0 {
		return found, true
	}

	return nil, false
}
