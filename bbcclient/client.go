package bbcclient

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	NewsUrl          string
	SportsUrl        string
	NewsCategories   map[string]Category
	SportsCategories map[string]Category
}

type Category struct {
	Uri    string
	Pretty string
}

func NewClient() *Client {
	var newsCategories = map[string]Category{
		"top":           Category{Uri: "", Pretty: "Top Stories"},
		"world":         Category{Uri: "/world", Pretty: "World"},
		"science":       Category{Uri: "/science_and_environment", Pretty: "Science and Environment"},
		"tech":          Category{Uri: "/technology", Pretty: "Technology"},
		"uk":            Category{Uri: "/uk", Pretty: "UK"},
		"business":      Category{Uri: "/business", Pretty: "Business"},
		"politics":      Category{Uri: "/politics", Pretty: "Politics"},
		"health":        Category{Uri: "/health", Pretty: "Health"},
		"education":     Category{Uri: "/education", Pretty: "Education"},
		"entertainment": Category{Uri: "/entertainment_and_arts", Pretty: "Entertainment and Arts"},
	}
	var sportsCategories = map[string]Category{
		"sports":       Category{Uri: "", Pretty: "Sports"},
		"football":     Category{Uri: "/football", Pretty: "Football"},
		"cricket":      Category{Uri: "/cricket", Pretty: "Cricket"},
		"rugby":        Category{Uri: "/rugby-union", Pretty: "Rugby Union"},
		"rugby_league": Category{Uri: "/rugby-league", Pretty: "Rugby League"},
		"tennis":       Category{Uri: "/tennis", Pretty: "Tennis"},
		"golf":         Category{Uri: "/Golf", Pretty: "Golf"},
		"snooker":      Category{Uri: "/snooker", Pretty: "Snooker"},
	}
	var c = Client{
		NewsUrl:          "http://www.bbc.com/news%s/rss.xml",
		SportsUrl:        "http://feeds.bbci.co.uk/sport/0%s/rss.xml?edition=uk",
		NewsCategories:   newsCategories,
		SportsCategories: sportsCategories,
	}
	return &c
}

func (c *Client) RequestFeed(url string) ([]byte, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (c *Client) GetFeed(category string) (Rss, error) {
	var url string
	var feed Rss
	if val, ok := c.NewsCategories[category]; ok {
		url = fmt.Sprintf(c.NewsUrl, val.Uri)
	} else if val, ok := c.SportsCategories[category]; ok {
		url = fmt.Sprintf(c.SportsUrl, val.Uri)
	} else {
		return feed, fmt.Errorf("Invalid feed selection: %s\n", category)
	}

	rep, err := c.RequestFeed(url)
	if err != nil {
		return feed, err
	}
	xml.Unmarshal(rep, &feed)
	return feed, nil
}

func (c *Client) GetPretty(category string) string {
	if val, ok := c.NewsCategories[category]; ok {
		return val.Pretty
	} else if val, ok := c.SportsCategories[category]; ok {
		return val.Pretty
	} else {
		return ""
	}
}

func (c *Client) GetUrl(category string) string {
	if _, ok := c.NewsCategories[category]; ok {
		uri := c.NewsCategories[category].Uri
		return fmt.Sprintf("http://www.bbc.com/news%s", uri)
	} else if _, ok := c.SportsCategories[category]; ok {
		uri := c.SportsCategories[category].Uri
		return fmt.Sprintf("http://www.bbc.com/sport/0%s", uri)
	} else {
		return ""
	}
}
