package model

import (
	"database/sql"
	"config"
	"fmt"
)

func Credit(uid int) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT credit FROM users WHERE userid = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	fmt.Println(uid)
	row := stmt.QueryRow(uid)
	credit := -1
	row.Scan(&credit)
	return credit
}

func SetCredit(credit int, uid int) {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE users SET credit = credit + ? WHERE userid = ?")
	defer stmt.Close()
	stmt.Exec(credit, uid)
	return
}