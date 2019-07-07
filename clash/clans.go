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

// Get fetches a clan by its tag.
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

// GetMembers fetches all members of a clan by its tag.
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
