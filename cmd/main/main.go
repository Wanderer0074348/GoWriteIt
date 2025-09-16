package main

import (
	"github.com/Wanderer0074348/GoWriteIt/pkg/app"
	"github.com/Wanderer0074348/GoWriteIt/pkg/utils"
)

func main() {
	utils.ConnectDatabase()
	r := app.SetupRouter()
	r.Run(":8080")
}
