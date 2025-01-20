package main

import (
    "fmt"
    "net/http"
    "log"
    "database/sql"
    "github.com/lib/pq"
)

// Глобальная переменная для работы с базой данных
var db *sql.DB

// Функция для подключения к базе данных
func initDB() {
    var err error
    connStr := "user=user password=password dbname=mydb sslmode=disable host=postgres port=5432"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }
}

// Функция для записи фразы в базу данных
func savePhraseToDB(phrase string) {
    _, err := db.Exec("INSERT INTO phrases (content) VALUES ($1)", phrase)
    if err != nil {
        log.Println("Error saving phrase to DB: ", err)
    }
}

// Функция для обработки запросов
func handler(w http.ResponseWriter, r *http.Request) {
    phrase := "Hello, Vladislav!"
    
    // Если запрос POST, то сохраняем фразу в БД
    if r.Method == http.MethodPost {
        phrase = r.FormValue("phrase")
        savePhraseToDB(phrase)
    }
    
    // Выводим фразу на страницу
    fmt.Fprintf(w, "Current phrase: %s", phrase)
}

func main() {
    initDB()
    defer db.Close()
    
    // Устанавливаем обработчик маршрута
    http.HandleFunc("/", handler)
    
    port := "6003"
    fmt.Printf("Server running on port %s\n", port)
    http.ListenAndServe(":"+port, nil)
}
