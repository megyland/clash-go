package clash

import (
	"context"
	"fmt"
	"net/url"
)

type LocationsService service

type (
	Location struct {
		Id        int    `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		IsCountry bool   `json:"isCountry,omitempty"`
	}

	LocationList struct {
		Items []Location `json:"items,omitempty"`
	}

	ClanRanking struct {
		Tag              string       `json:"tag,omitempty"`
		Name             string       `json:"name,omitempty"`
		Location         Location     `json:"location,omitempty"`
		BadgeUrls        UrlContainer `json:"badgeUrls,omitempty"`
		ClanLevel        int          `json:"clanLevel,omitempty"`
		ClanPoints       int          `json:"clanPoints,omitempty"`
		ClanVersusPoints int          `json:"clanVersusPoints,omitempty"`
		Members          int          `json:"members,omitempty"`
		Rank             int          `json:"rank,omitempty"`
		PreviousRank     int          `json:"previousRank,omitempty"`
	}

	ClanRankingList struct {
		Items []ClanRanking `json:"items,omitempty"`
	}

	PlayerRanking struct {
		Tag            string `json:"tag,omitempty"`
		Name           string `json:"name,omitempty"`
		ExpLevel       int    `json:"expLevel,omitempty"`
		League         League `json:"league,omitempty"`
		Trophies       int    `json:"trophies,omitempty"`
		VersusTrophies int    `json:"versusTrophies,omitempty"`
		AttackWins     int    `json:"attackWins,omitempty"`
		DefenseWins    int    `json:"defenseWins,omitempty"`
		Clan           Clan   `json:"clan,omitempty"`
		Rank           int    `json:"rank,omitempty"`
		PreviousRank   int    `json:"previousRank,omitempty"`
	}

	PlayerRankingList struct {
		Items []PlayerRanking `json:"items,omitempty"`
	}

	PlayerVersusRanking struct {
		Tag              string `json:"tag,omitempty"`
		Name             string `json:"name,omitempty"`
		ExpLevel         int    `json:"exp_level,omitempty"`
		Rank             int    `json:"rank,omitempty"`
		PreviousRank     int    `json:"previousRank,omitempty"`
		VersusTrophies   int    `json:"versusTrophies,omitempty"`
		Clan             Clan   `json:"clan,omitempty"`
		VersusBattleWins int    `json:"versusBattleWins,omitempty"`
	}

	PlayerVersusRankingList struct {
		Items []PlayerVersusRanking `json:"items,omitempty"`
	}
)

// List fetches a list of all available locations.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/locations/getLocations
func (s *LocationsService) List(ctx context.Context, opt *Options) (*LocationList, *Response, error) {
	u := "locations"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	locations := new(LocationList)
	resp, err := s.client.Do(ctx, req, locations)
	if err != nil {
		return nil, resp, err
	}

	return locations, resp, nil
}

// Get fetches information about specific location.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/locations/getLocation
func (s *LocationsService) Get(ctx context.Context, locationId string) (*Location, *Response, error) {
	locationId = url.QueryEscape(locationId)
	u := fmt.Sprintf("locations/%v", locationId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	location := new(Location)
	resp, err := s.client.Do(ctx, req, location)
	if err != nil {
		return nil, resp, err
	}

	return location, resp, nil
}

// GetClanRankings fetches clan rankings for a specific location.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/locations/getClanRanking
func (s *LocationsService) GetClanRankings(ctx context.Context, locationId string, opt *Options) (*ClanRankingList, *Response, error) {
	locationId = url.QueryEscape(locationId)
	u := fmt.Sprintf("locations/%v/rankings/clans", locationId)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	clanRankingList := new(ClanRankingList)
	resp, err := s.client.Do(ctx, req, clanRankingList)
	if err != nil {
		return nil, resp, err
	}

	return clanRankingList, resp, nil
}

// GetPlayerRankings fetches player rankings for a specific location.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/locations/getPlayerRanking
func (s *LocationsService) GetPlayerRankings(ctx context.Context, locationId string, opt *Options) (*PlayerRankingList, *Response, error) {
	locationId = url.QueryEscape(locationId)
	u := fmt.Sprintf("locations/%v/rankings/players", locationId)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	playerRankingList := new(PlayerRankingList)
	resp, err := s.client.Do(ctx, req, playerRankingList)
	if err != nil {
		return nil, resp, err
	}

	return playerRankingList, resp, nil
}

// GetClanVersusRankings fetches clan versus rankings for a specific location.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/locations/getClanVersusRanking
func (s *LocationsService) GetClanVersusRankings(ctx context.Context, locationId string, opt *Options) (*ClanRankingList, *Response, error) {
	locationId = url.QueryEscape(locationId)
	u := fmt.Sprintf("locations/%v/rankings/clans-versus", locationId)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	clanRankingList := new(ClanRankingList)
	resp, err := s.client.Do(ctx, req, clanRankingList)
	if err != nil {
		return nil, resp, err
	}

	return clanRankingList, resp, nil
}

// GetClanVersusRankings fetches player versus rankings for a specific location.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/locations/getPlayerVersusRanking
func (s *LocationsService) GetPlayerVersusRankings(ctx context.Context, locationId string, opt *Options) (*PlayerVersusRankingList, *Response, error) {
	locationId = url.QueryEscape(locationId)
	u := fmt.Sprintf("locations/%v/rankings/players-versus", locationId)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	playerVersusRankingList := new(PlayerVersusRankingList)
	resp, err := s.client.Do(ctx, req, playerVersusRankingList)
	if err != nil {
		return nil, resp, err
	}

	return playerVersusRankingList, resp, nil
}
