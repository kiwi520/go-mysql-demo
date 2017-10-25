package main

// 添加数据
import (
    "fmt"
    //"unicode/utf8"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    db, err := sql.Open("mysql", "homestead:secret@tcp(192.168.13.106:3306)/go?charset=utf8")
    if err != nil {
           panic(err.Error())
    }
    //插入数据
    stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
    checkErr(err)
    
    res, err := stmt.Exec("liyan", "欧萨克斯;理科", "2012-12-09")
    //res, err := stmt.Exec("liyan", utf8.ValidString("理科"), "2012-12-09")   
    
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
