package models

type Spell struct {
	Name     string `json:"name,omitempty"`
	Level    int    `json:"level,omitempty"`
	MaxLevel int    `json:"maxLevel,omitempty"`
	Village  string `json:"village,omitempty"`
}
