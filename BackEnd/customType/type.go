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

type UserLogin struct {
	Login      string `json:"login"`
	Password   string `json:"password"`
	Name       string
	Surname    string
	PasswordDB string
}

type RouteDB struct {
	Sending     string `json:"sending"`
	Arrival     string `json:"arrival"`
	FromStation string `json:"from_station"`
	ToStation   string `json:"to_station"`
	Distance    int    `json:"distance"`
}
