package ctrl

import (
	"fmt"
	"net/http"
	api "github.com/itnopadol/api_inventory/resp"
	lc "github.com/itnopadol/api_inventory/model"
	"github.com/gin-gonic/gin"
	"log"
)

//=======================================================API App========================================================
func GetLabelList(l *gin.Context){

	log.Println("call GET Label List")
	l.Keys=headerKeys
	access_token := l.Request.URL.Query().Get("access_token")
	keyword :=l.Request.URL.Query().Get("keyword")

	fmt.Println("access_token = ",access_token)
	label := lc.Label{}

	fmt.Println("call Label.GetLabelList :",keyword)

	ls,err := label.GetByUser(keyword,dbc)
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
		if ls==nil{
			fmt.Println("2")
			rs.Status = "error"
			rs.Message = "No Content: NotData"
			l.JSON(http.StatusNotFound, rs)
		}else {
			rs.Status = "success"
			rs.Data = ls
			l.JSON(http.StatusOK, rs)
		}
	}
}