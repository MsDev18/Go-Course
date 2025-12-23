package middleware

import (
	"E-01/param"
	"E-01/pkg/claim"
	"E-01/pkg/errmsg"
	"E-01/pkg/timestamp"
	"E-01/service/presenceservice"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpsertPresence(service presenceservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			claims := claim.GetClaimsFromEchoContext(c)
			_, err = service.Upsert(c.Request().Context(), param.UpsertPrecenseRequest{
				UserID: claims.UserID,
				Timestamp: 	timestamp.Now(),
			})

			if err != nil {
				// TODO - log unexpected error
				// we can just log the error and go to the next step (middleware, handler)  
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message" : errmsg.ErrMsgSomethingWentWrong,
				})
			}

			return next(c)
		}
	}
}
