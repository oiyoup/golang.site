package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "log"
)

func main() {
    db, err := sql.Open("mysql", "lpop:lpop5913kc@tcp(127.0.0.1:3306)/dc_lpop")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    var id int64
    var yyyymmdd string
    var merchantId string
    var orderCode string
    var productCode string
    rows, err := db.Query(`
        select id, yyyymmdd, merchant_id, order_code, product_code 
        from   ttranslog
        where  id = ?
    `, 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        if err := rows.Scan(&id, &yyyymmdd, &merchantId, &orderCode, &productCode); err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, yyyymmdd, merchantId, orderCode, productCode)
    }
}
