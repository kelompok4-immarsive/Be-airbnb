package mysql

import (
	"fajar/testing/config"

	gambar "fajar/testing/features/gambar/repository"
	komentar "fajar/testing/features/komentar/repository"

	reservation "fajar/testing/features/reservation/repository"
	room "fajar/testing/features/room/repository"
	user "fajar/testing/features/user/repository"

	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	MigrateDB(db)

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&room.Room{})
	db.AutoMigrate(&gambar.Gambar{})
	db.AutoMigrate(&komentar.Komentar{})
	db.AutoMigrate(&reservation.Reservation{})
}
