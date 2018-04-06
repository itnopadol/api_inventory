package ctrl

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var dbc *sqlx.DB

var headerKeys = make(map[string]interface{})

func init() {
	dbc = ConnectSql()

}


//func ConnectDB(dbName string)(db *sqlx.DB, err error) {
//	dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/"+ dbName +"?parseTime=true&charset=utf8&loc=Local"
//	db, err = sqlx.Connect("mysql",dsn)
//
//	if err != nil {
//		return nil, err
//	}
//	return db, err
//}

func ConnectSql()(msdb *sqlx.DB) {
	db_host := "192.168.0.7"
	db_name := "npmaster"
	db_user := "sa"
	db_pass := "[ibdkifu"
	port := "1433"
	//dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	msdb = sqlx.MustConnect("mssql",dsn)

	if (msdb.Ping() != nil) {
		fmt.Println("Error")
	}
	return msdb
}



func setHeader(){

	headerKeys = map[string]interface{}{
		"Server":"ProjectCard API",
		"Host":"nebula",
		"Content_Type":"application/json",
		"Access-Control-Allow-Origin":"*",
		"Access-Control-Allow-Methods":"GET, POST, PUT, DELETE",
		"Access-Control-Allow-Headers":"Origin, Content-Type, X-Auth-Token",
	}
}
