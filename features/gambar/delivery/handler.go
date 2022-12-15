package delivery

import (
	"errors"
	"fajar/testing/features/gambar"
	"fajar/testing/utils/helper"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type GambarDeliv struct {
	GambarService gambar.ServiceInterface
}

func NewGambar(service gambar.ServiceInterface, e *echo.Echo) {
	handler := &GambarDeliv{
		GambarService: service,
	}
	e.POST("/gambar", handler.CreateGambar)
	e.GET("/gambar", handler.Getall)

}

func (delivery *GambarDeliv) CreateGambar(c echo.Context) error {

	gambar := GambarRequest{}
	roomid := c.FormValue("room_id")
	id, _ := strconv.Atoi(roomid)
	gambar.RoomID = uint(id)
	file, _ := c.FormFile("file")
	if file != nil {
		res, err := helper.UploadProfile(c)
		if err != nil {
			return errors.New("Create gambar Failed. Cannot Upload Data.")
		}
		gambar.Image_url = res
	} else {
		gambar.Image_url = "https://project3bucker.s3.ap-southeast-1.amazonaws.com/dummy-profile-pic.png"
	}
	result := gambar.reqToCore()

	err := delivery.GambarService.CreateGambar(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error insert into database"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Add Gambar"))

}
func (delivery *GambarDeliv) Getall(c echo.Context) error {

	result, err := delivery.GambarService.Getall() //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListGambarCoreToGambarRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca  user", ResponData))

}
