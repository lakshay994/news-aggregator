package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

var wpSitemap string = "https://www.washingtonpost.com/wp-stat/sitemaps/index.xml"

type Location struct {
	Loc string `xml:"loc"`
}

type SitemapIndex struct {
	Locations []*Location `xml:"sitemap"`
}

func (l *Location) String() string {
	return l.Loc
}

func main() {
	resp, _ := http.Get(wpSitemap)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	s := &SitemapIndex{}
	xml.Unmarshal(bytes, s)

	fmt.Println(s.Locations)
}
