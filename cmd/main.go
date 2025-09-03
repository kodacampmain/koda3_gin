package main

import (
	"log"

	"github.com/kodacampmain/koda3_gin/internal/configs"
	"github.com/kodacampmain/koda3_gin/internal/routers"

	// "github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// manual load env
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("Failed to load env\nCause: ", err.Error())
	// 	return
	// }
	// log.Println(os.Getenv("DBUSER"))

	// inisialisasi db
	// psql string: postgres://username:password@host:port/namadb
	db, err := configs.InitDB()
	if err != nil {
		log.Println("Failed to connect to database\nCause: ", err.Error())
		return
	}
	defer db.Close()
	// testing koneksi db

	if err := configs.TestDB(db); err != nil {
		log.Println("Ping to DB failed\nCause: ", err.Error())
		return
	}
	log.Println("DB Connected")
	// inisialisasi engine gin
	router := routers.InitRouter(db)
	// client => (router => handler => repo => handler) => client
	// jalankan engine gin
	router.Run("localhost:3000")
}
