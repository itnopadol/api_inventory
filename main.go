package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itnopadol/api_inventory/ctrl"
	"gopkg.in/gin-contrib/cors.v1"

)

func main() {
	fmt.Println("BC API Inventory")
	// 1 = MsSql server , 0 = MySql
	app := gin.Default()
	app.Use(cors.Default())

	// pjc
	app.GET("/labels", ctrl.GetLabelList)

	app.Run(":9010")
}