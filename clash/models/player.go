package models

type Player struct {
	Tag                  string `json:"tag,omitempty"`
	Name                 string `json:"name,omitempty"`
	TownHallLevel        int    `json:"townHallLevel,omitempty"`
	ExpLevel             int    `json:"expLevel,omitempty"`
	Trophies             int    `json:"trophies,omitempty"`
	BestTrophies         int    `json:"bestTrophies,omitempty"`
	WarStars             int    `json:"warStars,omitempty"`
	AttackWins           int    `json:"attackWins,omitempty"`
	DefenseWins          int    `json:"defenseWins,omitempty"`
	BuilderHallLevel     int    `json:"builderHallLevel,omitempty"`
	VersusTrophies       int    `json:"versusTrophies,omitempty"`
	BestVersusTrophies   int    `json:"bestVersusTrophies,omitempty"`
	VersusBattleWins     int    `json:"versusBattleWins,omitempty"`
	VersusBattleWinCount int    `json:"versusBattleWinCount,omitempty"`
	Role                 string `json:"role,omitempty"`
	Donations            int    `json:"donations,omitempty"`
	DonationsReceived    int    `json:"donationsReceived,omitempty"`
}
