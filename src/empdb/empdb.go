package empdb

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	DB_USER = "postgres"
	DB_PASSWD = "test1234"
	DB_NAME = "test"
)

func checkErr(err error) {
	if err != nil {
		panic(err)	
	}
}

type EmpData struct {
	FirstName string
	LastName string
	EmpId string
	Phone string 
}

func (table *EmpData) Dbcommit() (result bool) {
	result = false
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	DB_USER, DB_PASSWD, DB_NAME)
	db , err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	if db.Ping() == nil {
		val , err := db.Query(`CREATE TABLE IF NOT EXISTS empdata( 
			id integer, 
			FirstName varchar(20), 
			LastName varchar(10),
			EmpId varchar(20), 
			Phone varchar(10))`)
		checkErr(err)
		insertQuery := fmt.Sprintf(`INSERT INTO empdata (FirstName, LastName, 
		EmpId, Phone)
		VALUES ('%s', '%s', '%s', '%s')`,table.FirstName, table.LastName, table.EmpId, table.Phone)
		value , err := db.Query(insertQuery)
		fmt.Println(val, value, err)
		result = true
	}
	return

}

