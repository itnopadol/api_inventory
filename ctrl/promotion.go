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