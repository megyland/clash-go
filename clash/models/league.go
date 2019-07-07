package models

type League struct {
	Id       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	IconUrls []string `json:"iconUrls,omitempty"`
}
