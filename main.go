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
	app := gin.New()
	app.Use(cors.Default())

	// pjc
	app.GET("/labels", ctrl.GetLabelList)
	app.POST("/label", ctrl.LabelSave)
	app.POST("/labelcancel", ctrl.LabelCancel)

	// gendocno
	app.GET("/gendocno", ctrl.GenDocno)

	//promotion
	app.POST("/promotion", ctrl.InsertAndUpdatePromotion)
	app.PUT("/promotioncancelitem", ctrl.PromotionCancel)
	app.PUT("/promotioncancel", ctrl.PromotionCancelItem)

	app.GET("/requests", ctrl.GetRequestList)
	app.GET("/promotiontype", ctrl.GetPromotionTypeList)
	app.GET("/promotionmaster", ctrl.GetPromotionMasterList)
	app.GET("/sectionman", ctrl.GetSectionManList)

	app.Run(":9010")
}