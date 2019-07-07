package models

type Clan struct {
	Tag       string   `json:"tag,omitempty"`
	Name      string   `json:"name,omitempty"`
	ClanLevel int      `json:"clanLevel,omitempty"`
	BadgeUrl  []string `json:"badgeUrl,omitempty"`
}
