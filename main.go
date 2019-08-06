package main

import (
	"fmt"

	"gameframework/core"

	"github.com/Unknwon/goconfig"
)

func main() {
	cfg, _ := goconfig.LoadConfigFile("./config/localdata.ini")
	value, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "path")
	fmt.Println(value)
	core.DriverStart()
}
