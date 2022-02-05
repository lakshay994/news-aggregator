package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
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

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>titile"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	s := &SitemapIndex{}
	n := &News{}
	xml.Unmarshal([]byte(wpSitemap), s)

	for _, Location := range s.Locations {
		Location = strings.TrimSpace(Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
	}
}
