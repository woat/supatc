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
	name     string
	category string
	slug     string
}

type PageFetcher func(url string) string

type Downloader struct {
	fetchPage PageFetcher
}

func (d *Downloader) download() string {
	return d.fetchPage(url)
}

func NewDownloader(fp PageFetcher) *Downloader {
	return &Downloader{fetchPage: fp}
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

// Once the slugs are retrieved then the item names can be found through
// parsing the <title> of each page.
func fetchItemNames(l []Item) {
	var wg sync.WaitGroup
	wg.Add(len(l))

	for i := 0; i < len(l); i++ {
		go func(i int, l []Item) {
			defer wg.Done()
			raw := fetchHTML(url + "/" + l[i].category + "/" + l[i].slug)
			l[i].name = parseItemName(raw)
		}(i, l)
	}

	wg.Wait()
}

func parseItemName(raw string) string {
	doc, err := html.Parse(strings.NewReader(raw))
	if err != nil {
		panic(err)
	}

	return rNameSearch(doc)
}

// Connecting the item slug to a name will allow for features such as
// add-to-cart by name become very easy.
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

func StdDl() *Downloader {
	return NewDownloader(fetchHTML)
}

func Retrieve(d *Downloader) []Item {
	raw := d.download()
	items := parseLinkNodes(raw)
	fetchItemNames(items)
	return items
}

func Find(sep string, l []Item) ([]Item, bool) {
	found := make([]Item, 0)

	for _, item := range l {
		in := strings.ToLower(item.name)
		sep = strings.ToLower(sep)
		if strings.Contains(in, sep) {
			fmt.Println(item)
		}
	}

	if len(found) != 0 {
		return found, true
	}

	return nil, false
}
