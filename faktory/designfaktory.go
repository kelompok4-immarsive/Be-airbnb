package factory

import (
	GambarDelivery "fajar/testing/features/gambar/delivery"
	GambarRepo "fajar/testing/features/gambar/repository"
	GambarService "fajar/testing/features/gambar/service"

	authDelivery "fajar/testing/features/auth/delivery"
	authRepo "fajar/testing/features/auth/repository"
	authService "fajar/testing/features/auth/service"
	komentarDelivery "fajar/testing/features/komentar/delivery"
	komentarRepo "fajar/testing/features/komentar/repository"
	komentarService "fajar/testing/features/komentar/service"

	roomDelivery "fajar/testing/features/room/delivery"
	roomRepo "fajar/testing/features/room/repository"
	roomService "fajar/testing/features/room/service"
	userDelivery "fajar/testing/features/user/delivery"
	userRepo "fajar/testing/features/user/repository"
	userService "fajar/testing/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	komentarRepoFactory := komentarRepo.New(db)
	komentarServiceFactory := komentarService.New(komentarRepoFactory)
	komentarDelivery.New(komentarServiceFactory, e)

	authRepoFactory := authRepo.NewAuth(db)
	authServiceFactory := authService.NewAuth(authRepoFactory)
	authDelivery.NewAuth(authServiceFactory, e)

	userRepofaktory := userRepo.NewUser(db) //menginiasialisasi func new yang ada di repository
	userServiceFaktory := userService.NewUser(userRepofaktory)
	userDelivery.NewUser(userServiceFaktory, e)

	RoomRepofaktory := roomRepo.NewRoom(db) //menginiasialisasi func new yang ada di repository
	RoomServiceFaktory := roomService.NewRoom(RoomRepofaktory)
	roomDelivery.NewRoom(RoomServiceFaktory, e)

	GambarRepofaktory := GambarRepo.NewGambar(db) //menginiasialisasi func new yang ada di repository
	GambarServiceFaktory := GambarService.NewGambar(GambarRepofaktory)
	GambarDelivery.NewGambar(GambarServiceFaktory, e)
}
