package static

import (
	"encoding/json"
	"github.com/MathisBurger/Server-Manager/system"
	"github.com/gofiber/fiber/v2"
	"os"
)

func GetStaticStorageController(c *fiber.Ctx) error {
	f, err := os.Open("./static/hardware/Storage.json")
	if err != nil {
		return err
	}
	ans := new(system.StorageInfo)
	err = json.NewDecoder(f).Decode(ans)
	return c.JSON(ans)
}
