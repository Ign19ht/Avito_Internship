package models

type Balance struct {
	Id      int     `repository:"id"`
	Balance float32 `repository:"balance"`
}

type Reservation struct {
	IdUser    int     `repository:"id_user"`
	IdService int     `repository:"id_service"`
	IdOrder   int     `repository:"id_order"`
	Amount    float32 `repository:"amount"`
	Status    string  `repository:"status"`
	Date      string  `repository:"date"`
}

type History struct {
	IdUser  int     `repository:"id_user"`
	Date    string  `repository:"date"`
	Amount  float32 `repository:"amount"`
	Message string  `repository:"message"`
}

type Reports struct {
	Date string `repository:"date"`
	Link string `repository:"link"`
}
