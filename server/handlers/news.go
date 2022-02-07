package handlers

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

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

func init() {
	s := &SitemapIndex{}
	n := &News{}
	xml.Unmarshal([]byte(wpSitemap), &s)

	for _, Location := range s.Locations {
		Location = strings.TrimSpace(Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)

		for i := range n.Keywords {
			newsMap[n.Titles[i]] = NewsMap{Keyword: n.Keywords[i], Location: n.Locations[i]}
		}
	}
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "News agrregator", News: newsMap}
	directory, _ := os.Getwd()
	t, _ := template.ParseFiles(path.Join(directory, "server", "templates", "news.html"))
	t.Execute(w, p)
}
