package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// RootResponse type
type RootResponse struct {
	Status     string     `json:"status"`
	Response   Response   `json:"response"`
	ServerInfo ServerInfo `json:"serverInfo"`
}

// Response type
type Response struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// ServerInfo type
type ServerInfo struct {
	ServerTime string `json:"serverTime"`
}

const (
	ERROR   = "ERROR"
	SUCCESS = "OK"
)

// Result response
func Result(httpStatus int, status string, data interface{}, msg string, c echo.Context) error {

	respObj := RootResponse{
		Status: status,
		Response: Response{
			Message: msg,
			Result:  data,
		},
		ServerInfo: ServerInfo{
			ServerTime: time.Now().Format(time.RFC3339Nano),
		},
	}

	resp, err := json.Marshal(respObj)
	if err != nil {
		return err
	}

	return c.JSONBlob(httpStatus, resp)
}

// Ok response
func Ok(c echo.Context) error {
	return Result(http.StatusOK, SUCCESS, map[string]interface{}{}, "Operation succeeded", c)
}

// OKWithMesssage response
func OKWithMesssage(msg string, c echo.Context) error {
	return Result(http.StatusOK, SUCCESS, map[string]interface{}{}, msg, c)
}

// OkWithData response
func OkWithData(data interface{}, c echo.Context) error {
	return Result(http.StatusOK, SUCCESS, data, "Operation succeeded", c)
}

// CreatedWithData response
func CreatedWithData(data interface{}, c echo.Context) error {
	return Result(http.StatusCreated, SUCCESS, data, "Operation succeeded", c)
}

// OkDetailed response
func OkDetailed(data interface{}, message string, c echo.Context) error {
	return Result(http.StatusOK, SUCCESS, data, message, c)
}

// Fail response
func Fail(c echo.Context) error {
	return Result(http.StatusOK, ERROR, map[string]interface{}{}, "Operation failed", c)
}

// FailWithMessage response
func FailWithMessage(message string, c echo.Context) error {
	return Result(http.StatusOK, ERROR, map[string]interface{}{}, message, c)
}

// FailWithDetailed response
func FailWithDetailed(status string, data interface{}, message string, c echo.Context) error {
	return Result(http.StatusOK, status, data, message, c)
}

// FailWithDataMessage response
func FailWithDataMessage(data interface{}, message string, c echo.Context) error {
	return Result(http.StatusOK, ERROR, data, message, c)
}

// FailDetailedwithCode response
func FailDetailedwithCode(code int, data interface{}, message string, c echo.Context) error {
	return Result(code, ERROR, data, message, c)
}

// FailWithMessageWithCode response
func FailWithMessageWithCode(code int, message string, c echo.Context) error {
	return Result(code, ERROR, map[string]interface{}{}, message, c)
}
