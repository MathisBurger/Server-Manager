package monitoring

import (
	"encoding/json"
	"github.com/MathisBurger/Server-Manager/system"
	"github.com/gofiber/fiber/v2"
	"os"
)

func GetStaticMemoryController(c *fiber.Ctx) error {
	f, err := os.Open("./static/hardware/Memory.json")
	if err != nil {
		return err
	}
	ans := new(system.MemoryInfo)
	err = json.NewDecoder(f).Decode(ans)
	return c.JSON(ans)
}
