package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Post struct {
	Emailaddress string `json:"emailaddress"`
	Password     string `json:"password"`
}

type Res struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Name    string `json:"name"`
}

// GOのexecuteTemplateが何もないhtmlファイルで動作するか？
func executeTemplate(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./front/dist/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, "test")
	fmt.Println("success")
}



func getLoginJSON(w http.ResponseWriter, r *http.Request) {
	// 受け取ったpostをparseする
	r.ParseForm()
	// 受け取った値を出力
	emailaddress := r.PostFormValue("emailaddress")
	password := r.PostFormValue("password")
	fmt.Println(emailaddress)
	fmt.Println(password)
	post(emailaddress,password)
	// 返却JSONの生成
	 res := Res{
		Id:1,
		Content:"Hello World!",
		Name:"David",
	}
	// 返却JSONをヘッダーにセットするまで
	output , err := json.MarshalIndent(&res,"","\t\t")
	if err != nil{
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}

func main() {
	staticDir := "./dist"
	// staticDir := "./front/dist"
	fmt.Println("First Access:       " + time.Now().String() + "\n")

	server := http.Server{
		Addr: ":8080",
	}
	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	http.HandleFunc("/login", getLoginJSON)

	server.ListenAndServe()
}
