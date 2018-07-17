package wakatime

import (
	"encoding/json"
	"net/http"
	"time"
)

type Summaries struct {
	Data  []Summary `json:"data"`
	End   time.Time `json:"end"`
	Start time.Time `json:"start"`
}

type Summary struct {
	Categories       []Category        `json:"categories"`
	Dependencies     []Dependency      `json:"dependencies"`
	Editors          []Editor          `json:"editors"`
	GrandTotal       GrandTotal        `json:"grand_total"`
	Languages        []Language        `json:"languages"`
	OperatingSystems []OperatingSystem `json:"operating_systems"`
	Projects         []Project         `json:"projects"`
	Range            Range             `json:"range"`
}

type Category struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      int     `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds int     `json:"total_seconds"`
}

type Dependency struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Text         string  `json:"text"`
	TotalSeconds int     `json:"total_seconds"`
}

type Editor struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      int     `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds int     `json:"total_seconds"`
}

type GrandTotal struct {
	Digital      string `json:"digital"`
	Hours        int    `json:"hours"`
	Minutes      int    `json:"minutes"`
	Text         string `json:"text"`
	TotalSeconds int    `json:"total_seconds"`
}

type Language struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      int     `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds int     `json:"total_seconds"`
}

type OperatingSystem struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      int     `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds int     `json:"total_seconds"`
}

type Project struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      int     `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds int     `json:"total_seconds"`
}

type Range struct {
	Date     string    `json:"date"`
	End      time.Time `json:"end"`
	Start    time.Time `json:"start"`
	Text     string    `json:"text"`
	Timezone string    `json:"timezone"`
}

func (c *Client) GetSummaries(start, end string) (*Summaries, error) {

	req, err := http.NewRequest("GET", c.urlFor("/users/current/summaries").String(), nil)

	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("start", start)
	q.Add("end", end)
	req.URL.RawQuery = q.Encode()

	resp, err := c.request(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var summary Summaries
	err = json.NewDecoder(resp.Body).Decode(&summary)
	if err != nil {
		return nil, err
	}
	return &summary, nil
}
