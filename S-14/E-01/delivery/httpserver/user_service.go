package httpserver

import (
	"E-01/service/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) userLogin(c echo.Context) error {

	var req user.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err, code := s.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(int(code), err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (s Server) userRegister(c echo.Context) error {

	var req user.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := s.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (s Server) userProfile(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	
	claims, err := s.authSvc.ParseToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := s.userSvc.Profile(user.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
