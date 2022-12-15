package delivery

import (
	"fajar/testing/features/room"
	"fajar/testing/utils/helper"
	"fmt"
	"net/http"
	"strconv"

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
	e.GET("/room", handler.GetAll)
	e.PUT("/room/:id", handler.Updates)
	e.DELETE("/room/:id", handler.Deletes)
	e.GET("/room/:id/user", handler.GetAllRoomUser)

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
func (delivery *RoomDeliv) GetAll(c echo.Context) error {

	result, err := delivery.RoomService.GetAll() //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListRoomCoreToRoomRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca  Room", ResponData))

}
func (delivery *RoomDeliv) Updates(c echo.Context) error {
	room := RoomRequest{}
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}

	errBind := c.Bind(&room)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	result := room.reqToCore()
	err := delivery.RoomService.UpdateRoom(id, result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Update Failed"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Update"))

}
func (delivery *RoomDeliv) Deletes(c echo.Context) error {
	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.FailedResponse("tidak bisa diakses khusus admin!!!"))
	// }

	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}
	data, err := delivery.RoomService.DeleteRoom(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Delete rooms"))
	}

	result := RoomCoreToRoomRespon(data)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Delete rooms", result))

}
func (delivery *RoomDeliv) GetAllRoomUser(c echo.Context) error {
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}

	userId, err := delivery.RoomService.GetUserRoom(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Id not Found"))
	}

	// result := coreToResponse(userId)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get mentee", userId))

}
