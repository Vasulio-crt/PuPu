package customType

type User struct {
	ID         int
	Login      string `json:"login"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type RouteDB struct {
	Sending     string `json:"sending"`
	Arrival     string `json:"arrival"`
	FromStation string `json:"from_station"`
	ToStation   string `json:"to_station"`
	Distance    int    `json:"distance"`
}
