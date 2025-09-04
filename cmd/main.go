package main

import (
	"fmt"

	"github.com/kodacampmain/koda3_gin/pkg"

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
	// db, err := configs.InitDB()
	// if err != nil {
	// 	log.Println("Failed to connect to database\nCause: ", err.Error())
	// 	return
	// }
	// defer db.Close()
	// testing koneksi db

	// if err := configs.TestDB(db); err != nil {
	// 	log.Println("Ping to DB failed\nCause: ", err.Error())
	// 	return
	// }
	// log.Println("DB Connected")
	// inisialisasi engine gin
	// router := routers.InitRouter(db)
	// client => (router => handler => repo => handler) => client
	// jalankan engine gin
	// router.Run("localhost:3000")
	hc := pkg.NewHashConfig()
	hc.UseRecommended()

	password := "koda"
	hash, err := hc.GenHash(password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	isMatch, err := hc.CompareHashAndPassword(password, hash)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(isMatch)
	// hash, err = hc.GenHash(password)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(hash)
}
