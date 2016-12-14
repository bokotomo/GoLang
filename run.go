package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

const(
  PORT = "8889"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
      fmt.Println("key:", k)
      fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "<a href='/login'>login</a>")
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("static/login.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        token := r.Form.Get("token")
        if token != "" {
            //tokenの合法性を検証します。
        } else {
            //tokenが存在しなければエラーを出します。
        }
        for i:=0;i<10;i++ {
          template.HTMLEscape(w, []byte("UseName " + r.Form.Get("username")))
          template.HTMLEscape(w, []byte("PassWord " + r.Form.Get("password")))
        }
    }
}

func main() {
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":" + PORT, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}