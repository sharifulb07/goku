package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	
)

var DB *sql.DB

func Init(connStr string)error{
	var err error

	DB, err=sql.Open("postgres", connStr)
	if err !=nil{
		return nil 
	}

	if err=DB.Ping(); err!=nil{
		return err 
	}
	fmt.Println("Connected to PostgreSql")
	return nil 
}


func Close(){
	if DB!=nil{
		DB.Close()
	}
}