package factory

import (
	userDelivery "fajar/testing/features/user/delivery"
	userRepo "fajar/testing/features/user/repository"
	userService "fajar/testing/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepofaktory := userRepo.NewUser(db) //menginiasialisasi func new yang ada di repository
	userServiceFaktory := userService.NewUser(userRepofaktory)
	userDelivery.NewUser(userServiceFaktory, e)

}
