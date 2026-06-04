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

type UserLogin struct {
	Login      string `json:"login"`
	Password   string `json:"password"`
	Name       string
	Surname    string
	PasswordDB string
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
