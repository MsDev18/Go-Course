package middleware

import (
	"E-01/entity"
	"E-01/pkg/claim"
	"E-01/pkg/errmsg"
	"E-01/service/authorizationservice"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AccessCheck(service authorizationservice.Service, permissions ...entity.PermissionTitle) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			claims := claim.GetClaimsFromEchoContext(c)
			isAllowed, err := service.CheckAccess(claims.UserID, claims.Role, permissions...)
			
			
			if err != nil {
				// TODO - log unexpected error
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": errmsg.ErrMsgSomethingWentWrong,
				})
			}

			if !isAllowed {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": errmsg.ErrMsgUserNotAllowd,
				})
			}

			return next(c)
		}
	}
}
