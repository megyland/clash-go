package clash

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
)
