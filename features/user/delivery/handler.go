package delivery

import (
	"fajar/testing/features/user"
	"fajar/testing/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
	e.PUT("/user/:id", Handler.Update)
	e.GET("/user/:id", Handler.GetById)
	e.DELETE("/user/:id", Handler.DeleteById)
}
func (delivery *UserDeliv) Register(c echo.Context) error {
	Inputuser := RequestUser{} //request pengisian postman atau request kontrak user
	errbind := c.Bind(&Inputuser)

	generatePass := user.Bcript(Inputuser.Password)
	Inputuser.Password = generatePass

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errbind.Error()))
	}
	dataCore := RequestUserToCoreUser(Inputuser) //dari request mengirim ke coreuser
	errResultCore := delivery.ServiceUser.Register(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("berhasil create user"))
}
func (delivery *UserDeliv) Getall(c echo.Context) error {

	result, err := delivery.ServiceUser.Getall() //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListUserCoreToUserRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca  user", ResponData))

}
func (delivery *UserDeliv) Update(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	userInput := RequestUser{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := RequestUserToCoreUser(userInput)
	err := delivery.ServiceUser.Update(id, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success Update data"))
}
func (delivery *UserDeliv) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := delivery.ServiceUser.GetById(id) //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = UserCoreToUserRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca  user", ResponData))
}
func (delivery *UserDeliv) DeleteById(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	del, err := delivery.ServiceUser.DeleteId(id) //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr Hapus data"))
	}
	result := UserCoreToUserRespon(del)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil menghapus user", result))
}
