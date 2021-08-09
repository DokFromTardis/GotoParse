package main

import (
	"fmt"
	"net/http"
	"github.com/go-vk-api/vk"
	lp "github.com/go-vk-api/vk/longpoll/user"
)
type serv struct{

}
func (s*serv) ServeHTTP(rw http.ResponseWriter, r*http.Request){
	fmt.Println(r.Method, r.URL.Path)
	rw.Header().Set("Access-Control-Alow-Origin", "*")
}
func main(){
	s:= &serv
	http.ListenAndServe("0.0.0.00", s)
}
