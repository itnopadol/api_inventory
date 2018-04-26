package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"strconv"
)

type RequestDocno struct {
	LastDocno int `json:"last_number" db:"LastDocno"`
	DocYear string `json:"doc_year" db:"DocYear"`
	DocMonth string `json:"doc_month" db:"DocMonth"`
	RequestDocno string `json:"Request_docno"`
}


func (r *RequestDocno)GenDocno(keyword string,db *sqlx.DB) error {
	
if keyword == "RQ"{
	
	var last_number string
	var doc_year string
	var doc_month string
	var snumber string

	sql := `select substring(cast(dbo.FT_CG_ThaiYear(getdate()) as varchar(4)),3,2) as DocYear
	,case when len(cast(month(getdate()) as varchar(2))) = 1 then '0' + rtrim(cast(month(getdate()) as varchar(2))) else rtrim(cast(month(getdate()) as varchar(2))) end DocMonth
	,isnull((select Count(*) + 1 from NpMaster.dbo.TB_PM_Request where year(Docdate) = year(getdate()) and  month(DocDate) = month(getdate()) ),1) as LastDocno`
	err := db.Get(r,sql)
	if err != nil {
		return err
	}
	last_number = strconv.Itoa(r.LastDocno)
	doc_year=(r.DocYear)
	doc_month=(r.DocMonth)
	fmt.Println("Last No = ", last_number)

	if(len(string(last_number))==1){
		snumber = "000"+last_number
	}
	if(len(string(last_number))==2){
		snumber = "00"+last_number
	}
	if(len(string(last_number))==3){
		snumber = "0"+last_number
	}
	if(len(string(last_number))==4) {
		snumber = last_number
	}
	fmt.Println("Number = ",snumber)
	r.RequestDocno = (keyword+doc_year+doc_month+"-"+snumber)
	fmt.Println("RequestDocno = ",r.RequestDocno)

}else {
		fmt.Println("keyword = ",keyword)
	}
	return nil
}