package delivery

import (
	"fajar/testing/features/room"
	"fajar/testing/utils/helper"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoomDeliv struct {
	RoomService room.ServiceInterface
}

func NewRoom(service room.ServiceInterface, e *echo.Echo) {
	handler := &RoomDeliv{
		RoomService: service,
	}
	e.POST("/room", handler.CreateRoom)
}

func (delivery *RoomDeliv) CreateRoom(c echo.Context) error {

	room := RoomRequest{}
	errBind := c.Bind(&room)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}
	fmt.Println("port input", room)
	result := room.reqToCore()

	err := delivery.RoomService.CreateRoom(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error insert into database"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Add Room"))

}
