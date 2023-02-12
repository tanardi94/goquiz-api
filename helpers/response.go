package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Output struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func CobaResponse(c echo.Context, data interface{}, errors interface{}) error {

	response := Output{
		Success: true,
		Message: "Request is successfully processed",
		Errors:  errors,
		Data:    data,
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(response)
}

func ResponseSuccess(c echo.Context, data interface{}) error {

	response := Output{
		Success: true,
		Message: "Request is successfully processed",
		Errors:  make(map[string]string, 0),
		Data:    data,
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(response)
}

func ResponseCreated(c echo.Context, item string) error {

	response := Output{
		Success: true,
		Message: item + " is successfully created",
		Errors:  make(map[string]string, 0),
		Data:    make(map[string]string, 0),
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(response)

}

func ResponseFailure(c echo.Context, message string) error {

	response := Output{
		Success: false,
		Message: "Something went wrong",
		Errors: map[string]string{
			"error_message": message,
		},
		Data: make(map[string]string, 0),
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)
}

func ResponseNotFound(c echo.Context, message string) error {

	response := Output{
		Success: false,
		Message: message + " is not found",
		Errors:  make(map[string]string, 0),
		Data:    make(map[string]string, 0),
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusNotFound)
	return json.NewEncoder(c.Response()).Encode(response)

}

func ResponseServerError(c echo.Context, message string) error {
	response := Output{
		Success: false,
		Message: "Something went wrong from our side",
		Errors:  make(map[string]string, 0),
		Data:    make(map[string]string, 0),
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(c.Response()).Encode(response)
}

func ResponseForbidden(c echo.Context) error {
	response := Output{
		Success: false,
		Message: "You don't have access to this request",
		Errors:  make(map[string]string, 0),
		Data:    make(map[string]string, 0),
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusForbidden)
	return json.NewEncoder(c.Response()).Encode(response)

}

func ResponseUnauthorized(c echo.Context) error {

	response := Output{
		Success: false,
		Message: "Invalid Authorization, Please Log in to access",
		Errors:  make(map[string]string, 0),
		Data:    make(map[string]string, 0),
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusUnauthorized)
	return json.NewEncoder(c.Response()).Encode(response)

}

func ResponseCustom(c echo.Context, status bool, message string, code int, data interface{}, errors interface{}) error {
	response := Output{
		Success: status,
		Message: message,
		Errors:  errors,
		Data:    data,
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(code)
	return json.NewEncoder(c.Response()).Encode(response)
}

func ResponseMessage(c echo.Context, status bool, code int, message string) error {
	response := Output{
		Success: status,
		Message: message,
		Errors:  make(map[string]string, 0),
		Data:    make(map[string]string, 0),
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(code)
	return json.NewEncoder(c.Response()).Encode(response)
}
