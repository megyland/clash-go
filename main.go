package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/megyland/clash-go/clash"
	"net/http"
)

func main() {
	client := clash.NewClient(new(http.Client))

	/*player, _, err := client.Players.Get(context.Background(), "#8UCG8C0U")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(player))*/

	/*clan, _, err := client.Clans.Get(context.Background(), "#28GQRLV2J")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(clan))*/

	/*opt := &clash.SearchOptions{
		Name:    "Team4Playing",
		Options: clash.Options{Limit: 1},
	}
	clanResults, _, err := client.Clans.Search(context.Background(), opt)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(clanResults))*/

	/*clanMembers, _, err := client.Clans.GetMembers(context.Background(), "#28GQRLV2J", &clash.Options{Limit: 1})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(clanMembers))*/

	// TODO: Test with a public clan
	/*warLog, _, err := client.Clans.GetWarLog(context.Background(), "#28GQRLV2J", &clash.Options{Limit: 3})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(warLog))*/

	// TODO: Test with a public clan
	/*currentWar, _, err := client.Clans.GetCurrentWar(context.Background(), "#28GQRLV2J")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(currentWar))*/

	// TODO: Test with a public clan
	/*currentWarLeagueGroup, _, err := client.Clans.GetCurrentWarLeagueGroup(context.Background(), "#28GQRLV2J")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(currentWarLeagueGroup))*/

	// TODO: Test with a public clan
	/*currentWarLeagueWar, _, err := client.Clans.GetClanLeagueWar(context.Background(), "#28GQRLV2J")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(currentWarLeagueWar))*/

	/*locationList, _, err := client.Locations.List(context.Background(), nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(locationList))*/

	/*location, _, err := client.Locations.Get(context.Background(), "32000087")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(location))*/

	/*locationRankingClan, _, err := client.Locations.GetClanRankings(context.Background(), "32000087", &clash.Options{Limit: 10})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(locationRankingClan))*/

	/*locationRankingPlayer, _, err := client.Locations.GetPlayerRankings(context.Background(), "32000087", &clash.Options{Limit: 10})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(locationRankingPlayer))*/

	/*locationRankingClanVersus, _, err := client.Locations.GetClanVersusRankings(context.Background(), "32000087", &clash.Options{Limit: 10})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(locationRankingClanVersus))*/

	/*locationRankingPlayerVersus, _, err := client.Locations.GetPlayerVersusRankings(context.Background(), "32000087", &clash.Options{Limit: 10})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(locationRankingPlayerVersus))*/

	/*leagues, _, err := client.Leagues.List(context.Background(), nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(leagues))*/

	/*league, _, err := client.Leagues.Get(context.Background(), "29000000")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(league))*/

	leagueSeason, _, err := client.Leagues.GetLeagueSeason(context.Background(), "29000022", nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(leagueSeason))

	seasonPlayersRanking, _, err := client.Leagues.GetLeagueSeasonRankings(context.Background(), "29000022", "2019-02", nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(seasonPlayersRanking))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
