package main

import (
	"github.com/Floor-Gang/DikDik-Bot/internal"
	util "github.com/Floor-Gang/utilpkg"
)

func main() {
	//set config
	config := internal.GetConfig()
	internal.Start(config)

	util.KeepAlive()
}
