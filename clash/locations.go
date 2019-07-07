package clash

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
)
