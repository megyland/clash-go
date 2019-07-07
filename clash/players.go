package clash

import (
	"context"
	"fmt"
	"net/url"
)

// PlayersService handles communication with the players related
// methods of the Clash of Clan API.
type PlayersService service

// Player represents a Clash of Clan player
type (
	Player struct {
		Tag                  string        `json:"tag,omitempty"`
		Name                 string        `json:"name,omitempty"`
		TownHallLevel        int           `json:"townHallLevel,omitempty"`
		ExpLevel             int           `json:"expLevel,omitempty"`
		Trophies             int           `json:"trophies,omitempty"`
		BestTrophies         int           `json:"bestTrophies,omitempty"`
		WarStars             int           `json:"warStars,omitempty"`
		AttackWins           int           `json:"attackWins,omitempty"`
		DefenseWins          int           `json:"defenseWins,omitempty"`
		BuilderHallLevel     int           `json:"builderHallLevel,omitempty"`
		VersusTrophies       int           `json:"versusTrophies,omitempty"`
		BestVersusTrophies   int           `json:"bestVersusTrophies,omitempty"`
		VersusBattleWins     int           `json:"versusBattleWins,omitempty"`
		VersusBattleWinCount int           `json:"versusBattleWinCount,omitempty"`
		Role                 string        `json:"role,omitempty"`
		Donations            int           `json:"donations,omitempty"`
		DonationsReceived    int           `json:"donationsReceived,omitempty"`
		Clan                 Clan          `json:"clan,omitempty"`
		Achievements         []Achievement `json:"achievements,omitempty"`
		Troops               []Troop       `json:"troops,omitempty"`
		Heroes               []Hero        `json:"heroes,omitempty"`
		Spells               []Spell       `json:"spells,omitempty"`
	}

	Achievement struct {
		Name           string `json:"name,omitempty"`
		Stars          int    `json:"stars,omitempty"`
		Value          int    `json:"value,omitempty"`
		Target         int    `json:"target,omitempty"`
		Info           string `json:"info,omitempty"`
		CompletionInfo string `json:"completionInfo,omitempty"`
		Village        string `json:"village,omitempty"`
	}

	Troop struct {
		Name     string `json:"name,omitempty"`
		Level    int    `json:"level,omitempty"`
		MaxLevel int    `json:"maxLevel,omitempty"`
		Village  string `json:"village,omitempty"`
	}

	Hero struct {
		Name     string `json:"name,omitempty"`
		Level    int    `json:"level,omitempty"`
		MaxLevel int    `json:"maxLevel,omitempty"`
		Village  string `json:"village,omitempty"`
	}

	Spell struct {
		Name     string `json:"name,omitempty"`
		Level    int    `json:"level,omitempty"`
		MaxLevel int    `json:"maxLevel,omitempty"`
		Village  string `json:"village,omitempty"`
	}
)

// Get fetches a player by its tag
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/players/getPlayer
func (s *PlayersService) Get(ctx context.Context, playerTag string) (*Player, *Response, error) {
	playerTag = url.QueryEscape(playerTag)
	u := fmt.Sprintf("players/%v", playerTag)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	player := new(Player)
	resp, err := s.client.Do(ctx, req, player)
	if err != nil {
		return nil, resp, err
	}

	return player, resp, nil
}
