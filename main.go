package main

import (
 "database/sql"
 "fmt"
 "log"
 "net/http"
 "os"

 _ "github.com/lib/pq"
)

var db *sql.DB

func main() {
 // Получаем настройки подключения из переменных окружения
 dbHost := os.Getenv("DB_HOST")
 dbPort := os.Getenv("DB_PORT")
 dbUser := os.Getenv("DB_USER")
 dbPassword := os.Getenv("DB_PASSWORD")
 dbName := os.Getenv("DB_NAME")

 // Формируем строку подключения
 connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
  dbUser, dbPassword, dbName, dbHost, dbPort)

 // Открываем подключение к БД
 var err error
 db, err = sql.Open("postgres", connStr)
 if err != nil {
  log.Fatal(err)
 }
 defer db.Close()

 // Пингуем БД
 err = db.Ping()
 if err != nil {
  log.Fatal(err)
 }

 // Печатаем сообщение об успешном подключении
 fmt.Println("Successfully connected to the database!")

 // Маршруты
 http.HandleFunc("/", handleNameChange)

 // Запускаем сервер
 port := "6003"
 fmt.Printf("Server running on port %s\n", port)
 http.ListenAndServe(":"+port, nil)
}

// Обработчик изменения имени
func handleNameChange(w http.ResponseWriter, r *http.Request) {
 if r.Method == http.MethodPost {
  // Извлекаем имя из тела запроса
  name := r.FormValue("name")
  if name == "" {
   http.Error(w, "Name is required", http.StatusBadRequest)
   return
  }

  // Сохраняем имя в базе данных
  _, err := db.Exec("INSERT INTO users (name) VALUES ($1)", name)
  if err != nil {
   http.Error(w, "Failed to save name", http.StatusInternalServerError)
   log.Println("Error saving name:", err)
   return
  }

  fmt.Fprintf(w, "Name %s saved to database!", name)
 } else {
  http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
 }
}
