package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/kisielk/sqlstruct"
    "log"
)

type TTranslog1 struct {
    Id          int64      `sql:"id"`
    YYYYMMDD    string     `sql:"yyyymmdd"`
    AffiliateId sql.NullString `sql:"affiliate_id"`
    MerchantId  string     `sql:"merchant_id"`
    OrderCode   string     `sql:"order_code"`
    ProductCode string     `sql:"product_code"`
}

func main() {
    //db, err := sql.Open("mysql", "lpop:lpop5913kc@tcp(127.0.0.1:3306)/dc_lpop")
    db, err := sql.Open("mysql", "lpop:lpop5913kc@tcp(127.0.0.1:3306)/dc_lpop")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query(fmt.Sprintf(`
        select %s 
        from   ttranslog
        where  id = ?
    `, sqlstruct.Columns(TTranslog1{})), 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var tt TTranslog1
        if err := sqlstruct.Scan(&tt, rows); err != nil {
            log.Fatal(err)
        }

        //fmt.Printf("%+v\n", tt)
        fmt.Println(tt.Id, tt.YYYYMMDD, tt.AffiliateId.String, tt.MerchantId, tt.OrderCode, tt.ProductCode)
    }
}
