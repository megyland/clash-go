package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/megyland/clash-go/models"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	Token     string

	httpClient *http.Client
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: "/v1" + path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(u.String())
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", "Bearer "+c.Token)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Client) GetPlayer(playerTag string) (models.Player, error) {
	playerTag = url.QueryEscape(playerTag)

	var player models.Player
	req, err := c.newRequest("GET", "/players/"+playerTag, nil)
	if err != nil {
		return player, err
	}
	_, err = c.do(req, &player)
	return player, err
}

func main() {
	apiUrl := &url.URL{Scheme: "https", Host: "api.clashofclans.com"}

	ClashGo := Client{BaseURL: apiUrl, Token: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImJhZDA1NTAxLWVhNWQtNDI3Yy1iM2MwLTBkNzUxMDQ4NjNhMiIsImlhdCI6MTU2MjQyMzEyMCwic3ViIjoiZGV2ZWxvcGVyLzkyMjFmZjM1LWRkNjctNTQ0ZC00ZTI5LTQ4ZThlOTgwZmZjNSIsInNjb3BlcyI6WyJjbGFzaCJdLCJsaW1pdHMiOlt7InRpZXIiOiJkZXZlbG9wZXIvc2lsdmVyIiwidHlwZSI6InRocm90dGxpbmcifSx7ImNpZHJzIjpbIjEwOS4xNC4xMTQuMTczIl0sInR5cGUiOiJjbGllbnQifV19.s0qxcdluiJr3MifW_E6jOEq6ZTY6-Nc-SYB_3lsQAmfO5c9GMkCQQv1ZpArUmIkzc0sbrDO2-KYt8VdsnU3dWg"}
	res, _ := ClashGo.GetPlayer("#8UCG8C0U")
	fmt.Println(res)
}
