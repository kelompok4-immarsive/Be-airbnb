package delivery

import (
	"fajar/testing/features/auth"
	"fajar/testing/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService auth.ServiceInterface
}

func New(service auth.ServiceInterface, e *echo.Echo) {
	handler := &AuthHandler{
		authService: service,
	}
	e.POST("/auth", handler.Login)
}

func (handler *AuthHandler) Login(c echo.Context) error {
	reqBody := UserRequest{}
	errBind := c.Bind(&reqBody)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("failed to bind data"))
	}

	token, err := handler.authService.Login(reqBody.Email, reqBody.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.PesanGagalHelper("failed to get token data"+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("login success", map[string]interface{}{
		"token": token,
	}))
}
