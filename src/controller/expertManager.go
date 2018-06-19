package controller

import (
	"net/http"
	"model"
	"encoding/json"
	"fmt"
	"config"
	"math/rand"
	"time"
	"strconv"
	"sort"
	"io/ioutil"
)

type MapRet struct {
	EIDs    []int    `json:"eid"`
}

type ExpertRegReq struct {
	UserName     string    `json:"nickname"`
	Email        string    `json:"email"`
}

type EmailValReq struct {
	Captcha     string    `json:"captcha"`
}

func ExpertRegister(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	body, _ := ioutil.ReadAll(r.Body)
	token, _ := r.Cookie("token")
	var ereg ExpertRegReq
	_ = json.Unmarshal(body, &ereg)
	email := ereg.Email
	nickname := ereg.UserName
	uid := CheckSession(token.Value)
	if uid == -1 {
		info := OffRet{-1}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
		return
	}
	eid := model.NewExpert(uid, nickname)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	model.SetCaptcha(vcode, eid)
	config.SendMail(email,nickname,"专家认证验证码","您的验证码是：" + vcode + "。")
	info := OffRet{0}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
	fmt.Println(string(ret))
	return
}

func EmailValidate(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	body, _ := ioutil.ReadAll(r.Body)
	token, _ := r.Cookie("token")
	var ereg EmailValReq
	_ = json.Unmarshal(body, &ereg)
	uid := CheckSession(token.Value)
	eid := model.CheckCaptcha(ereg.Captcha, uid)
	info := OffRet{eid}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
	fmt.Println(string(ret))
	return
}

func InfoCRUD(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	r.ParseMultipartForm(32 << 20)
	if len(r.MultipartForm.Value["eid"]) > 0 {
		eid, _ := strconv.Atoi(r.MultipartForm.Value["eid"][0])
		if len(r.MultipartForm.Value["token"]) > 0 {
			token := r.MultipartForm.Value["token"][0]
			uid := CheckSession(token)
			eids := model.GetMapping(uid)
			sort.Ints(eids)
			has := sort.SearchInts(eids, eid)
			if has == len(eids) {
				info := OffRet{-1}
				ret, _ := json.Marshal(info)
				fmt.Fprint(w, string(ret))
				fmt.Println(string(ret))
				return
			}
			result := 1
			if len(r.MultipartForm.Value["nickname"]) > 0 {
				nickname := r.MultipartForm.Value["nickname"][0]
				result *= model.SetNickname(nickname, eid)
			}
			if len(r.MultipartForm.Value["gender"]) > 0 {
				gender := r.MultipartForm.Value["gender"][0]
				result *= model.SetGender(gender, eid)
			}
			if len(r.MultipartForm.Value["email"]) > 0 {
				email := r.MultipartForm.Value["email"][0]
				result *= model.SetEmail(email, eid)
			}
			if len(r.MultipartForm.Value["tel"]) > 0 {
				tel := r.MultipartForm.Value["tel"][0]
				result *= model.SetTel(tel, eid)
			}
			if len(r.MultipartForm.Value["subgroup"]) > 0 {
				subgroup := r.MultipartForm.Value["subgroup"][0]
				result *= model.SetSubgroup(subgroup, eid)
			}
			if len(r.MultipartForm.Value["category"]) > 0 {
				category := r.MultipartForm.Value["category"][0]
				result *= model.SetCategory(category, eid)
			}
			if len(r.MultipartForm.Value["avatar"]) > 0 {
				avatar := r.MultipartForm.Value["avatar"][0]
				result *= model.SetAvatar(avatar, eid)
			}
			if result == 1 {
				info := OffRet{0}
				ret, _ := json.Marshal(info)
				fmt.Fprint(w, string(ret))
				fmt.Println(string(ret))
				return
			} else {
				info := OffRet{-1}
				ret, _ := json.Marshal(info)
				fmt.Fprint(w, string(ret))
				fmt.Println(string(ret))
				return
			}
		}
		info := model.FetchExpertData(eid)
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
		return
	}
}

func Mapping(w http.ResponseWriter, r *http.Request) {
	PreprocessXHR(&w,r)
	token, _ := r.Cookie("token")
	id := CheckSession(token.Value)
	eid := model.GetMapping(id)
	info := MapRet{eid}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
	fmt.Println(string(ret))
	return
}