package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//"golang.org/x/tools/go/gcimporter15/testdata"
)

type Request struct{
	DocNo string `json:"doc_no" db:"doc_no"`
	DocDate string `json:"doc_date" db:"doc_date"`
	SecMan string `json:"sec_man" db:"sec_man"`
	IsConfirm int `json:"is_con_firm" db:"is_con_firm"`
	IsCancel int `json:"is_cancel" db:"is_cancel"`
	PMCode string `json:"pm_code" db:"pm_code"`
	CreatorCode string `json:"creator_code" db:"creator_code"`
	CreateDate string `json:"create_date" db:"create_date"`
	EditorCode string `json:"editor_code" db:"editor_code"`
	EditDate string `json:"edit_date" db:"edit_date"`
	IsCompleteSave int `json:"is_complete_save" db:"is_complete_save"`
    Subs []*Requestsub `json:"subs"`
}

type Requestsub struct{
	ItemCode string `json:"item_code" db:"item_code"`
	ItemName string `json:"item_name" db:"item_name"`
	UnitCode string `json:"unit_code" db:"unit_code"`
	Price float64 `json:"price" db:"price"`
	FromQty float64 `json:"from_qty" db:"from_qty"`
	ToQty float64 `json:"to_qty" db:"to_qty"`
	Discount float64 `json:"discount" db:"discount"`
	DiscountType int `json:"discount_type" db:"discount_type"`
	DiscountWord string `json:"discount_word" db:"discount_word"`
	PromoPrice float64 `json:"promo_price" db:"promo_price"`
	Mydescription string `json:"mydescription" db:"mydescription"`
	LineNumber int `json:"line_number" db:"line_number"`
	IsBrochure int `json:"is_brochure" db:"is_brochure"`
	PromoMember int `json:"promo_member" db:"promo_member"`
	PromotionType string `json:"promotion_type" db:"promotion_type"`
	CancelCode string `json:"cancel_code" db:"cancel_code"`
	CancelDate string `json:"cancel_date" db:"cancel_date"`
	DateStart string `json:"date_start" db:"date_start"`
	DateEnd string `json:"date_end" db:"date_end"`
}

type Promotiontype struct{
	Code string `json:"code" db:"Code"`
	NameEng string `json:"name_eng" db:"NameEng"`
	NameThai string `json:"name_thai" db:"NameThai"`
	Mydescription string `json:"mydescription" db:"Mydescription"`
	NameFull string `json:"name_full" db:"NameFull"`
}

func(rq *Request)GetByKeyWordRequest(keyword string,db *sqlx.DB)(rqs []*Request,err error){
	
	lcCommand := 
	"select top 20 isnull(DocNo,'') as doc_no"+
		",isnull(DocDate,'') as doc_date"+
		",isnull(SecMan,'') as sec_man"+
		",isnull(IsConfirm,0) as is_con_firm"+
		",isnull(IsCancel,0) as is_cancel"+
		",isnull(PMCode,'') as pm_code"+
		",isnull(CreatorCode,'') as creator_code"+
		",isnull(CreateDate,'') as create_date"+
		",isnull(EditorCode,'') as editor_code"+
		",isnull(EditDate,'') as edit_date"+
		",isnull(IsCompleteSave,'') as is_complete_save"+
	 " from NPMaster.dbo.TB_PM_Request"+
		" where isnull(Docno,'') like '%"+keyword+"%' or isnull(SecMan,'') like '%"+keyword+"%'"+
		" or isnull(PMCode,'') like '%"+keyword+"%' and isnull(iscancel,0) = 0"+
		" order by DocNo desc"

	err = db.Select(&rqs,lcCommand)
	//fmt.Println(lcCommand,"keyword",&rqs) 
	if err !=nil{
		return nil,err
	}

	for _, rq2 := range rqs {
		lcCommandsub := `select isnull(b.ItemCode,'') as item_code
		,isnull(b.ItemName,'') as item_name
		,isnull(b.UnitCode,'') as unit_code
		,isnull(b.Price,0) as price
		,isnull(b.FromQty,0) as from_qty
		,isnull(b.ToQty,0) as to_qty
		,isnull(b.Discount,0) as discount
		,isnull(b.DiscountType,0) as discount_type
		,isnull(b.DiscountWord,'') as discount_word
		,isnull(b.PromoPrice,0) as promo_price
		,isnull(b.Mydescription,'') as mydescription
		,isnull(b.LineNumber,0) as line_number
		,isnull(b.IsBrochure,0) as is_brochure
		,isnull(b.PromoMember,0) as promo_member
		,isnull(b.PromotionType,'') as promotion_type
		,isnull(b.CancelCode,'') as cancel_code
		,isnull(b.CancelDate,'') as cancel_date
		,isnull(e.datestart,'') as date_start
		,isnull(e.dateend,'') as date_end
	from NPMaster.dbo.TB_PM_Request a 
		left join NPMaster.dbo.TB_PM_Requestsub b on a.docno = b.docno
		left join NPMaster.dbo.TB_PM_Type c on b.promotiontype = c.code
		left join npdb.dbo.tb_inc_saleteam as d on a.secman = d.salecode and  isnull(d.degreecode,'')='L2' and isnull(d.teamstatus,'')<>1 and enyear=year(getdate()) and monthofyear=month(getdate())
		left join NPMaster.dbo.TB_PM_PromotionMaster e on a.pmcode = e.pmcode
	where a.Docno is not null and b.iscancel = 0 and a.docno = ?
	order by a.docno desc,a.docdate desc,b.linenumber`

		err = db.Select(&rq2.Subs, lcCommandsub, rq2.DocNo)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	fmt.Println("keyword : ",keyword) 
	return rqs,nil
}

func(pt *Promotiontype)GetPromotionType(db *sqlx.DB)(pts []*Promotiontype,err error){

	lcCommand := `select Code,NameEng,NameThai,Mydescription, Code+'/'+NameEng+'/'+NameThai as NameFull
				from Nebula.NPMaster.dbo.TB_PM_Type 
				order by Code`
	err = db.Select(&pts,lcCommand)
	if err !=nil{
		return nil,err
	}
	return pts,nil
}


