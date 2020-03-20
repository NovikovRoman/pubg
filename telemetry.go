package pubg

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Telemetry structure.
type Telemetry struct {
	index int
	Raw   []json.RawMessage
}
// NewTelemetryFromBytes create new Telemetry structure from bytes.
func NewTelemetryFromBytes(b []byte) (telemetry *Telemetry, err error) {
	var raw []json.RawMessage
	err = json.Unmarshal(b, &raw)
	telemetry = &Telemetry{
		Raw: raw,
	}
	return
}
// NewTelemetryFromFile create new Telemetry structure from file.
func NewTelemetryFromFile(filename string) (telemetry *Telemetry, err error) {
	var b []byte
	if b, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	telemetry, err = NewTelemetryFromBytes(b)
	return
}
// NewTelemetryFromURL create new Telemetry structure from url.
func NewTelemetryFromURL(url string, transport *http.Transport) (telemetry *Telemetry, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	request.Header.Add("Accept-Encoding", "gzip")

	if transport == nil {
		transport = &http.Transport{}
	}

	client := http.Client{
		Transport: transport,
	}

	resp, err := client.Do(request)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%d %s", resp.StatusCode, resp.Status)
		return
	}

	defer func() {
		if derr := resp.Body.Close(); derr != nil {
			err = derr
		}
	}()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer func() {
			if derr := reader.Close(); derr != nil {
				err = derr
			}
		}()

	default:
		reader = resp.Body
	}

	var b []byte
	if b, err = ioutil.ReadAll(reader); err != nil {
		return
	}
	telemetry, err = NewTelemetryFromBytes(b)
	return
}
// Len returns the number of events.
func (t *Telemetry) Len() int {
	return len(t.Raw)
}
// Rewind reset pointer.
func (t *Telemetry) Rewind() {
	t.index = 0
}
// SetIndex set pointer.
func (t *Telemetry) SetIndex(index int) error {
	if !t.validIndex(t.index) {
		return errors.New("Invalid argument. ")
	}

	t.index = index
	return nil
}
// Next returns the next event (date, type, raw).
func (t *Telemetry) Next() (eventDate time.Time, eventType string, raw json.RawMessage, ok bool, err error) {
	if !t.validIndex(t.index) {
		return
	}

	var info struct {
		DateRaw string `json:"_D"`
		Type    string `json:"_T"`
	}

	raw = t.Raw[t.index]
	if err = json.Unmarshal(raw, &info); err != nil {
		return
	}

	ok = true
	t.index++

	eventDate, _ = time.Parse(layoutDateTime, info.DateRaw)
	eventType = info.Type
	return
}

func (t Telemetry) validIndex(index int) bool {
	return index >= 0 && index < t.Len()
}
