package ctrl

import (
	"fmt"
	"net/http"
	api "github.com/itnopadol/api_inventory/resp"
	m "github.com/itnopadol/api_inventory/model"
	"github.com/gin-gonic/gin"
	"log"
)

func GenDocno(c *gin.Context){
	log.Println("call Get GenDocno")
	c.Keys = headerKeys
	keyword :=c.Request.URL.Query().Get("keyword")

	r := new(m.RequestDocno)
	fmt.Println("keyword = ",keyword)
	err := r.GenDocno(keyword,dbc)
	if err != nil {
		fmt.Println(err)
	}

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = r
		c.JSON(http.StatusOK, rs)
	}

}