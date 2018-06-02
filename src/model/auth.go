package model

import (
	"database/sql"
	"config"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	UserID   int
	UserName string
	Password string
}

func Login(user User) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("SELECT userid FROM users WHERE nickname = ? AND hashkey = ?")
	defer stmt.Close()
	rows, _ := stmt.Query(user.UserName, user.Password)
	UserID := -1
	for rows.Next() {
		_ = rows.Scan(&UserID)
	}
	return UserID
}
