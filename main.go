package main

import (
	"github.com/marcelwolf1989/docker-machine-driver-aliyunecs/aliyunecs"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(aliyunecs.NewDriver("", ""))
}
