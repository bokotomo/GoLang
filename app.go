package main

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
)

const(
  PORT = "8889"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  t, _ := template.ParseFiles("view/index.html")
  t.Execute(w, nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    t, _ := template.ParseFiles("view/login.html")
  	data := struct {
  		Title string
  		UserName string
  	}{
  		Title: "MyPage",
  		UserName: "",
  	}
    t.Execute(w, data)
  } else {
    r.ParseForm()
    t, _ := template.ParseFiles("view/login.html")
  	data := struct {
  		Title string
  		UserName string
  	}{
  		Title: "MyPage",
  		UserName: r.Form.Get("username"),
  	}
    t.Execute(w, data)
  }
}

func rooting(){
  fmt.Println("port :"+PORT)
  http.HandleFunc("/", indexPage)
  http.HandleFunc("/login", loginPage)
  err := http.ListenAndServe(":" + PORT, nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func main() {
  rooting()
}