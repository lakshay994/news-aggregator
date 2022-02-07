package handlers

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var wpSitemap = `
<sitemapindex>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/politics.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/opinions.xml
</loc>
</sitemap>
</sitemapindex>
`
var newsMap = make(map[string]NewsMap)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func fetchNews(c chan *News, location string) {
	defer wg.Done()

	n := &News{}
	location = strings.TrimSpace(location)
	resp, _ := http.Get(location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &n)

	c <- n
}

func init() {
	s := &SitemapIndex{}
	queue := make(chan *News, 50)
	xml.Unmarshal([]byte(wpSitemap), &s)

	for _, Location := range s.Locations {
		wg.Add(1)
		fetchNews(queue, Location)
	}
	wg.Wait()
	close(queue)

	for news := range queue {
		for i := range news.Keywords {
			newsMap[news.Titles[i]] = NewsMap{Keyword: news.Keywords[i], Location: news.Locations[i]}
		}
	}
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "News agrregator", News: newsMap}
	directory, _ := os.Getwd()
	t, _ := template.ParseFiles(path.Join(directory, "server", "templates", "news.html"))
	t.Execute(w, p)
}
