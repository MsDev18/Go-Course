package backofficeuserhandler

import (
	"E-01/delivery/httpserver/middleware"
	"E-01/entity"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/backoffice/users")
	userGroup.GET("/", h.listUsers, middleware.Auth(h.authSvc, h.authConfig), middleware.AccessCheck(h.authorizationSvc, entity.UserListPermissions))
}
