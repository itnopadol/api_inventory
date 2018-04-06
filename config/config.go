package config

import (
	"encoding/json"
	"log"
	"os"
	"fmt"
)

const (
	API_HOST string   = "http://nopadol.net:8001"
	ORG_NOPADOL int = 1
	ORG_NAVA   int  = 2
)

//TODO: เมื่อรันจริงต้องเปลี่ยนเป็น Docker Network Bridge IP เช่น 172.17.0.3 เป็นต้น
type Config struct {
	DBHost string `json:"db_host"`
	DBName string `json:"db_name"`
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
	DBPort string `json:"port"`
}

func LoadDSN(fileName string,dbType int) string {

	// dbType 0 = mySql , 1 = MsSql
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Err Open file %v: Error is: %v", file, err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	c := new(Config)
	err = decoder.Decode(&c)
	if err != nil {
		log.Println("error Decode Json:", err)
	}
	//log.Printf("Test Variable: %s", c.DBHost)
	// mysql connect string
	dsn := ""
	if dbType == 0 {
		dsn = c.DBUser + ":" + c.DBPass + "@" + c.DBHost + "/" + c.DBName + "?parseTime=true"
		//dsn := "root:mypass@tcp(nava.work:3306)/sys?parseTime=true"
	}
	// Microsoft SQLserver Pattern
	if dbType == 1 {
		dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", c.DBHost, c.DBUser, c.DBPass, c.DBPort, c.DBName)
		//dsn := "root:mypass@tcp(nava.work:3306)/sys?parseTime=true"
	}
	log.Println("DSN =", dsn)


	return dsn
}
