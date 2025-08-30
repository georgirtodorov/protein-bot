package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type App struct {
	db *sql.DB
}

func (a *App) routes() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/add", a.add)
}

func (a *App) add(w http.ResponseWriter, r *http.Request) {

	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// name := r.URL.Query().Get("name")
	// if name == "" {
	// 	http.Error(w, "Missing 'name'", http.StatusBadRequest)
	// 	return
	// }

	// _, err := a.db.Exec("INSERT INTO users (name) VALUES ($1)", name)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Insert failed: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	// fmt.Fprintf(w, "User %s added successfully!", name)

	fmt.Fprintf(w, "User pesho added successfully!")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// Connect to DB
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Setup routes
	app := &App{db: db}
	app.routes()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("server failed to start, PORT environment variable not set")
	}
	addr := ":" + port
	log.Printf("Starting server on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func connectDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("✅ Connected to database successfully")
	return db, nil
}
