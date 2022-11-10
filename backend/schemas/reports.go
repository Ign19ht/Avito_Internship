package schemas

type ReportRequest struct {
	Month int `json:"month" binding:"required"`
	Year  int `json:"year" binding:"required"`
}

type ReportResponse struct {
	Link string `json:"link"`
}

type HistoryRequest struct {
	Id        int    `json:"id" binding:"required"`
	SortType  string `json:"sortType" binding:"required"`
	Amount    int    `json:"amount" binding:"required"`
	OrderType string `json:"orderType" binding:"required"`
	Offset    int    `json:"offset"`
}

type HistoryResponse struct {
	History []HistoryResponseItem `json:"history"`
}

type HistoryResponseItem struct {
	Date    string  `json:"date"`
	Amount  float32 `json:"amount"`
	Message string  `json:"message"`
}
