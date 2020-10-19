package main

import (
	"github.com/mniudanri/store/config"
	"github.com/mniudanri/store/router"
)

func main() {
	setter := config.GetConfig()
  defer setter.DB.Close()

	router.InitRoute(":5100")
}
