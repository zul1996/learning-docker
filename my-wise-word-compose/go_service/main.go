package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var dbpool *pgxpool.Pool

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

	// Membuat pool koneksi
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	dbpool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Tidak dapat membuat pool koneksi: %v\n", err)
	}
	defer dbpool.Close()

	// Membuat aplikasi Fiber
	app := fiber.New()

	// Rute untuk "/"
	app.Get("/", func(c fiber.Ctx) error {
		fmt.Println("ouch!")
		return c.SendString("The true measure of a shinobi is not how he lives, but how he dies. - Jiraiya👋!")
	})

	// Rute untuk "/ping"
	app.Get("/ping", func(c fiber.Ctx) error {
		
		if err != nil {
			log.Printf("Gagal mendapatkan koneksi: %v\n", err)
			return c.SendString("Connection to database failed")
		}

		return c.SendString("Connection to database successful")
	})

	// Memulai server
	app.Listen(":78")
}
