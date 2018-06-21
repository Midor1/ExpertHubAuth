package controller

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"model"
	"fmt"
)


type CreditReq struct {
	Credit     int    `json:"credit"`
}

func GetCredit(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	token, _ := r.Cookie("token")
	uid := CheckSession(token.Value)
	credit := model.Credit(uid)
	info := OffRet{credit}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
}
func ModifyCredit(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	body, _ := ioutil.ReadAll(r.Body)
	token, _ := r.Cookie("token")
	var ereg CreditReq
	_ = json.Unmarshal(body, &ereg)
	uid := CheckSession(token.Value)
	fmt.Println(ereg.Credit)
	fmt.Println(uid)
	model.SetCredit(ereg.Credit, uid)
	info := OffRet{0}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
	fmt.Println(string(ret))
	return
}