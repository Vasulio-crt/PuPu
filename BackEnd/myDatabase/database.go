package myDatabase

import (
	"database/sql"
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
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE login = ? OR email = ?)"
	err := DB.QueryRow(query, login, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// !-- Для админов --!
func GetAllUsers() ([]*customType.User, error) {
	rows, err := DB.Query("SELECT * FROM users")
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
	err := DB.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Login,
		&user.Password, &user.Email, &user.Name, &user.Surname, &user.Patronymic)
	if err != nil {
		return user, err
	}
	return user, nil
}
