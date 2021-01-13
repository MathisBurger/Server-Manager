package hardware

import (
	"encoding/json"
	"github.com/MathisBurger/Server-Manager/system"
	"io/ioutil"
)

func ReloadCPU() {
	info := system.GetCPU_Info()
	data, _ := json.MarshalIndent(info, "", " ")
	_ = ioutil.WriteFile("./static/hardware/CPU.json", data, 0664)
}
