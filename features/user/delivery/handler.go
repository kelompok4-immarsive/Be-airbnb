package delivery

import (
	"fajar/testing/features/user"
	"fajar/testing/helper"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserDeliv struct {
	ServiceUser user.ServiceEntities
}

func NewUser(Service user.ServiceEntities, e *echo.Echo) {
	Handler := &UserDeliv{
		ServiceUser: Service,
	}
	e.POST("/user/register", Handler.Register)
	e.GET("/user", Handler.Getall)
}
func (delivery *UserDeliv) Register(c echo.Context) error {
	Inputuser := RequestUser{} //request pengisian postman atau request kontrak user
	generate, _ := bcrypt.GenerateFromPassword([]byte(Inputuser.Password), bcrypt.DefaultCost)
	Inputuser.Password = string(generate)
	errbind := c.Bind(&Inputuser)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errbind.Error()))
	}
	dataCore := RequestUserToCoreUser(Inputuser) //dari request mengirim ke coreuser
	errResultCore := delivery.ServiceUser.Register(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("berhasil melakukan register"))
}
func (delivery *UserDeliv) Getall(c echo.Context) error {

	result, err := delivery.ServiceUser.Getall() //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"))
	}
	var ResponData = ListUserCoreToUserRespon(result)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("berhasil membaca  user", ResponData))

}
