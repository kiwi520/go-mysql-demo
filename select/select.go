package main

//查询数据
import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    db, err := sql.Open("mysql", "homestead:secret@tcp(192.168.13.106:3306)/go?charset=utf8")
    if err != nil {
           panic(err.Error())
    }

    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
