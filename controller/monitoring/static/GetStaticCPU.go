package static

import (
	"encoding/json"
	"github.com/MathisBurger/Server-Manager/system"
	"github.com/gofiber/fiber/v2"
	"os"
)

func GetStaticCPUController(c *fiber.Ctx) error {
	f, err := os.Open("./static/hardware/CPU.json")
	if err != nil {
		return err
	}
	ans := new(system.CPU_Info)
	err = json.NewDecoder(f).Decode(ans)
	return c.JSON(ans)
}
