package bbcclient

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Url              string
	Categories       map[string]string
	PrettyCategories map[string]string
}

func NewClient() *Client {
	var categories = map[string]string{
		"top":           "",
		"world":         "world",
		"science":       "science_and_environment",
		"tech":          "technology",
		"uk":            "uk",
		"business":      "business",
		"politics":      "politics",
		"health":        "health",
		"education":     "education",
		"entertainment": "entertainment_and_arts",
	}
	var pretty = map[string]string{
		"top":           "Top News",
		"world":         "World",
		"science":       "Science and Environment",
		"tech":          "Technology",
		"uk":            "UK",
		"business":      "Business",
		"politics":      "Politics",
		"health":        "Health",
		"education":     "Education",
		"entertainment": "Entertainment and Arts",
	}
	var c = Client{
		Url:              "http://www.bbc.com/news/%s/rss.xml",
		Categories:       categories,
		PrettyCategories: pretty,
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
	url := fmt.Sprintf(c.Url, category)
	rep, err := c.RequestFeed(url)
	var feed Rss
	if err != nil {
		return feed, err
	}
	xml.Unmarshal(rep, &feed)
	return feed, nil
}

// func (c *Client) GetTop10(category string) (string, error) {
// 	if val, ok := c.Categories[category]; ok {
// 		rep, err := c.GetFeed(category)
// 		if (err != nil){
// 			return "", err
// 		}
// 		var urls [11]string
// 		urls[0] = "Top Stories from BBC " + c.PrettyCategories[category]
//     	items := world.Channel.Items
//     	for _, element := range items {
//     		fmt.Printf("%s %s\n", index, element)
//     	}
// 	} else {
// 		return "", fmt.Errorf("Invalid feed selection: %s\n", category)
// 	}
// }
