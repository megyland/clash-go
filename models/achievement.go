package models

type Achievement struct {
	Name           string `json:"name,omitempty"`
	Stars          int    `json:"stars,omitempty"`
	Value          int    `json:"value,omitempty"`
	Target         int    `json:"target,omitempty"`
	Info           string `json:"info,omitempty"`
	CompletionInfo string `json:"completionInfo,omitempty"`
	Village        string `json:"village,omitempty"`
}
