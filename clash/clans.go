package clash

import (
	"context"
	"fmt"
	"net/url"
)

type ClansService service

type (
	Clan struct {
		Tag              string       `json:"tag,omitempty"`
		Name             string       `json:"name,omitempty"`
		Location         Location     `json:"location,omitempty"`
		ClanLevel        int          `json:"clanLevel,omitempty"`
		ClanPoints       int          `json:"clanPoints,omitempty"`
		ClanVersusPoints int          `json:"clanVersusPoints,omitempty"`
		BadgeUrls        UrlContainer `json:"badgeUrls,omitempty"`
		Members          int          `json:"members,omitempty"`
	}

	ClanMember struct {
		Player
		ClanRank         int `json:"clanRank,omitempty"`
		PreviousClanRank int `json:"previousClanRank,omitempty"`
	}

	ClanMemberList struct {
		Items []ClanMember `json:"items,omitempty"`
	}

	WarLogList struct {
		Items []WarLog `json:"items,omitempty"`
	}

	WarLog struct {
		Result   string  `json:"result,omitempty"`
		EndTime  string  `json:"endTime,omitempty"`
		TeamSize int     `json:"teamSize,omitempty"`
		Clan     WarClan `json:"clan,omitempty"`
		Opponent WarClan `json:"opponent,omitempty"`
	}

	WarClan struct {
		Tag       string          `json:"tag,omitempty"`
		Name      string          `json:"name,omitempty"`
		BadgeUrls UrlContainer    `json:"badgeUrls,omitempty"`
		ClanLevel int             `json:"clanLevel,omitempty"`
		Attacks   int             `json:"attacks,omitempty"`
		Stars     int             `json:"stars,omitempty"`
		ExpEarned int             `json:"expEarned,omitempty"`
		Members   []ClanWarMember `json:"members,omitempty"`
	}

	CurrentWar struct {
		State                string  `json:"state,omitempty"`
		TeamSize             int     `json:"teamSize,omitempty"`
		PreparationStartTime string  `json:"preparationStartTime,omitempty"`
		StartTime            string  `json:"startTime,omitempty"`
		EndTime              string  `json:"endTime,omitempty"`
		Clan                 WarClan `json:"clan,omitempty"`
		Opponent             WarClan `json:"opponent,omitempty"`
	}

	CurrentWarLeagueGroup struct {
		Tag    string   `json:"tag,omitempty"`
		State  string   `json:"state,omitempty"`
		Season string   `json:"season,omitempty"`
		Clans  WarClan  `json:"clans,omitempty"`
		Rounds []WarTag `json:"rounds,omitempty"`
	}

	WarTag struct {
		Tag           string `json:"tag,omitempty"`
		Name          string `json:"name,omitempty"`
		TownHallLevel int    `json:"townHallLevel,omitempty"`
	}

	ClanWarMember struct {
		Tag                string          `json:"tag,omitempty"`
		Name               string          `json:"name,omitempty"`
		TownHallLevel      int             `json:"townHallLevel,omitempty"`
		MapPosition        int             `json:"mapPosition,omitempty"`
		Attacks            []ClanWarAttack `json:"attacks,omitempty"`
		OpponentAttacks    int             `json:"opponentAttacks,omitempty"`
		BestOpponentAttack ClanWarAttack   `json:"bestOpponentAttack,omitempty"`
	}

	ClanWarAttack struct {
		AttackerTag           string `json:"attackerTag,omitempty"`
		DefenderTag           string `json:"defenderTag,omitempty"`
		Stars                 int    `json:"stars,omitempty"`
		DestructionPercentage int    `json:"destructionPercentage,omitempty"`
		Order                 int    `json:"order,omitempty"`
	}

	UrlContainer struct {
		Small  string `json:"small,omitempty"`
		Large  string `json:"large,omitempty"`
		Medium string `json:"medium,omitempty"`
	}

	Options struct {
		Limit  int `url:"limit,omitempty"`
		After  int `url:"after,omitempty"`
		Before int `url:"before,omitempty"`
	}

	// ClanOptions specifies the optional parameters for clan search
	SearchOptions struct {
		Options
		Name          string `url:"name,omitempty"`
		WarFrequency  string `url:"warFrequency,omitempty"`
		MinMembers    int    `url:"minMembers,omitempty"`
		MaxMembers    int    `url:"maxMembers,omitempty"`
		MinClanPoints int    `url:"minClanPoints,omitempty"`
		MinClanLevel  string `url:"minClanLevel,omitempty"`
	}

	ClanList struct {
		Items []Clan `json:"items,omitempty"`
	}

	ClanWarLeagueWar struct {
		Tag                  string `json:"tag,omitempty"`
		State                string `json:"state,omitempty"`
		TeamSize             int    `json:"teamSize,omitempty"`
		PreparationStartTime string `json:"preparationStartTime,omitempty"`
		StartTime            string `json:"startTime,omitempty"`
		EndTime              string `json:"endTime,omitempty"`
		Clan                 string `json:"clan,omitempty"`
		Opponent             string `json:"opponent,omitempty"`
	}

	ClanWarLeagueWarClan struct {
		Tag                   string        `json:"tag,omitempty"`
		Name                  string        `json:"name,omitempty"`
		BadgeUrls             UrlContainer  `json:"badgeUrls,omitempty"`
		ClanLevel             int           `json:"clanLevel,omitempty"`
		Attacks               int           `json:"attacks,omitempty"`
		Stars                 int           `json:"stars,omitempty"`
		DestructionPercentage int           `json:"destructionPercentage,omitempty"`
		Members               ClanWarMember `json:"members,omitempty"`
	}
)

// List the repositories for a user. Passing the empty string will list
// repositories for the authenticated user.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/clans/searchClans
func (s *ClansService) Search(ctx context.Context, opt *SearchOptions) (*ClanList, *Response, error) {
	u := "clans"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	clans := new(ClanList)
	resp, err := s.client.Do(ctx, req, &clans)
	if err != nil {
		return nil, resp, err
	}

	return clans, resp, nil
}

// Get fetches information about a single clan by clan tag.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/clans/getClan
func (s *ClansService) Get(ctx context.Context, clanTag string) (*Clan, *Response, error) {
	clanTag = url.QueryEscape(clanTag)
	u := fmt.Sprintf("clans/%v", clanTag)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	clan := new(Clan)
	resp, err := s.client.Do(ctx, req, clan)
	if err != nil {
		return nil, resp, err
	}

	return clan, resp, nil
}

// GetMembers fetches a list of clan members.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/clans/getClanMembers
func (s *ClansService) GetMembers(ctx context.Context, clanTag string, opt *Options) (*ClanMemberList, *Response, error) {
	clanTag = url.QueryEscape(clanTag)
	u := fmt.Sprintf("clans/%v/members", clanTag)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	members := new(ClanMemberList)
	resp, err := s.client.Do(ctx, req, members)
	if err != nil {
		return nil, resp, err
	}

	return members, resp, nil
}

// GetWarLog fetches clan's clan war log.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/clans/getClanWarLog
func (s *ClansService) GetWarLog(ctx context.Context, clanTag string, opt *Options) (*WarLogList, *Response, error) {
	clanTag = url.QueryEscape(clanTag)
	u := fmt.Sprintf("clans/%v/warlog", clanTag)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	warLogs := new(WarLogList)
	resp, err := s.client.Do(ctx, req, warLogs)
	if err != nil {
		return nil, resp, err
	}

	return warLogs, resp, nil
}

// GetCurrentWar fetches  information about clan's current clan war.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/clans/getCurrentWar
func (s *ClansService) GetCurrentWar(ctx context.Context, clanTag string) (*CurrentWar, *Response, error) {
	clanTag = url.QueryEscape(clanTag)
	u := fmt.Sprintf("clans/%v/currentwar", clanTag)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	currentWar := new(CurrentWar)
	resp, err := s.client.Do(ctx, req, currentWar)
	if err != nil {
		return nil, resp, err
	}

	return currentWar, resp, nil
}

// GetCurrentWarLeagueGroup fetches information about clan's current clan war league group.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/clans/getClanWarLeagueGroup
func (s *ClansService) GetCurrentWarLeagueGroup(ctx context.Context, clanTag string) (*CurrentWarLeagueGroup, *Response, error) {
	clanTag = url.QueryEscape(clanTag)
	u := fmt.Sprintf("clans/%v/currentwar/leaguegroup", clanTag)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	currentWarLeagueGroup := new(CurrentWarLeagueGroup)
	resp, err := s.client.Do(ctx, req, currentWarLeagueGroup)
	if err != nil {
		return nil, resp, err
	}

	return currentWarLeagueGroup, resp, nil
}

// GetClanLeagueWar fetches information about clan's current clan war league war.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/clans/getClanWarLeagueWar
func (s *ClansService) GetClanLeagueWar(ctx context.Context, warTag string) (*ClanWarLeagueWar, *Response, error) {
	warTag = url.QueryEscape(warTag)
	u := fmt.Sprintf("clanwarleagues/wars/%v", warTag)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	clanWarLeagueWar := new(ClanWarLeagueWar)
	resp, err := s.client.Do(ctx, req, clanWarLeagueWar)
	if err != nil {
		return nil, resp, err
	}

	return clanWarLeagueWar, resp, nil
}
