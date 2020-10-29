package main

import (
    "database/sql"
    "fmt"
    "github.com/samonzeweb/godb"
    "github.com/samonzeweb/godb/adapters/mysql"
    "log"
    "os"
)

type TTranslog struct {
    Id          int64  `db:"id"`
    YYYYMMDD    string `db:"yyyymmdd"`
    AffiliateId  string `db:"affiliate_id"`
    MerchantId  string `db:"merchant_id"`
    OrderCode   string `db:"order_code"`
    ProductCode string `db:"product_code"`
}

func main() {
    //db, err := sql.Open("mysql", "lpop:lpop5913kc@tcp(127.0.0.1:3306)/dc_lpop")
    db, err := godb.Open(mysql.Adapter, "lpop:lpop5913kc@tcp(127.0.0.1:3306)/dc_lpop")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    db.SetLogger(log.New(os.Stderr, "", 0))

    tt := TTranslog{}

    if err := db.Select(&tt).Where("id = ?", 1).Do(); err == sql.ErrNoRows {
        fmt.Println("TTranslog not found!!")
    }

    fmt.Println(tt.Id, tt.YYYYMMDD, tt.MerchantId, tt.OrderCode, tt.ProductCode)
}
