package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	//"golang.org/x/tools/go/gcimporter15/testdata"
)

type Label struct{
	ItemCode string `json:"item_code,omitempty" db:"item_code"`
	BarCode string `json:"bar_code,omitempty" db:"bar_code"`
	ItemName string `json:"item_name,omitempty" db:"item_name"`
	UnitCode string `json:"unit_code,omitempty" db:"unit_code"`
	Price float64 `json:"price,omitempty" db:"price"`
	Qty float64 `json:"qty,omitempty" db:"qty"`
	LabelType string `json:"label_type,omitempty" db:"label_type"`
	LabSize string `json:"lab_size,omitempty" db:"lab_size"`
	LabForm string `json:"lab_from,omitempty" db:"lab_from"`
	LabelTypeName string `json:"label_type_name,omitempty" db:"label_type_name"`
	CreatorCode string `json:"creator_code,omitempty" db:"creator_code"`
	CreateDatetime string `json:"create_datetime,omitempty" db:"create_datetime"`
	IsUsed float64 `json:"is_used,omitempty" db:"is_used"`
}

type InsertLabel struct{
	JobID int64 
	ItemCode string `json:"ItemCode" db:"ItemCode"`
	BarCode string `json:"BarCode" db:"BarCode"`
	Qty float64 `json:"Qty" db:"Qty"`
	ReOrder float64 `json:"ReOrder" db:"ReOrder"`
	Suggest float64 `json:"Suggest" db:"Suggest"`
	WHCode string `json:"WHCode" db:"WHCode"`
	ZoneID string `json:"ZoneID" db:"ZoneID"`
	ShelfCode string `json:"ShelfCode" db:"ShelfCode"`
	RowID string `json:"RowID" db:"RowID"`
	ShelfID string `json:"ShelfID" db:"ShelfID"`
	Price float64 `json:"Price" db:"Price"`
	LabelType string `json:"LabelType" db:"LabelType"`
	DateTimeStamp string `json:"DateTimeStamp" db:"DateTimeStamp"`
	CreatorCode string `json:"CreatorCode" db:"CreatorCode"`
	CreateDateTime string `json:"CreateDateTime" db:"CreateDateTime"`
	PathFileName string `json:"PathFileName" db:"PathFileName"`
	UnitCode string `json:"unitcode" db:"unitcode"`
	ReasonCode string `json:"reasoncode" db:"reasoncode"`
}
func(l *Label)GetByUser(keyword string,db *sqlx.DB)(ls []*Label,err error){
	lcCommand := `select	
			a.itemcode as item_code
			,isnull(a.barcode,'') as bar_code
			,isnull(b.name1,'') as item_name
			,a.unitcode as unit_code
			,b.SalePrice1 as price
			,a.qty
			,isnull(a.labeltype,'') as label_type
			,c.LabSize as lab_size
			,d.LabForm as lab_from
			,case when d.LabForm='F1'and c.LabSize='P1' then 'ป้ายธรรมดา 21 ดวง/หน้า'
				  when d.LabForm='F1'and c.LabSize='P2' then 'ป้ายธรรมดา 3 ดวง/หน้า'
				  when d.LabForm='F1'and c.LabSize='P3' then 'ป้ายธรรมดา 2 ดวง/หน้า'
				  when d.LabForm='F1'and c.LabSize='P4' then 'ป้ายธรรมดา A4'
				  when d.LabForm='F2'and c.LabSize='P1' then 'ป้ายราคาพิเศษ  21 ดวง/หน้า'
				  when d.LabForm='F2'and c.LabSize='P2' then 'ป้ายราคาพิเศษ 3 ดวง/หน้า'
				  when d.LabForm='F2'and c.LabSize='P3' then 'ป้ายราคาพิเศษ 2 ดวง/หน้า'
				  when d.LabForm='F2'and c.LabSize='P4' then 'ป้ายราคาพิเศษ A4'
			else 'อื่นๆ' end as label_type_name
			,a.creatorcode as creator_code
			,a.datetimestamp as create_datetime
			,a.isused as is_used
	from	npmaster.dbo.TB_NP_ItemDataOfflineCenter a
		left join bcnp.dbo.bcitem b on a.itemcode = b.code 
		left join npmaster.dbo.TB_PM_Label c on left(a.LabelType,2)=c.LabSize and c.LabUsed = 1
		left join npmaster.dbo.TB_PM_Label d on right(a.LabelType,2)=d.LabForm and c.LabUsed = 1
	where jobid = 4 and a.creatorcode = ? and isused = 0
	group by a.itemcode,isnull(a.barcode,''),a.qty,a.unitcode,isnull(a.labeltype,''),a.datetimestamp,isnull(b.name1,'')
	,a.creatorcode,a.isused,c.LabSize,d.LabForm,b.SalePrice1
	order by datetimestamp,a.itemcode`

	//fmt.Println(lcCommand) 

	err = db.Select(&ls,lcCommand,keyword)
	//fmt.Println("CMD",lcCommand,keyword)
	if err !=nil{
		return nil,err
	}


	return ls,nil
}

func(il *InsertLabel)CheckExists(db *sqlx.DB, itemcode string, barcode string, unitcode string, labeltype string, CreatorCode string) int {
	var chkRow int 
	fmt.Println("Begin CheckExists")
	lccommand := `select isnull(count(itemcode),0) as vCount 
					from npmaster.dbo.TB_NP_ItemDataOfflineCenter 
				  where jobid = 4 and isused = 0 and itemcode = ? and barcode = ? 
				  and unitcode = ? and labeltype = ? and CreatorCode = ?`
	err := db.Get(&chkRow, lccommand, itemcode,barcode,unitcode,labeltype,CreatorCode)
	if err !=nil{
		return 0
	}
	//chkRow, _ := rs.RowsAffected()
	fmt.Println("itemcode",itemcode,barcode,unitcode,labeltype,CreatorCode,chkRow)
	// if chkRow > 0 {
	// 	fmt.Println("data aleady exists!!! cannot insert this number : ", pj.Code)
	// 	return 1	
	// }
	return chkRow
}


func (il *InsertLabel)Insert(db *sqlx.DB) (NewProject string, err error) {

	lccommand := `set dateformat dmy
		INSERT INTO nebula.npmaster.dbo.TB_NP_ItemDataOfflineCenter
		(JobID
		,ItemCode
		,BarCode
		,Qty
		,ReOrder
		,Suggest
		,WHCode
		,ZoneID
		,ShelfCode
		,RowID
		,ShelfID
		,Price
		,LabelType
		,DateTimeStamp
		,CreatorCode
		,CreateDateTime
		,PathFileName
		,UnitCode
		,ReasonCode) 
	VALUES (4,?,?,?,0,0,'','','','','',?,?
			,cast(rtrim(day(getdate()))+'/'+rtrim(month(getdate()))+'/'+rtrim(year(getdate())) as datetime)
			,?,getdate(),'',?,'Mobile App')`
	_, err = db.Exec(lccommand,il.ItemCode,il.BarCode,il.Qty,il.Price,il.LabelType,il.CreatorCode,il.UnitCode)
	fmt.Println(lccommand)
	if err != nil {
		return il.ItemCode, err
	}
	return il.ItemCode+" Completed Insert", err
}

func(il *InsertLabel)Update(db *sqlx.DB) (msg string, err error) {
	// Update Project
	lccommand := `set dateformat dmy 
				update nebula.npmaster.dbo.TB_NP_ItemDataOfflineCenter 
				set	qty = ? 
				where jobid = 4 and itemcode = ? and barcode = ? and unitcode = ? and labeltype = ?`
			
	_, err = db.Exec(lccommand,il.Qty,il.ItemCode,il.BarCode,il.UnitCode,il.LabelType)
	if err != nil {
		msg = "Update Error "
		return msg, err
		}
		fmt.Println("CMD",lccommand)
	msg = "Completed updated"
	return msg, err
}

