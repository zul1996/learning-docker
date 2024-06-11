package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
)


func main() {
	// Membaca variabel lingkungan
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Membuat string koneksi database
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Println(connStr)

	db, err := sql.Open("postgres",connStr)
	if err != nil {
		fmt.Println("Connect db failed")
	}

	// Membuat aplikasi Fiber
	app := fiber.New()

	// Rute untuk "/"
	app.Get("/", func(c fiber.Ctx) error {
		fmt.Println("ouch!")
		return c.SendString("The true measure of a shinobi is not how he lives, but how he dies. - Jiraiya👋!")
	})

	// Rute untuk "/ping"
	app.Get("/ping", func(c fiber.Ctx) error {
		if err := db.Ping(); err != nil {
			log.Printf("Gagal mendapatkan koneksi: %v\n", err)
			return c.SendString("Connection to database failed")
		}

		return c.SendString("Connection to database successful")
	})

	// Memulai server
	app.Listen(":78")
}
