package schemas

type AccrualRequest struct {
	Id      int     `json:"id" binding:"required"`
	Amount  float32 `json:"amount" binding:"required"`
	Date    string  `json:"date" binding:"required"`
	Message string  `json:"message" binding:"required"`
}

type BalanceResponse struct {
	Balance float32 `json:"balance"`
}

type BalanceRequest struct {
	Id int `json:"id" binding:"required"`
}

type SendRequest struct {
	IdFrom  int     `json:"idFrom" binding:"required"`
	IdTo    int     `json:"idTo" binding:"required"`
	Amount  float32 `json:"amount" binding:"required"`
	Date    string  `json:"date" binding:"required"`
	Message string  `json:"message" binding:"required"`
}
