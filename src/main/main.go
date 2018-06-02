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

func main() {
	var conf, _ = config.C.GetConfig()
	rtr := mux.NewRouter()
	rtr.HandleFunc("/hello", SayHello)
	rtr.HandleFunc("/login",controller.UserLogin).Methods("POST")
	rtr.HandleFunc("/logout",controller.UserLogout).Methods("POST")
	http.Handle("/", rtr)
	http.ListenAndServe(conf.Server.Port, nil)
}