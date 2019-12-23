package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/add", addHandler)
	e.POST("/mul", mulHandler)
	e.Start(":8080")
}

type addRequest struct {
	A int
	B int
}

type addResponse struct {
	Result int `json:"res"`
}

func addHandler(c echo.Context) error {
	// parse request
	var reqBody addRequest
	err := c.Bind(&reqBody)
	if err != nil {
		return err
	}

	// calculate
	result := reqBody.A + reqBody.B

	// send response
	return c.JSON(200, addResponse{
		Result: result,
	})
}

type mulRequest struct {
	A int
	B int
}

type mulResponse struct {
	Result int `json:"res"`
}

func mulHandler(c echo.Context) error {
	// parse request
	var reqBody mulRequest
	err := c.Bind(&reqBody)
	if err != nil {
		return err
	}

	// calculate
	result := reqBody.A * reqBody.B

	// send response
	return c.JSON(200, mulResponse{
		Result: result,
	})
}
