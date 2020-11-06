package handler

import (
	"app/pkg/db/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// BodyDumper output logs
func BodyDumper(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

// RootHandler return response like hello world
func RootHandler(c echo.Context) error {
	res := model.Response{
		Message: "hello echo framework",
		Status:  http.StatusOK,
	}
	bytes, _ := json.Marshal(res)

	return c.JSONBlob(http.StatusOK, bytes)
}

// TestHandler return map response
func TestHandler(c echo.Context) error {
	response := map[string]string{"ping": "pong"}

	return c.JSON(http.StatusOK, response)
}
