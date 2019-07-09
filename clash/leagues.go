package clash

import (
	"context"
	"fmt"
	"net/url"
)

type LeaguesService service

type (
	League struct {
		Id       int          `json:"id,omitempty"`
		Name     string       `json:"name,omitempty"`
		IconUrls UrlContainer `json:"iconUrls,omitempty"`
	}

	LeagueList struct {
		Items []League `json:"items,omitempty"`
	}

	LeagueSeason struct {
		Id string `json:"id,omitempty"`
	}

	LeagueSeasonList struct {
		Items []LeagueSeason `json:"items,omitempty"`
	}

	SeasonPlayerRanking struct {
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
	}

	SeasonPlayerRankingList struct {
		Items []SeasonPlayerRanking `json:"items,omitempty"`
	}
)

// List fetches a list of all leagues.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/leagues/getLeagues
func (s *LeaguesService) List(ctx context.Context, opt *Options) (*LeagueList, *Response, error) {
	u := "leagues"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	leagues := new(LeagueList)
	resp, err := s.client.Do(ctx, req, leagues)
	if err != nil {
		return nil, resp, err
	}

	return leagues, resp, nil
}

// Get fetches information about specific league.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/leagues/getLeague
func (s *LeaguesService) Get(ctx context.Context, leagueId string) (*League, *Response, error) {
	leagueId = url.QueryEscape(leagueId)
	u := fmt.Sprintf("leagues/%v", leagueId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	lea := new(League)
	resp, err := s.client.Do(ctx, req, lea)
	if err != nil {
		return nil, resp, err
	}

	return lea, resp, nil
}

// GetLeagueSeason fetches league seasons. Note that league season information is available only for Legend League.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/leagues/getLeagueSeasons
func (s *LeaguesService) GetLeagueSeason(ctx context.Context, leagueId string, opt *Options) (*LeagueSeasonList, *Response, error) {
	leagueId = url.QueryEscape(leagueId)
	u := fmt.Sprintf("leagues/%v/seasons", leagueId)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	leagueSeasons := new(LeagueSeasonList)
	resp, err := s.client.Do(ctx, req, leagueSeasons)
	if err != nil {
		return nil, resp, err
	}

	return leagueSeasons, resp, nil
}

// GetLeagueSeason fetches league season rankings. Note that league season information is available only for Legend League.
//
// Clash of Clan API docs: https://developer.clashofclans.com/api-docs/index.html#!/leagues/getLeagueSeasonRankings
func (s *LeaguesService) GetLeagueSeasonRankings(ctx context.Context, leagueId string, seasonId string, opt *Options) (*SeasonPlayerRankingList, *Response, error) {
	leagueId = url.QueryEscape(leagueId)
	seasonId = url.QueryEscape(seasonId)
	u := fmt.Sprintf("leagues/%v/seasons/%v", leagueId, seasonId)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	seasonPlayerRankingList := new(SeasonPlayerRankingList)
	resp, err := s.client.Do(ctx, req, seasonPlayerRankingList)
	if err != nil {
		return nil, resp, err
	}

	return seasonPlayerRankingList, resp, nil
}
