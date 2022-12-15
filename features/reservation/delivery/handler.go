package delivery

import (
	"fajar/testing/features/reservation"
	"fajar/testing/utils/helper"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ReservationDeliv struct {
	ReservasiService reservation.ServiceInterface
}

func NewReservation(service reservation.ServiceInterface, e *echo.Echo) {
	handler := &ReservationDeliv{
		ReservasiService: service,
	}
	e.POST("/reservasi", handler.CreateReservasi)

}

func (delivery *ReservationDeliv) CreateReservasi(c echo.Context) error {

	// userIdtoken := middlewares.ExtractTokenUserId(c)

	Input := ReservationRequest{} //penangkapan data user reques dari entities user
	errbind := c.Bind(&Input)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data0"+errbind.Error()))
	}
	fmt.Println("start", Input.Start_date)
	res, errConvtime1 := time.Parse("02/01/2006", Input.Start_date)
	if errConvtime1 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data1"))
	}
	res1, errConvtime2 := time.Parse("02/01/2006", Input.End_date)
	if errConvtime2 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data2"))
	}

	log.Println(Input.UserID)

	dataCore := RequestToCore(Input, res, res1) //data mapping yang diminta create\
	// dataCore.UserID = uint(userIdtoken)
	errResultCore := delivery.ReservasiService.CreateReservasi(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data3"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success Create Reservation "))
}
