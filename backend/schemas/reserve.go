package schemas

type ReservationRequest struct {
	IdUser    int     `json:"idUser" binding:"required"`
	IdService int     `json:"idService" binding:"required"`
	IdOrder   int     `json:"idOrder" binding:"required"`
	Amount    float32 `json:"amount" binding:"required"`
	Date      string  `json:"date" binding:"required"`
}

type ConfirmRequest struct {
	IdUser    int     `json:"idUser" binding:"required"`
	IdService int     `json:"idService" binding:"required"`
	IdOrder   int     `json:"idOrder" binding:"required"`
	Amount    float32 `json:"amount" binding:"required"`
	Date      string  `json:"date" binding:"required"`
	Message   string  `json:"message" binding:"required"`
}

type CancelRequest struct {
	IdUser    int    `json:"idUser" binding:"required"`
	IdService int    `json:"idService" binding:"required"`
	IdOrder   int    `json:"idOrder" binding:"required"`
	Date      string `json:"date" binding:"required"`
}
