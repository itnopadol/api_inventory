package ctrl

import (
	"fmt"
	"net/http"
	api "github.com/itnopadol/api_inventory/resp"
	md "github.com/itnopadol/api_inventory/model"
	"github.com/gin-gonic/gin"
	"log"
)

//=======================================================API App========================================================
func GetRequestList(c *gin.Context){

	log.Println("call GET Request List")
	c.Keys=headerKeys
	access_token := c.Request.URL.Query().Get("access_token")
	keyword :=c.Request.URL.Query().Get("keyword")

	fmt.Println("access_token = ",access_token)
	request := md.Request{}

	fmt.Println("call Label.GetRequestList :",keyword)

	rqs,err := request.GetByKeyWordRequest(keyword,dbc)
	if err != nil{
		fmt.Println("111")
		log.Println(err.Error())
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: "+err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		if rqs==nil{
			fmt.Println("2")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			c.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = rqs
			c.JSON(http.StatusOK, rs)
		}
	}
}

func GetPromotionTypeList(l *gin.Context){

	log.Println("call Get PromotionType List")
	l.Keys=headerKeys

	promotionType := md.Promotiontype{}

	pt,err := promotionType.GetPromotionType(dbc)
	if err != nil{
		fmt.Println("1")
		log.Println(err.Error())
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: "+err.Error()
		l.JSON(http.StatusNotFound, rs)
	} else {
		if pt==nil{
			fmt.Println("2")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			l.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = pt
			l.JSON(http.StatusOK, rs)
		}
	}
}

func GetPromotionMasterList(l *gin.Context){

	log.Println("call Get PromotionMaster List")
	l.Keys=headerKeys

	promotionmaster := md.Promotionmaster{}

	pm,err := promotionmaster.GetPromotionMaster(dbc)
	if err != nil{
		fmt.Println("1")
		log.Println(err.Error())
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: "+err.Error()
		l.JSON(http.StatusNotFound, rs)
	} else {
		if pm==nil{
			fmt.Println("2")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			l.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = pm
			l.JSON(http.StatusOK, rs)
		}
	}
}


func GetSectionManList(l *gin.Context){

	log.Println("call Get SectionMan List")
	l.Keys=headerKeys

	sectionman := md.SectionMan{}

	sm,err := sectionman.GetSectionMan(dbc)
	if err != nil{
		fmt.Println("1")
		log.Println(err.Error())
	}
	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: "+err.Error()
		l.JSON(http.StatusNotFound, rs)
	} else {
		if sm==nil{
			fmt.Println("2")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			l.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data =sm
			l.JSON(http.StatusOK, rs)
		}
	}
}