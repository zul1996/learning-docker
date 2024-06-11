package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	// Membuat koneksi database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v\n", err)
	}
	defer db.Close()

	// Membuat aplikasi Fiber
	app := fiber.New()

	// Rute untuk "/"
	app.Get("/", func(c fiber.Ctx) error {

		_, err := http.Get("http://localhost:4331")

		if err != nil {
			log.Println("Ping failed", err)
			return c.SendString("Allah tidak membebani seseorang melainkan sesuai dengan kesanggupannya. (Q.S Al Baqarah: 286")



		}
		log.Printf("Ping succesful")
		return c.SendString("Pong!")
	})

	// Rute untuk "/ping"
	app.Get("/ping", func(c fiber.Ctx) error {

		// Menyisipkan data timestamp saat ini ke dalam kolom timestamp
		_, err = db.Exec("INSERT INTO timestamp (timestamp) VALUES ($1)", time.Now())
		if err != nil {
            log.Printf("Ping failed :%v\n", err)
			return c.SendString("Ping failed")
		}

		// Mengirimkan respons jika semua operasi berhasil
        log.Printf("Ping successful")
		return c.SendString("Pong")
	})

	// Memulai server
	log.Fatal(app.Listen(":77"))
}
