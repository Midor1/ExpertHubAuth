package main

import (
	"config"
	"github.com/gorilla/mux"
	"net/http"
	"controller"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func Options(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods","POST")
}

func main() {
	var conf, _ = config.C.GetConfig()
	rtr := mux.NewRouter()
	rtr.HandleFunc("/hello", SayHello)
	rtr.HandleFunc("/login",controller.UserLogin).Methods("POST")
	rtr.HandleFunc("/logout",controller.UserLogout).Methods("POST")
	rtr.HandleFunc("/check",controller.UserCheck).Methods("POST")
	rtr.HandleFunc("/register",controller.UserRegister).Methods("POST")
	rtr.HandleFunc("/grant",controller.ExpertRegister).Methods("POST")
	rtr.HandleFunc("/validate",controller.EmailValidate).Methods("POST")
	rtr.HandleFunc("/map",controller.Mapping).Methods("POST")
	rtr.HandleFunc("/crud",controller.InfoCRUD).Methods("POST")
	rtr.HandleFunc("/credit",controller.GetCredit).Methods("GET")
	rtr.HandleFunc("/credit",controller.ModifyCredit).Methods("POST")
	rtr.HandleFunc("/{any}",Options).Methods("OPTIONS")
	http.Handle("/", rtr)
	http.ListenAndServe(conf.Server.Port, nil)
}