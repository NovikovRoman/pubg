package pubg

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Samples struct {
	Data struct {
		// Identifier for this object type ("sample")
		Type string `json:"type"`
		// Sample ID
		ID string `json:"id"`

		// Sample objects contain the ID of a match.
		Attributes struct {
			CreatedAtRaw string    `json:"createdAt"`
			CreatedAt    time.Time `json:"-"`
			TitleId      string    `json:"titleId"`
			// Platform ID
			ShardId string `json:"shardId"`
		} `json:"attributes"`

		// Array:
		//     type - Identifier for this object type ("match")
		//     id - MatchID. Used to find the full match object in the included array
		Relationships struct {
			Matches dataArrayIDAndType `json:"matches"`
		} `json:"relationships"`
	} `json:"data"`
}

// Get a list of sample matches.
//
// The number of matches per shard may vary. Requests for samples need to be at least 24hrs in the past UTC time using
// the `start` query parameter. The default if not specified is the latest sample.
//
// The starting search date in UTC.
//
// Usage: 2018-01-01T00:00:00Z // Note this time will not work
//
// Default: time.Now() - 24hrs
func (c Client) Samples(platform Platform, start time.Time) (samples *Samples, err error) {
	t := ""
	if !start.IsZero() {
		if start.Before(time.Date(2018, 01, 01, 0, 0, 0, 0, time.UTC)) {
			err = errors.New("Date must be later than 2018-01-01T00:00:00Z ")
			return
		}
		t = start.Format("2006-01-02T15:04:05Z")
	}

	b, _, err := c.requestGET(platform, fmt.Sprintf("/samples?filter[createdAt-start]=%s", t))
	if err != nil {
		return
	}

	samples = &Samples{}
	err = json.Unmarshal(b, samples)

	samples.Data.Attributes.CreatedAt, _ = time.Parse(layoutDateTime, samples.Data.Attributes.CreatedAtRaw)
	return
}
