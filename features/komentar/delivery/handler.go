package delivery

import (
	"fajar/testing/features/komentar"
	"fajar/testing/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KomentarDelivery struct {
	komentarService komentar.ServiceInterface
}

func New(service komentar.ServiceInterface, e *echo.Echo) {
	Handler := &KomentarDelivery{
		komentarService: service,
	}

	e.POST("/komentars", Handler.Create)
	e.GET("/komentars", Handler.GetAll)
}
func (delivery *KomentarDelivery) Create(c echo.Context) error {
	komentarInput := KomentarRequest{}
	errBind := c.Bind(&komentarInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(komentarInput)
	err := delivery.komentarService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.PesanGagalHelper("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("success create data"))
}

func (delivery *KomentarDelivery) GetAll(c echo.Context) error {
	results, err := delivery.komentarService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("success read all komentar", dataResponse))
}
