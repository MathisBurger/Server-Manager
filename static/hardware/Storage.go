package hardware

import (
	"encoding/json"
	"github.com/MathisBurger/Server-Manager/system"
	"io/ioutil"
)

func ReloadStorage() {
	info := system.GetStorageInfo()
	data, _ := json.MarshalIndent(info, "", " ")
	_ = ioutil.WriteFile("./static/hardware/Storage.json", data, 0664)
}
