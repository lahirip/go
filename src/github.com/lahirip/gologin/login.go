package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)


func home(w http.ResponseWriter, r *http.Request) {
   t, _ := template.ParseFiles("home.html")
   t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}


func main() {
    http.HandleFunc("/", home) // setting router rule
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":8090", nil) //listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
