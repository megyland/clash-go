package models

type Location struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IsCountry bool   `json:"isCountry,omitempty"`
}
