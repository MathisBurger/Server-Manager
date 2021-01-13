package main

import (
	"bytes"
	"encoding/json"
	"github.com/MathisBurger/Server-Manager/controller/monitoring"
	"github.com/MathisBurger/Server-Manager/static/hardware"
	"github.com/MathisBurger/Server-Manager/system"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
	"log"
	"time"
)

func main() {
	UpdateHardwareStatus()
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	// Static Calls
	app.Get("/static/getCPU-Info", monitoring.GetStaticCPUController)
	app.Get("/static/getMemory-Info", monitoring.GetStaticMemoryController)
	app.Get("/static/getStorage-Info", monitoring.GetStaticStorageController)
	app.Get("/static/updateHardware", monitoring.UpdateHardwareController)

	// Websocket Updates
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws/liveCPU", websocket.New(func(c *websocket.Conn) {
		var (
			err error
		)
		for {
			for _ = range time.Tick(1 * time.Second) {
				data := system.GetCPU_Stats()
				reqBodyBytes := new(bytes.Buffer)
				json.NewEncoder(reqBodyBytes).Encode(data)
				if err = c.WriteMessage(1, reqBodyBytes.Bytes()); err != nil {
					log.Println("write:", err)
				}
			}

		}

	}))
	app.Get("/ws/liveMemory", websocket.New(func(c *websocket.Conn) {
		var (
			err error
		)
		for {
			for _ = range time.Tick(1 * time.Second) {
				data := system.GetMemory_Stats()
				reqBodyBytes := new(bytes.Buffer)
				json.NewEncoder(reqBodyBytes).Encode(data)
				if err = c.WriteMessage(1, reqBodyBytes.Bytes()); err != nil {
					log.Println("write:", err)
				}
			}

		}

	}))

	app.Listen(":8080")
}

func UpdateHardwareStatus() {
	hardware.ReloadCPU()
	hardware.ReloadMemory()
	hardware.ReloadStorage()
}
