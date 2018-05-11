package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//"golang.org/x/tools/go/gcimporter15/testdata"
	"github.com/kataras/iris/core/errors"
)

type Promotion struct {
	CheckJob       int             `json:"check_job" db:"check_job"`
	DocNo          string          `json:"doc_no" db:"doc_no"`
	DocDate        string          `json:"doc_date" db:"doc_date"`
	SecMan         string          `json:"sec_man" db:"sec_man"`
	StartPromo     string          `json:"start_promo" db:"start_promo"`
	IsConfirm      int             `json:"is_con_firm" db:"is_con_firm"`
	IsCancel       int             `json:"is_cancel" db:"is_cancel"`
	PMCode         string          `json:"pm_code" db:"pm_code"`
	CreatorCode    string          `json:"creator_code" db:"creator_code"`
	CreateDate     string          `json:"create_date" db:"create_date"`
	EditorCode     string          `json:"editor_code" db:"editor_code"`
	EditDate       string          `json:"edit_date" db:"edit_date"`
	IsCompleteSave int             `json:"is_complete_save" db:"is_complete_save"`
	Subs           []*PromotionSub `json:"subs"`
}

type PromotionSub struct {
	ItemCode      string  `json:"item_code" db:"item_code"`
	ItemName      string  `json:"item_name" db:"item_name"`
	UnitCode      string  `json:"unit_code" db:"unit_code"`
	Price         float64 `json:"price" db:"price"`
	FromQty       float64 `json:"from_qty" db:"from_qty"`
	ToQty         float64 `json:"to_qty" db:"to_qty"`
	Discount      float64 `json:"discount" db:"discount"`
	DiscountType  int     `json:"discount_type" db:"discount_type"`
	DiscountWord  string  `json:"discount_word" db:"discount_word"`
	PromoPrice    float64 `json:"promo_price" db:"promo_price"`
	Mydescription string  `json:"mydescription" db:"mydescription"`
	LineNumber    int     `json:"line_number" db:"line_number"`
	IsBrochure    int     `json:"is_brochure" db:"is_brochure"`
	PromoMember   int     `json:"promo_member" db:"promo_member"`
	PromotionType string  `json:"promotion_type" db:"promotion_type"`
	CancelCode    string  `json:"cancel_code" db:"cancel_code"`
	CancelDate    string  `json:"cancel_date" db:"cancel_date"`
	DateStart     string  `json:"date_start" db:"date_start"`
	DateEnd       string  `json:"date_end" db:"date_end"`
}

type PromotionType struct {
	Code          string `json:"code" db:"Code"`
	NameEng       string `json:"name_eng" db:"NameEng"`
	NameThai      string `json:"name_thai" db:"NameThai"`
	Mydescription string `json:"mydescription" db:"Mydescription"`
	NameFull      string `json:"name_full" db:"NameFull"`
}

type PromotionMaster struct {
	PmCode        string `json:"pm_code" db:"pm_code"`
	PmName        string `json:"pm_name" db:"pm_name"`
	NameFull      string `json:"name_full" db:"name_full"`
	PmType        int    `json:"pm_type" db:"pm_type"`
	DateStart     string `json:"date_start" db:"date_start"`
	DateEnd       string `json:"date_end" db:"date_end"`
	Mydescription string `json:"mydescription" db:"mydescription"`
	IsCancel      int    `json:"is_cancel" db:"is_cancel"`
	CreatorCode   string `json:"creator_code" db:"creator_code"`
	CreateDate    string `json:"create_date" db:"create_date"`
	EditorCode    string `json:"editor_code" db:"editor_code"`
	EditDate      string `json:"edit_date" db:"edit_date"`
}

type SectionMan struct {
	SecmanCode string `json:"secman_code" db:"secman_code"`
	SecmanName string `json:"secman_name" db:"secman_name"`
	NameFull   string `json:"name_full" db:"name_full"`
	SaleCode   string `json:"sale_code" db:"sale_code"`
	UserId     string `json:"user_id" db:"user_id"`
}

type checkerror struct {
	IsDuplicate int    `json:"is_duplicate" db:"IsDuplicate"`
	Duplicate   string `json:"duplicate" db:"Duplicate"`
	PMCode      string `json:"pm_code" db:"PMCode"`
	PMName      string `json:"pm_name" db:"PMName"`
	DateStart   string `json:"date_start" db:"DateStart"`
	DateEnd     string `json:"date_end" db:"DateEnd"`
}

func (rq *Promotion) GetByKeyWordRequest(keyword string, db *sqlx.DB) (rqs []*Promotion, err error) {

	lcCommand :=
		"select top 20 isnull(DocNo,'') as doc_no" +
			",isnull(DocDate,'') as doc_date" +
			",isnull(SecMan,'') as sec_man" +
			",isnull(IsConfirm,0) as is_con_firm" +
			",isnull(IsCancel,0) as is_cancel" +
			",isnull(PMCode,'') as pm_code" +
			",isnull(CreatorCode,'') as creator_code" +
			",isnull(CreateDate,'') as create_date" +
			",isnull(EditorCode,'') as editor_code" +
			",isnull(EditDate,'') as edit_date" +
			",isnull(IsCompleteSave,'') as is_complete_save" +
			" from NPMaster.dbo.TB_PM_Request" +
			" where isnull(Docno,'') like '%" + keyword + "%' or isnull(SecMan,'') like '%" + keyword + "%'" +
			" or isnull(PMCode,'') like '%" + keyword + "%' and isnull(iscancel,0) = 0" +
			" order by DocNo desc"

	err = db.Select(&rqs, lcCommand)
	//fmt.Println(lcCommand,"keyword",&rqs) 
	if err != nil {
		return nil, err
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
	fmt.Println("keyword : ", keyword)
	return rqs, nil
}

func (pt *PromotionType) GetPromotionType(db *sqlx.DB) (pts []*PromotionType, err error) {

	lcCommand := `select Code
				,NameEng
				,NameThai
				,Mydescription
				, Code+'/'+NameEng+'/'+NameThai as NameFull
				from Nebula.NPMaster.dbo.TB_PM_Type 
				order by Code`
	err = db.Select(&pts, lcCommand)
	if err != nil {
		return nil, err
	}
	return pts, nil
}

func (pm *PromotionMaster) GetPromotionMaster(db *sqlx.DB) (pms []*PromotionMaster, err error) {

	lcCommand := `set dateformat dmy
		select isnull(PMCode,'') as pm_code
		,isnull(PMName,'') as pm_name
		,isnull(PMCode+'/'+PMName,'') as name_full
		,isnull(pm_type,0) as pm_type
		,isnull(DateStart,'') as date_start
		,isnull(DateEnd,'') as date_end
		,isnull(MyDescription,'') as mydescription
		,isnull(Iscancel,'') as is_cancel
		,isnull(CreatorCode,'') as creator_code
		,isnull(CreateDate,'') as create_date
		,isnull(EditorCode,'') as editor_code
		,isnull(EditDate,'') as edit_date
		from NPMaster.dbo.TB_PM_PromotionMaster with(index(PK_TB_PM_PromotionMaster))
		where cast(dbo.FT_CG_DateToString(DateEnd) as datetime) >= cast(dbo.FT_CG_DateToString(getdate()-7) as datetime)
		order by RowOrder desc`

	err = db.Select(&pms, lcCommand)
	if err != nil {
		return nil, err
	}
	return pms, nil
}

func (sm *SectionMan) GetSectionMan(db *sqlx.DB) (sms []*SectionMan, err error) {

	lcCommand := `set dateformat dmy
		select distinct rtrim(st.salecode)+'/'+ltrim(st.salename) as secman_code
		,isnull(st.salename,'') as secman_name
		,rtrim(st.salecode)+'/'+ltrim(st.salename) as name_full 
		,isnull(st.salecode,'') as sale_code
		,isnull(sL.Userid,'') as user_id
		from npdb.dbo.tb_inc_saleteam as st
		inner join bcnp.dbo.BCsale as sL on st.salecode = sL.code and sL.activestatus=1
		where isnull(st.degreecode,'')='L2' and isnull(st.teamstatus,'')<>1 and st.profitcenter in ('S01','S02')
		and enyear=year(getdate()) and monthofyear=month(getdate())`

	err = db.Select(&sms, lcCommand)
	if err != nil {
		return nil, err
	}
	return sms, nil
}

func (pm *Promotion) InsertAndUpdatePromotion(db *sqlx.DB) error {
	c := checkerror{}

	fmt.Println("Insert Promotion")

	sql := `exec bcnp.dbo.USP_PM_NewInsertRequest ?, ?, ?, ?, ?, ?`
	fmt.Println("Sql =", sql, pm.CheckJob, pm.DocNo, pm.DocDate, pm.SecMan, pm.PMCode, pm.CreatorCode)
	_, err := db.Exec(sql, pm.CheckJob, pm.DocNo, pm.DocDate, pm.SecMan, pm.PMCode, pm.CreatorCode)
	if err != nil {
		return err
	}
	for _, sub := range pm.Subs {
		var hotprice string

		fmt.Println("ItemCode = ",sub.ItemCode)
		err := c.CheckErrDuplicate(db, pm.PMCode, sub.ItemCode, sub.UnitCode)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("IsDuplicate = ",c.IsDuplicate)
		if (c.IsDuplicate == 0) {
			if (sub.PromotionType == "11") {
				hotprice = "S02"
			} else {
				hotprice = ""
			}

			sqlsub := `exec bcnp.dbo.USP_PM_NewInsertRequestSub ?,?,?,?,?,?,?,1,99999,?,?,?,?,?,0,?,?,?,?,?`
			_, err = db.Exec(sqlsub, c.IsDuplicate, pm.IsCompleteSave, pm.DocNo, sub.ItemCode, sub.ItemName, sub.UnitCode, sub.Price, sub.Discount, sub.DiscountType, sub.DiscountWord, sub.PromoPrice, sub.Mydescription, sub.LineNumber, sub.IsBrochure, sub.PromoMember, sub.PromotionType, hotprice)
			if err != nil {
				return err
			}
			sqlduplicate := `exec bcnp.dbo.USP_PM_DeleteCheckDuplicatItemLine ?,?,?`
			_, err = db.Exec(sqlduplicate, pm.DocNo, pm.PMCode, pm.CreatorCode)
			if err != nil {
				return err
			}
		}else{
			return errors.New("ItemCode have promotion request")
		}

	}

	return nil
}

func (c *checkerror) CheckErrDuplicate(db *sqlx.DB, pmcode string, itemcode string, unitcode string) error {
	sqlerr := `exec bcnp.dbo.USP_PM_ItemDuplicate_exist ?, ?, ?`
	fmt.Println("sqlerror = ",sqlerr, pmcode, itemcode, unitcode)
	err := db.Get(c, sqlerr, pmcode, itemcode, unitcode)
	if err != nil {
		return err
	}

	return nil
}

func (pm *Promotion) PromotionCancel(db *sqlx.DB) error {
	if (pm.DocNo != ""){
		sql := `exec bcnp.dbo.USP_PM_DeletePMRequest ?`
		_, err := db.Exec(sql, pm.DocNo)
		if err != nil {
			return err
		}
	}else{
		return errors.New("DocNo is empty")
	}


	return nil
}

func (pm *Promotion)PromotionCancelItem(db *sqlx.DB) error{
	if(len(pm.Subs) != 0){
		for _, sub := range pm.Subs {
			sql := `exec dbo.USP_PM_CancelItemPMRequest ?, ?, ?`
			_, err := db.Exec(sql, pm.DocNo, sub.ItemCode, sub.UnitCode)
			if err != nil {
				return err
			}
		}
	}else{
		return errors.New("DocNo not have list item cancel")
	}


	return nil
}
