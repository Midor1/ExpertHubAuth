package model

import (
	"database/sql"
	"config"
	"fmt"
)

type InfoRet struct {
	ExpertID       int    `json:"eid"`
	UserID         int    `json:"uid"`
	Nickname       string `json:"nickname"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	Tel            string `json:"tel"`
	Subgroup       string `json:"subgroup"`
	Category       string `json:"category"`
	Avatar         string `json:"avatar"`
	AccountStatus  int    `json:"status"`
}

func SetCaptcha(captha string, uid int) {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE experts SET captcha = ?,accountstatus = ? WHERE expertid = ?")
	defer stmt.Close()
	stmt.Exec(captha, 1, uid)
	return
}

func CheckCaptcha(captcha string, uid int) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("SELECT expertid FROM experts WHERE userid = ? AND captcha = ?")
	defer stmt.Close()
	rows, _ := stmt.Query(uid, captcha)
	eID := -1
	for rows.Next() {
		_ = rows.Scan(&eID)
		stmt, _ := db.Prepare("UPDATE experts SET accountstatus = ?,captcha = NULL WHERE expertid = ?")
		stmt.Exec(0,eID)
	}
	return eID
}

func NewExpert(uid int, nickname string) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO experts(userid, nickname) VALUES (?,?)")
	defer stmt.Close()
	res, _ := stmt.Exec(uid, nickname)
	eid, _ := res.LastInsertId()
	return int(eid)
}

func GetMapping(uid int) []int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("SELECT expertid FROM experts WHERE userid = ?")
	defer stmt.Close()
	rows, _ := stmt.Query(uid)
	var eID int
	var eIDs []int
	for rows.Next() {
		_ = rows.Scan(&eID)
		eIDs = append(eIDs, eID)
	}
	return eIDs
}

func FetchExpertData(eid int) InfoRet {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("SELECT * FROM experts WHERE expertid = ?")
	defer stmt.Close()
	row := stmt.QueryRow(eid)
	var info InfoRet
	row.Scan(&info.ExpertID,&info.UserID,&info.Nickname,&info.Gender,&info.Email,&info.Tel,&info.Subgroup,&info.Category,&info.Avatar,nil,&info.AccountStatus)
	return info
}

func SetNickname(nickname string, eid int) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE experts SET nickname = ? WHERE expertid = ?")
	defer stmt.Close()
	stmt.Exec(nickname,eid)
	return 1
}

func SetGender(gender string, eid int) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE experts SET gender = ? WHERE expertid = ?")
	defer stmt.Close()
	stmt.Exec(gender,eid)
	return 1
}

func SetEmail(email string, eid int) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE experts SET email = ? WHERE expertid = ?")
	defer stmt.Close()
	stmt.Exec(email,eid)
	return 1
}

func SetTel(tel string, eid int) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE experts SET tel = ? WHERE expertid = ?")
	defer stmt.Close()
	stmt.Exec(tel,eid)
	return 1
}

func SetSubgroup(subgroup string, eid int) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE experts SET subgroup = ? WHERE expertid = ?")
	defer stmt.Close()
	stmt.Exec(subgroup,eid)
	return 1
}

func SetCategory(category string, eid int) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE experts SET category = ? WHERE expertid = ?")
	defer stmt.Close()
	stmt.Exec(category,eid)
	return 1
}

func SetAvatar(avatar string, eid int) int {
	db, err := sql.Open("mysql", config.C.Database.SQLString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println(avatar)
	stmt, _ := db.Prepare("UPDATE experts SET avatar = ? WHERE expertid = ?")
	defer stmt.Close()
	stmt.Exec(avatar,eid)
	return 1
}