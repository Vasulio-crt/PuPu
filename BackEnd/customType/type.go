package customType

type User struct {
	ID         int
	Login      string `json:"login"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	BirthDate  string `json:"birth_date"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type RouteDB struct {
	Id          int    `json:"id"`
	Sending     string `json:"sending"`
	Arrival     string `json:"arrival"`
	FromStation string `json:"from_station"`
	ToStation   string `json:"to_station"`
	Distance    int    `json:"distance"`
}

type SeatDB struct {
	Number   int  `json:"num"`
	Occupied bool `json:"occ"`
}

type Ticket struct {
	RouteId       int    `json:"route_id"`
	Price         int    `json:"price"`
	PassengerType string `json:"passenger_type"`
	SeatNumber    int    `json:"seat_number"`
	CarriageId    int    `json:"carriage_id"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Patronymic    string `json:"patronymic"`
	BirthDate     string `json:"birth_date"`
	PassSeries    int16  `json:"passport_series"`
	PassNumber    int    `json:"passport_number"`
}
