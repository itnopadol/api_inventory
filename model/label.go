package model

import (
	"github.com/jmoiron/sqlx"
	//"fmt"
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

