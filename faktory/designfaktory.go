package factory

import (
	authDelivery "fajar/testing/features/auth/delivery"
	authRepo "fajar/testing/features/auth/repository"
	authService "fajar/testing/features/auth/service"
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
	authRepoFactory := authRepo.NewAuth(db)
	authServiceFactory := authService.NewAuth(authRepoFactory)
	authDelivery.NewAuth(authServiceFactory, e)

	userRepofaktory := userRepo.NewUser(db) //menginiasialisasi func new yang ada di repository
	userServiceFaktory := userService.NewUser(userRepofaktory)
	userDelivery.NewUser(userServiceFaktory, e)

	RoomRepofaktory := roomRepo.NewRoom(db) //menginiasialisasi func new yang ada di repository
	RoomServiceFaktory := roomService.NewRoom(RoomRepofaktory)
	roomDelivery.NewRoom(RoomServiceFaktory, e)
}
