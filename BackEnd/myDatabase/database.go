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
	dbExists := !os.IsNotExist(err)

	DB, err = sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("Ошибка при открытии/создании базы данных SQLite: %v", err)
	}

	if !dbExists {
		log.Printf("Файл базы данных '%s' не найден. Создание таблиц...", dbName)
		log.Println("Выполнение скрипта create_db.sql...")
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
		log.Printf("База данных '%s' уже существует. Подключение...", dbName)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Не удалось установить соединение с базой данных: %v", err)
	}
	log.Println("Успешное подключение к базе данных.")
}

func AddUserDB(user customType.User) {
	_, err := DB.Exec("INSERT INTO users (login, password, email, name, surname, patronymic) VALUES (?, ?, ?, ?, ?, ?)",
		user.Login, user.Password, user.Email, user.Name, user.Surname, user.Patronymic)
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

func GetRoutesDB(from, to, date string) ([]customType.RouteDB, error) {
	query := `SELECT sending, arrival, S1.name AS from_station, S2.name AS to_station, distance
		FROM Route
		JOIN Station AS S1 ON Route.from_station_id = S1.id
		JOIN Station AS S2 ON Route.to_station_id = S2.id
		WHERE S1.name = ?
			AND S2.name = ?
			AND sending >= datetime(?)
		LIMIT 15`

	rows, err := DB.Query(query, from, to, date, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []customType.RouteDB
	for rows.Next() {
		var route customType.RouteDB
		if err := rows.Scan(&route.Sending, &route.Arrival, &route.FromStation, &route.ToStation, &route.Distance); err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}
	return routes, nil
}

func CreateStationDB(req []string) ([]int64, error) {
	var results_id []int64
	for _, v := range req {
		result, err := DB.Exec("INSERT INTO Station(name) VALUES (?)", v)
		if err != nil {
			return nil, err
		}
		last_id, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}
		results_id = append(results_id, last_id)
	}

	return results_id, nil
}

func CreateRouteDB(req customType.RouteDB) error {
	var ids []int
	rows, err := DB.Query("SELECT id FROM Station WHERE name IN (?, ?);", req.FromStation, req.ToStation)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return err
		}
		ids = append(ids, id)
	}
	rows.Close()
	if len(ids) < 2 {
		return errors.New("ids stations < 2")
	}
	query := `INSERT INTO Route(sending, arrival, from_station_id, to_station_id, distance) VALUES
	(datetime(?), datetime(?), ?, ?, ?);`
	if _, err = DB.Exec(query, req.Sending, req.Arrival, ids[0], ids[1], req.Distance); err != nil {
		return err
	}
	return nil
}

// !-- Для админов --!
func GetAllUsers() ([]*customType.User, error) {
	rows, err := DB.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*customType.User
	for rows.Next() {
		var user customType.User
		err := rows.Scan(&user.ID, &user.Login, &user.Password, &user.Email,
			&user.Name, &user.Surname, &user.Patronymic)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func GetUser(id int) (customType.User, error) {
	var user customType.User
	err := DB.QueryRow("SELECT * FROM Users WHERE id = ?", id).Scan(&user.ID, &user.Login,
		&user.Password, &user.Email, &user.Name, &user.Surname, &user.Patronymic)
	if err != nil {
		return user, err
	}
	return user, nil
}
