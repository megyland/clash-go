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

	clanMembers, _, err := client.Clans.GetMembers(context.Background(), "#28GQRLV2J", &clash.Options{Limit: 1})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(prettyPrint(clanMembers))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
