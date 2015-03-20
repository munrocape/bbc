package bbcclient

import "encoding/xml"

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName       xml.Name `xml:"channel"`
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	Language      string   `xml:"language"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Copyright     string   `xml:"copyright"`
	Image         Image    `xml:"image"`
	Ttl           string   `xml:"ttl"`
	Atom          Atom     `xml:"atom"`
	Items         []Item   `xml:"item"`
}

type Image struct {
	XMLName xml.Name `xml:"image"`
	Url     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Width   string   `xml:"width"`
	Height  string   `xml:"height"`
}

type Atom struct {
	XMLName xml.Name `xml:"atom"`
	Href    string   `xml:"href,attr"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	Guid        string   `xml:"guid"`
	PubDate     string   `xml:"pubDate"`
	Media       []Media  `xml:"media:thumbnail"`
}

type Media struct {
	XMLName xml.Name `xml:"media:thumbnail"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Url     string   `xml:"url,attr"`
}
