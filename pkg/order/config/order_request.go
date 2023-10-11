package config

type GetOrdersRequest struct {
	StartDate   string   `json:"startDate" binding:"required"`
	EndDate     string   `json:"endDate" binding:"required"`
	RecordLimit int      `json:"recordLimit" binding:"required"`
	Offset      int      `json:"offset"`
	Statuses    []string `json:"statuses" binding:"required"`
}
