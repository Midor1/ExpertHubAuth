package controller

import (
	"net/http"
	"model"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type AuthRet struct {
	Status   int     `json:"status"`
}

type OffRet struct {
	Status   int     `json:"status"`
}

type UserReq struct {
	UserName     string    `json:"nickname"`
 	Password     string    `json:"hashkey"`
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	body, _ := ioutil.ReadAll(r.Body)
	var user UserReq
	_ = json.Unmarshal(body,&user)
	fmt.Println(user.Password + " " + user.UserName)
	userm := model.User{UserID: 0, UserName: user.UserName, Password: user.Password}
	id := model.Login(userm)
	if id == -1 {
		info := AuthRet{-1}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
	} else {
		token := NewSession(id)
		cookie := http.Cookie{Name:"token",Value:token}
		http.SetCookie(w, &cookie)
		info := AuthRet{0}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
	}
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	r.ParseForm()
	token, _ := r.Cookie("token")
	id := DropSession(token.Value)
	cookie := http.Cookie{Name:"token",Value:""}
	http.SetCookie(w, &cookie)
	info := OffRet{id}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
	fmt.Println(string(ret))
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	body, _ := ioutil.ReadAll(r.Body)
	var user UserReq
	_ = json.Unmarshal(body,&user)
	fmt.Println(user.Password + user.UserName)
	userm := model.User{UserID: 0, UserName: user.UserName, Password: user.Password}
	res := model.Register(userm)
	info := OffRet{res}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
	fmt.Println(string(ret))
}

func UserCheck(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	r.ParseForm()
	token, _ := r.Cookie("token")
	id := DropSession(token.Value)
	info := OffRet{id}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
	fmt.Println(string(ret))
}