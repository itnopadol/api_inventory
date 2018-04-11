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

func LabelSave(c *gin.Context){
	log.Println("call Post LabelSave")
	c.Keys=headerKeys

	fmt.Println("Ctrl.LabelSave")
	//pjc := &model.ProjectCard{}
	lc := lc.InsertLabel{}
	rs := api.Response{}
		if err := c.BindJSON(&lc); err != nil{
			fmt.Println(lc)
			log.Println("Error Decode(&lc) >>", err)
			rs.Status = "fail"
			rs.Message = err.Error()
			c.JSON(http.StatusOK,rs)

		}else{
			fmt.Println("ID = ",lc.JobID)
			//fmt.Println("Check Status = ", pjc.CheckExists())
			if lc.CheckExists(dbc,lc.ItemCode,lc.BarCode,lc.UnitCode,lc.LabelType,lc.CreatorCode) != 0 {
				//  มีรายการแล้ว
				updateProject,err := lc.Update(dbc)
				fmt.Println("<---------------update1")
				fmt.Println(lc.JobID)
				if err != nil {
					fmt.Println("Error Update DB:", err)
					rs.Status = "fail"
					rs.Message = "Error Update Label :"+err.Error()
					c.JSON(http.StatusBadRequest,rs)
					return
				}
					rs.Status = "success"
					rs.Data = updateProject
			}else{
				newProject,err := lc.Insert(dbc)
				fmt.Println("<---------------Start insert Label")
				fmt.Println(lc.JobID)
				if err != nil {
					fmt.Println("Error Insert DB:", err)
					rs.Status = "fail"
					rs.Message = "Error Insert Label :"+err.Error()
					c.JSON(http.StatusBadRequest,rs)
					return
			}
				rs.Status = "success"
				rs.Data = newProject
			}			
		}		
	rs.Status="success"
	c.JSON(http.StatusOK,rs)
}

func LabelCancel(c *gin.Context){
	log.Println("call Post LabelCancel")
	c.Keys=headerKeys

	fmt.Println("Ctrl.LabelCancel")
	//pjc := &model.ProjectCard{}
	lc := lc.InsertLabel{}
	rs := api.Response{}
		if err := c.BindJSON(&lc); err != nil{
			fmt.Println(lc)
			log.Println("Error Decode(&lc) >>", err)
			rs.Status = "fail"
			rs.Message = err.Error()
			c.JSON(http.StatusOK,rs)

		}else{
			fmt.Println("ID = ",lc.ItemCode)
			//fmt.Println("Check Status = ", pjc.CheckExists())
			if lc.CheckExists(dbc,lc.ItemCode,lc.BarCode,lc.UnitCode,lc.LabelType,lc.CreatorCode) != 0 {
				//  มีรายการแล้ว
				LabelCancel,err := lc.Cancel(dbc)
				fmt.Println("<---------------Cancel")
				fmt.Println(lc.ItemCode)
				if err != nil {
					fmt.Println("Error Cancel DB:", err)
					rs.Status = "fail"
					rs.Message = "Error Cancel Label :"+err.Error()
					c.JSON(http.StatusBadRequest,rs)
					return
				}
					rs.Status = "success"
					rs.Data = LabelCancel
			}else{
				newProject,err := lc.Insert(dbc)
				fmt.Println("<---------------Start insert Label")
				fmt.Println(lc.JobID)
				if err != nil {
					fmt.Println("Error Insert DB:", err)
					rs.Status = "fail"
					rs.Message = "Error Insert Label :"+err.Error()
					c.JSON(http.StatusBadRequest,rs)
					return
			}
				rs.Status = "success"
				rs.Data = newProject
			}			
		}		
	rs.Status="success"
	c.JSON(http.StatusOK,rs)
}