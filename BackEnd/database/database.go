package database

import (
	"database/sql"
	"log"
	"os"
	"sellTrainTicket/handlers"

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
	defer DB.Close()

	if !dbExists {
		log.Printf("Файл базы данных '%s' не найден. Создание таблиц...", dbName)
		log.Println("Выполнение скрипта create_db.sql...")
		sqlScript, err := os.ReadFile("database/create_db.sql")
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

func AddUserDB(user handlers.User) {

}
