package clash_go

import (
	"encoding/json"
	"github.com/megyland/clash-go/models"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

func (c *Client) GetPlayer(playerTag string) ([]models.Player, error) {
	rel := &url.URL{Path: "/players/" + playerTag}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var users []models.Player
	err = json.NewDecoder(resp.Body).Decode(&users)
	return users, err
}
