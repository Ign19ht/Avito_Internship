package models

type Balance struct {
	Id      int     `db:"id"`
	Balance float32 `db:"balance"`
}

type Reservation struct {
	IdUser    int     `db:"id_user"`
	IdService int     `db:"id_service"`
	IdOrder   int     `db:"id_order"`
	Amount    float32 `db:"amount"`
	Status    string  `db:"status"`
	Date      string  `db:"date"`
}

type History struct {
	IdUser  int     `db:"id_user"`
	Date    string  `db:"date"`
	Amount  float32 `db:"amount"`
	Message string  `db:"message"`
}

type Reports struct {
	Date string `db:"date"`
	Link string `db:"link"`
}
