package wakatime

import (
	"encoding/json"
	"time"
)

type Summaries struct {
	Data struct {
		BestDay struct {
			CreatedAt    time.Time   `json:"created_at"`
			Date         string      `json:"date"`
			ID           string      `json:"id"`
			ModifiedAt   interface{} `json:"modified_at"`
			TotalSeconds int         `json:"total_seconds"`
		} `json:"best_day"`
		Categories []struct {
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds int     `json:"total_seconds"`
		} `json:"categories"`
		CreatedAt             time.Time `json:"created_at"`
		DailyAverage          int       `json:"daily_average"`
		DaysIncludingHolidays int       `json:"days_including_holidays"`
		DaysMinusHolidays     int       `json:"days_minus_holidays"`
		Dependencies          []struct {
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds int     `json:"total_seconds"`
		} `json:"dependencies"`
		Editors []struct {
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds int     `json:"total_seconds"`
		} `json:"editors"`
		End                       time.Time `json:"end"`
		Holidays                  int       `json:"holidays"`
		HumanReadableDailyAverage string    `json:"human_readable_daily_average"`
		HumanReadableTotal        string    `json:"human_readable_total"`
		ID                        string    `json:"id"`
		IsAlreadyUpdating         bool      `json:"is_already_updating"`
		IsCodingActivityVisible   bool      `json:"is_coding_activity_visible"`
		IsOtherUsageVisible       bool      `json:"is_other_usage_visible"`
		IsStuck                   bool      `json:"is_stuck"`
		IsUpToDate                bool      `json:"is_up_to_date"`
		Languages                 []struct {
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds int     `json:"total_seconds"`
		} `json:"languages"`
		ModifiedAt       time.Time `json:"modified_at"`
		OperatingSystems []struct {
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds int     `json:"total_seconds"`
		} `json:"operating_systems"`
		Project  interface{} `json:"project"`
		Projects []struct {
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds int     `json:"total_seconds"`
		} `json:"projects"`
		Range        string      `json:"range"`
		Start        time.Time   `json:"start"`
		Status       string      `json:"status"`
		Timeout      int         `json:"timeout"`
		Timezone     string      `json:"timezone"`
		TotalSeconds int         `json:"total_seconds"`
		UserID       string      `json:"user_id"`
		Username     interface{} `json:"username"`
		WritesOnly   bool        `json:"writes_only"`
	} `json:"data"`
}

func (c *Client) GetLast7Days() (*Summaries, error) {
	resp, err := c.Get("/users/current/stats/last_7_days")
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
