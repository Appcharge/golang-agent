package config

type GetAnalyticsRequest struct {
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Metrics    []string `json:"metrics"`
	IncomeType string   `json:"incomeType"`
}

type GetAnalyticsResponse struct {
	Name   string        `json:"name"`
	Type   string        `json:"type"`
	Values []MetricValue `json:"values"`
}

type MetricValue struct {
	Value float64 `json:"value"`
	Date  string  `json:"date"`
}
