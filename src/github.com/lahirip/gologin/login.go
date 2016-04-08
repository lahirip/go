package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
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
        username := string(r.Form["username"][0])
	password := string(r.Form["password"][0])
	fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])	

	check := validateLogin(username, password) 
        fmt.Println("check:", check)
	if check == 1 {
            fmt.Println("here...")
	    t, _ := template.ParseFiles("welcome.html")
            t.Execute(w, nil)
	}  else {
	    t, _ := template.ParseFiles("login.html")
            t.Execute(w, nil)
        }
    }
}

func validateLogin(username , password string) int {
    //TCP on a remote host, e.g. Amazon RDS:
    //db, err := sql.Open("mysql", id:password@tcp(your-amazonaws-uri.com:3306)/dbname)    

    db, err := sql.Open("mysql", "root:alap@/test?charset=utf8")
    checkErr(err)

    // query
   cnt := 0
   err = db.QueryRow("SELECT count(*) FROM userlogin WHERE username = ? and password= ?", username, password).Scan(&cnt)
   checkErr(err)
    
   return cnt
}

func checkErr(err error) {
    if err != nil {
        panic(err)
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
