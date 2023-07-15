package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lathief/learn-fiber-go/pkg/utils"
)

type ReturnResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func HandleResponse(ctx *fiber.Ctx, message string, statusCode int) error {
	var res = ReturnResponse{
		StatusCode: statusCode,
		Message:    message,
	}
	utils.JoinString(GetStatusMsg(statusCode), message)
	return ctx.Status(statusCode).JSON(res)
}
func HandleResponseWithData(ctx *fiber.Ctx, data interface{}, message string, statusCode int) error {
	var res = ReturnResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	utils.JoinString(GetStatusMsg(statusCode), message)
	return ctx.Status(statusCode).JSON(res)
}
func HandleJustStatusCode(ctx *fiber.Ctx, statusCode int) error {
	return ctx.SendStatus(statusCode)
}
