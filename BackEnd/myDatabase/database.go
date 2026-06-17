package myDatabase

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"sellTrainTicket/customType"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "train_ticket.db"
	}

	_, err := os.Stat(dbName)
	dbNotExists := os.IsNotExist(err)

	DB, err = sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("Ошибка при открытии/создании базы данных SQLite: %v", err)
	}

	if dbNotExists {
		log.Printf("Файл базы данных '%s' не найден.", dbName)
		log.Println("Выполнение скрипта create_db.sql.")
		sqlScript, err := os.ReadFile("myDatabase/create_db.sql")
		if err != nil {
			log.Fatalf("Ошибка при чтении файла create_db.sql: %v", err)
		}
		_, err = DB.Exec(string(sqlScript))
		if err != nil {
			log.Fatalf("Ошибка при выполнении скрипта create_db.sql: %v", err)
		}
		log.Println("База данных и таблицы успешно созданы.")
	} else {
		log.Printf("База данных '%s' уже существует.", dbName)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Не удалось установить соединение с базой данных: %v", err)
	}
	log.Println("Успешное подключение к базе данных.")
}

func AddUserDB(user customType.User) {
	_, err := DB.Exec("INSERT INTO users(login, password, email, phone, birth_date, name, surname, patronymic) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.Login, user.Password, user.Email, user.Phone, user.BirthDate, user.Name, user.Surname, user.Patronymic)
	if err != nil {
		log.Fatalf("Ошибка при добавление пользователя: %v", err)
	}
}

func UserExists(login, email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM Users WHERE login = ? OR email = ?)"
	err := DB.QueryRow(query, login, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func getUserID(login string) int {
	var id int
	if err := DB.QueryRow("SELECT id FROM Users WHERE login = ?", login).Scan(&id); err != nil {
		return 0
	}
	return id
}

func GetLoginUserDB(loginUser *customType.User) (string, error) {
	query := "SELECT password, name, surname, email, phone, birth_date, patronymic FROM Users WHERE login = ?"
	var passwordDB string
	err := DB.QueryRow(query, loginUser.Login).Scan(&passwordDB, &loginUser.Name, &loginUser.Surname,
		&loginUser.Email, &loginUser.Phone, &loginUser.BirthDate, &loginUser.Patronymic)
	if err != nil {
		return "", err
	}
	return passwordDB, nil
}

func GetRoutesDB(from, to, date string) ([]customType.RouteDB, error) {
	query := `SELECT r.id, r.sending, r.arrival, s_from.name AS from_station, s_to.name AS to_station, r.distance
	FROM Route r
	LEFT JOIN Station s_from ON r.from_station_id = s_from.id
	LEFT JOIN Station s_to ON r.to_station_id = s_to.id
	WHERE s_from.name = ? AND s_to.name = ? AND r.sending BETWEEN DATE(?) AND DATE(?, '4 days')
	ORDER BY r.sending`

	rows, err := DB.Query(query, from, to, date, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []customType.RouteDB
	for rows.Next() {
		var route customType.RouteDB
		if err := rows.Scan(&route.Id, &route.Sending, &route.Arrival, &route.FromStation, &route.ToStation, &route.Distance); err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}
	return routes, nil
}

func GetAllStationDB() ([]string, error) {
	var res []string
	rows, err := DB.Query("SELECT name FROM Station")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var station string
		if err := rows.Scan(&station); err != nil {
			return nil, err
		}
		res = append(res, station)
	}
	return res, nil
}

// Возвращает (массив carriage_id) состав поезда по id маршрута
func GetTrainDB(id int) ([]int, error) {
	rows, err := DB.Query("SELECT carriage_id FROM Train_composition WHERE route_id = ?", id)
	if err != nil {
		return nil, err
	}
	var res []int
	for rows.Next() {
		var train_id int
		if err := rows.Scan(&train_id); err != nil {
			return nil, err
		}
		res = append(res, train_id)
	}
	return res, nil
}

func GetCarriageSeatDB(id int) ([]customType.SeatDB, error) {
	query := `SELECT number, CASE WHEN occupied >= 1 THEN 1 ELSE 0 END FROM Seat WHERE carriage_id = ?`
	rows, err := DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	var res []customType.SeatDB
	for rows.Next() {
		var seat customType.SeatDB
		if err := rows.Scan(&seat.Number, &seat.Occupied); err != nil {
			return nil, err
		}
		res = append(res, seat)
	}
	return res, nil
}

// Возвращает true если все места свободны
func CheckingSeatsDB(seats map[int][]int) (bool, error) {
	for key, carriage := range seats {
		for _, seat := range carriage {
			var occupied bool
			if err := DB.QueryRow("SELECT occupied FROM Seat WHERE number = ? AND carriage_id = ?", seat, key).Scan(&occupied); err != nil {
				return false, err
			}
			if occupied {
				return false, nil
			}
		}
	}
	return true, nil
}

func BookingSeatsDB(login string, seats map[int][]int) error {
	id_user := getUserID(login)
	if id_user == 0 {
		return errors.New("User not found")
	}
	for key, carriage := range seats {
		for _, seat := range carriage {
			_, err := DB.Exec("UPDATE Seat SET occupied = ? WHERE number = ? AND carriage_id = ?", id_user, seat, key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func BuyTicketDB(login string, passengerData *[]customType.Ticket) error {
	id_user := getUserID(login)
	if id_user == 0 {
		return errors.New("User not found")
	}
	query := `INSERT INTO Ticket(route_id, user_id, seat_id, price, passenger_type, name, surname, patronymic, birth_date, pass_series, pass_number)
	VALUES (?, ?, (SELECT id FROM Seat WHERE number = ? AND carriage_id = ?), ?, ?, ?, ?, ?, ?, ?, ?)`
	for _, passenger := range *passengerData {
		_, err := DB.Exec(query, passenger.RouteId, id_user, passenger.SeatNumber, passenger.CarriageId, passenger.Price, passenger.PassengerType,
			passenger.Name, passenger.Surname, passenger.Patronymic, passenger.BirthDate, passenger.PassSeries, passenger.PassNumber)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTicketDB(login string) ([]map[string]interface{}, error) {
	id_user := getUserID(login)
	if id_user == 0 {
		return nil, errors.New("User not found")
	}
	query := `SELECT 
		t.id AS ticket_id, r.sending, r.arrival,
		dep.name AS name_station_from, arr.name AS name_station_to,
		s.carriage_id, s.number AS seat_number,
		t.price, t.passenger_type,
		t.name, t.surname, t.patronymic, t.birth_date,
		t.pass_series, t.pass_number
	FROM Ticket t
	JOIN Route r ON t.route_id = r.id
	JOIN Station dep ON r.from_station_id = dep.id
	JOIN Station arr ON r.to_station_id = arr.id
	JOIN Seat s ON t.seat_id = s.id
	WHERE t.user_id = ?`
	rows, err := DB.Query(query, id_user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []map[string]interface{}
	for rows.Next() {
		var (
			ticketID      int
			sending       string
			arrival       string
			stationFrom   string
			stationTo     string
			carriageID    int
			seatNumber    int
			price         int
			passengerType string
			name          string
			surname       string
			patronymic    string
			birthDate     string
			passSeries    int16
			passNumber    int
		)
		err := rows.Scan(
			&ticketID, &sending, &arrival,
			&stationFrom, &stationTo,
			&carriageID, &seatNumber,
			&price, &passengerType,
			&name, &surname, &patronymic, &birthDate,
			&passSeries, &passNumber,
		)
		if err != nil {
			return nil, err
		}
		ticket := map[string]interface{}{
			"ticket_id":      ticketID,
			"sending":        sending,
			"arrival":        arrival,
			"station_from":   stationFrom,
			"station_to":     stationTo,
			"carriage_id":    carriageID,
			"seat_number":    seatNumber,
			"price":          price,
			"passenger_type": passengerType,
			"name":           name,
			"surname":        surname,
			"patronymic":     patronymic,
			"birth_date":     birthDate,
			"pass_series":    passSeries,
			"pass_number":    passNumber,
		}
		res = append(res, ticket)
	}

	return res, nil
}

// !-- Для админов --!
func CreateStationDB(req []string) ([]int64, error) {
	var results_id []int64
	for _, v := range req {
		result, err := DB.Exec(`INSERT OR IGNORE INTO Station(name) VALUES (?);`, v)
		if err != nil {
			return nil, err
		}
		last_id, err := result.RowsAffected()
		if err != nil {
			return nil, err
		}
		results_id = append(results_id, last_id)
	}

	return results_id, nil
}

func CreateRouteDB(req customType.RouteDB) error {
	var from, to int

	if err := DB.QueryRow("SELECT id FROM Station WHERE name = ?", req.FromStation).Scan(&from); err != nil {
		return err
	}
	if err := DB.QueryRow("SELECT id FROM Station WHERE name = ?", req.ToStation).Scan(&to); err != nil {
		return err
	}
	if from == to {
		return errors.New("Две одинаковые станции")
	}
	query := `INSERT INTO Route(sending, arrival, from_station_id, to_station_id, distance) VALUES
	(datetime(?), datetime(?), ?, ?, ?);`
	if _, err := DB.Exec(query, req.Sending, req.Arrival, from, to, req.Distance); err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]*customType.User, error) {
	rows, err := DB.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*customType.User
	for rows.Next() {
		var user customType.User
		err := rows.Scan(&user.ID, &user.Login, &user.Password, &user.Email, &user.Phone,
			&user.BirthDate, &user.Name, &user.Surname, &user.Patronymic)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func GetUser(id int) (customType.User, error) {
	var user customType.User
	err := DB.QueryRow("SELECT * FROM Users WHERE id = ?", id).Scan(&user.ID, &user.Login, &user.Password,
		&user.Email, &user.Phone, &user.BirthDate, &user.Name, &user.Surname, &user.Patronymic)
	if err != nil {
		return user, err
	}
	return user, nil
}
