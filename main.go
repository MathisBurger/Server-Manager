package main

import (
	"github.com/MathisBurger/Server-Manager/controller/monitoring"
	"github.com/MathisBurger/Server-Manager/controller/monitoring/static"
	"github.com/MathisBurger/Server-Manager/static/hardware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	UpdateHardwareStatus()
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	// Static Calls
	app.Get("/static/getCPU-Info", static.GetStaticCPUController)
	app.Get("/static/getMemory-Info", static.GetStaticmemoryController)
	app.Get("/static/updateHardware", monitoring.UpdateHardwareController)

	app.Listen(":8080")
}

func UpdateHardwareStatus() {
	hardware.ReloadCPU()
	hardware.ReloadMemory()
}
