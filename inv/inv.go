package inv

import (
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

// Downloader will hold a PageFetcher for use in dependency injection.
// See fetchHTML.
type Downloader struct {
	fetchPage PageFetcher
}

type PageFetcher func(url string) string

func (d *Downloader) download(slug string) string {
	return d.fetchPage(url + slug)
}

// See inv_test.go for mock example.
func NewDownloader(fp PageFetcher) *Downloader {
	return &Downloader{fetchPage: fp}
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
	if isTag("a", n) {
		for _, a := range n.Attr {
			if isItemLink(a) {
				s := strings.Split(a.Val, "/")
				*l = append(*l, Item{category: s[2], slug: s[3]})
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		rLinkSearch(c, l)
	}
}

func isTag(t string, n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == t
}

func isItemLink(a html.Attribute) bool {
	r, _ := regexp.Compile("/shop/.*?/.*")
	return a.Key == "href" && r.MatchString(a.Val)
}

// Once the slugs are retrieved then the item names can be found through
// parsing the <title> of each page.
func fetchItemNames(l []Item, d *Downloader) {
	var wg sync.WaitGroup
	wg.Add(len(l))

	for i := 0; i < len(l); i++ {
		go func(i int, l []Item) {
			defer wg.Done()
			raw := d.download("/" + l[i].category + "/" + l[i].slug)
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
	if isTag("title", n) {
		return n.FirstChild.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s := rNameSearch(c)
		// TODO: This might infinite loop.
		if s != "" {
			return s
		}
	}

	return ""
}

// Standard Downloader available for usuage in cross-lib.
// Might be an anti-pattern, but TDD GOD.
func StdDl() *Downloader {
	return NewDownloader(fetchHTML)
}

// Retrieves all inventory information available using a Downloader.
// See Downloader, PageFetcher, fetchHTML.
func Retrieve(d *Downloader) []Item {
	raw := d.download("")
	items := parseLinkNodes(raw)
	fetchItemNames(items, d)
	return items
}

// Takes a search query and item list returns a new item list of the matches.
// Use v, ok convention.
func Find(q string, l []Item) ([]Item, bool) {
	found := make([]Item, 0)

	for _, item := range l {
		r, _ := regexp.Compile("(?i)" + q)
		if r.MatchString(item.name) {
			found = append(found, item)
		}
	}

	if len(found) != 0 {
		return found, true
	}

	return found, false
}
