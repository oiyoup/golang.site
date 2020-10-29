package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "log"
)

type TTranslog2 struct {
    Id          int64          `db:"id"`
    YYYYMMDD    string         `db:"yyyymmdd"`
    AffiliateId sql.NullString `db:"affiliate_id"`
    MerchantId  string         `db:"merchant_id"`
    OrderCode   string         `db:"order_code"`
    ProductCode string         `db:"product_code"`
}

func main() {
    db, err := sqlx.Connect("mysql", "lpop:lpop5913kc@tcp(127.0.0.1:3306)/dc_lpop")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    tt := TTranslog2{}
    err = db.Get(&tt, `
        select id, yyyymmdd, affiliate_id, merchant_id, order_code, product_code 
        from ttranslog 
        where id = ?
    `, 2)
    //fmt.Printf("%+v\n", tt)

    fmt.Println(tt.Id, tt.YYYYMMDD, tt.AffiliateId.String, tt.MerchantId, tt.OrderCode, tt.ProductCode)
}
