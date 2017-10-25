package main

//更新数据
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

    stmt, err := db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)

    res, err := stmt.Exec("哈哈哈哈", 3)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
