package hardware

import (
	"encoding/json"
	"github.com/MathisBurger/Server-Manager/system"
	"io/ioutil"
)

func ReloadMemory() {
	info := system.GetMemoryInfo()
	data, _ := json.MarshalIndent(info, "", " ")
	_ = ioutil.WriteFile("./static/hardware/Memory.json", data, 0664)
}
