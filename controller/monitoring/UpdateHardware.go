package monitoring

import (
	"github.com/MathisBurger/Server-Manager/static/hardware"
	"github.com/gofiber/fiber/v2"
)

type UpdateHardwareResponse struct {
	Message    string `json:"message"`
	Alert      string `json:"alert"`
	HttpStatus int    `json:"htpp_status"`
	Status     string `json:"status"`
}

func UpdateHardwareController(c *fiber.Ctx) error {
	hardware.ReloadCPU()
	hardware.ReloadMemory()
	return c.JSON(UpdateHardwareResponse{
		"Successfully updated static hardware",
		"alert alert-success",
		200,
		"ok",
	})
}
