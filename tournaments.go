package pubg

import (
	"encoding/json"
	"fmt"
	"time"
)

type Tournaments struct {
	Data []struct {
		// Identifier for this object type ("tournament")
		Type string `json:"type"`
		// Tournament ID
		ID string `json:"id"`

		Attributes struct {
			CreatedAtRaw string    `json:"createdAt"`
			CreatedAt    time.Time `json:"-"`
		} `json:"attributes"`
	} `json:"data"`

	Links links `json:"links"`
}

type Tournament struct {
	Data struct {
		// Identifier for this object type ("tournament")
		Type string `json:"type"`
		// Tournament ID
		ID string `json:"id"`

		Relationships struct {
			Matches dataArrayIDAndType `json:"matches"`
		} `json:"relationships"`
	}

	Included []struct {
		// Identifier for this object type ("match")
		Type string `json:"type"`
		// Match ID
		ID string `json:"id"`

		Attributes struct {
			CreatedAtRaw string    `json:"createdAt"`
			CreatedAt    time.Time `json:"-"`
		} `json:"attributes"`
	} `json:"included"`

	Links links `json:"links"`
}

// Get information for a single tournament.
func (c Client) Tournament(tournamentID string) (tournament *Tournament, err error) {
	b, err := c.requestGET(EmptyPlatform, fmt.Sprintf("/tournaments/%s", tournamentID))
	if err != nil {
		return
	}

	tournament = &Tournament{}
	err = json.Unmarshal(b, tournament)

	for i := range tournament.Included {
		tournament.Included[i].Attributes.CreatedAt, _ = time.Parse(
			layoutDateTime,
			tournament.Included[i].Attributes.CreatedAtRaw,
		)
	}
	return
}

// Get the list of available tournaments.
func (c Client) Tournaments() (tournaments *Tournaments, err error) {
	b, err := c.requestGET(EmptyPlatform, "/tournaments")
	if err != nil {
		return
	}

	tournaments = &Tournaments{}
	err = json.Unmarshal(b, tournaments)

	for i := range tournaments.Data {
		tournaments.Data[i].Attributes.CreatedAt, _ = time.Parse(
			layoutDateTime,
			tournaments.Data[i].Attributes.CreatedAtRaw,
		)
	}
	return
}
