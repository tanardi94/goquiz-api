package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"ohas-api.com/v2/helpers"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authorized := strings.Split(
			c.Request().Header.Get("authorization"),
			"Bearer ",
		)

		if len(authorized) != 2 {
			return helpers.ResponseUnauthorized(c)
		}

		tokenString := authorized[1]

		claims := &helpers.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return helpers.JWTKey, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				return helpers.ResponseUnauthorized(c)
			case jwt.ValidationErrorExpired:
				response := "token expired"
				return helpers.ResponseMessage(c, false, http.StatusUnauthorized, response)
			default:
				return helpers.ResponseUnauthorized(c)
			}
		}

		if !token.Valid {
			return helpers.ResponseUnauthorized(c)
		}

		return next(c)
	}
}
