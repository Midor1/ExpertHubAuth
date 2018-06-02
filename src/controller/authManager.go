package controller

import (
	"net/http"
	"model"
	"encoding/json"
	"fmt"
)

type AuthRet struct {
	Token    string  `json:"token"`
	Status   int     `json:"status"`
}

type OffRet struct {
	Status   int     `json:"status"`
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	r.ParseMultipartForm(32 << 20)
	username := r.MultipartForm.Value["nickname"][0]
	password := r.MultipartForm.Value["hashkey"][0]
	user := model.User{UserID: 0, UserName: username, Password: password}
	id := model.Login(user)
	if id == -1 {
		info := AuthRet{"",-1}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
	} else {
		token := NewSession(id)
		info := AuthRet{token,0}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
	}
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	r.ParseMultipartForm(32 << 20)
	token := r.MultipartForm.Value["token"][0]
	id := DropSession(token)
	info := OffRet{id}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
	fmt.Println(string(ret))
}