// TODO: Test inStock
package inv

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	verboseLogging = false
)

func TestTestdataFolder(t *testing.T) {
	if _, err := os.Stat("testdata"); os.IsNotExist(err) {
		t.Errorf("FATAL: /testdata is not found.")
	}
}

func TestFind(t *testing.T) {
	tl := []Item{
		Item{"Pants", "item_cat_1", "item_slug_1"},
		Item{"Extendo Pants", "item_cat_2", "item_slug_2"},
		Item{"Hat", "item_cat_3", "item_slug_3"},
	}

	tb := []struct {
		in       string
		expected []Item
	}{
		{"PaNts", []Item{tl[0], tl[1]}},
		{"shorts", nil},
		{"ten", []Item{tl[1]}},
		{"a", tl},
	}

	for _, q := range tb {
		actual, _ := Find(q.in, tl)
		if compareItems(q.expected, actual) {
			t.Errorf("Could not find items for query: `%s` (slices not equal).\n"+
				"Expected:\n %v \n"+
				"Actual:\n %v \n", q.in, q.expected, actual)
		}
	}
}

func TestRetrieve(t *testing.T) {
	d := NewDownloader(bmPageFetcher)
	actual := Retrieve(d)
	expected := []Item{
		Item{"item_name_1", "item_cat_1", "item_slug_1"},
		Item{"item_name_2", "item_cat_2", "item_slug_2"},
		Item{"item_name_3", "item_cat_3", "item_slug_3"},
	}

	if compareItems(expected, actual) {
		t.Errorf("Could not retrieve items (slices not equal).\n"+
			"Expected:\n %v \n"+
			"Actual:\n %v \n", expected, actual)
	}
}

// URLs are generated from shop.html
func tdPageFetcher(url string) string {
	s := strings.Split(url, "/")

	slug := s[3]
	if len(s) > 4 {
		slug = s[4] + "-" + s[5]
	}

	ttslug := "testdata/" + slug + ".html"

	if verboseLogging {
		fmt.Printf("Actual URL -> %v \n", url)
		fmt.Printf("Test URL   -> %v \n\n", ttslug)
	}

	d, _ := ioutil.ReadFile(ttslug)
	return string(d)
}

func compareItems(a, b []Item) bool {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(append(a, b...))

	c := 0
	for _, x := range buf.Bytes() {
		c ^= int(x)
	}
	return c == 0
}

func bmPageFetcher(url string) string {
	s := strings.Split(url, "/")

	slug := s[3]
	if len(s) > 4 {
		slug = s[4] + "-" + s[5]
	}

	ttslug := "testdata/benchmarks/" + slug + ".html"

	if verboseLogging {
		fmt.Printf("Actual URL -> %v \n", url)
		fmt.Printf("Test URL   -> %v \n\n", ttslug)
	}

	d, _ := ioutil.ReadFile(ttslug)
	return string(d)
}

func BenchmarkRetrieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDownloader(bmPageFetcher)
		Retrieve(d)
	}
}

func BenchmarkParseLinkNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		raw := "testdata/benchmarks/shop.html"
		parseLinkNodes(raw)
	}
}

func BenchmarkFetchItemInfo(b *testing.B) {
	d := NewDownloader(bmPageFetcher)
	raw := d.download("")
	items := parseLinkNodes(raw)
	mockItems := items
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fetchItemInfo(&mockItems, d)
		mockItems = items
	}
}
