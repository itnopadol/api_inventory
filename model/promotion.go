package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//"golang.org/x/tools/go/gcimporter15/testdata"
)

type Request struct{
	DocNo string `json:"DocNo" db:"DocNo"`
	DocDate string `json:"DocDate" db:"DocDate"`
	SecMan string `json:"SecMan" db:"SecMan"`
	IsConfirm float64 `json:"IsConfirm" db:"IsConfirm"`
	IsCancel float64 `json:"IsCancel" db:"IsCancel"`
	PMCode string `json:"PMCode" db:"PMCode"`
	CreatorCode string `json:"CreatorCode,omitempty" db:"CreatorCode"`
	CreateDate string `json:"CreateDate,omitempty" db:"CreateDate"`
	EditorCode string `json:"EditorCode,omitempty" db:"EditorCode"`
	EditDate string `json:"EditDate,omitempty" db:"EditDate"`
	IsCompleteSave float64 `json:"IsCompleteSave" db:"IsCompleteSave"`
    
}

func(rq *Request)GetByKeyWordRequest(db *sqlx.DB)(rqs []*Request,err error){
	lcCommand := 
	"select top 20 DocNo"+
		",DocDate"+
		",SecMan"+
		",IsConfirm"+
		",IsCancel"+
		",PMCode"+
		",CreatorCode"+
		",CreateDate"+
		",isnull(EditorCode,'') as EditorCode"+
		",isnull(EditDate,'') as EditDate"+
		",IsCompleteSave"+
	" from NPMaster.dbo.TB_PM_Request with(index(PK_TB_PM_Request))"+
		" where Docno like '%%' and iscancel = 0"+
		" order by DocNo desc"

	err = db.Select(&rqs,lcCommand)
	fmt.Println(lcCommand,"keyword",&rqs) 

	if err !=nil{
		return nil,err
	}
	//fmt.Println("CMD : ",rqs) 
	return rqs,nil
}
