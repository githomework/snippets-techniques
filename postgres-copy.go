import (
	"database/sql"
	"github.com/lib/pq"
)
	

_, err := db.Exec("set search_path to xxxx")
		if err != nil {
			log.Panicln(err)
		}
txn, err := db.Begin()
		if err != nil {
			log.Panicln(err)
		}
		defer txn.Commit()

		statement, err := txn.Prepare(pq.CopyIn("products", "product", "uuid", "product_long", "description", "hex", "uom"))
		if err != nil {
			log.Fatalln(err)
		}
		defer statement.Close()
		for _, v := range rows {
			product := ""
			if len(v.productLong.String) >= 8 {
				product = v.productLong.String[len(v.productLong.String)-8:]
			} else {
				product = v.productLong.String
			}

			_, err := statement.Exec(product, v.uuidC22.String, v.productLong.String, v.description.String, text.EWMDecodeUUID(v.uuidC22.String), v.uom.String)
			if err != nil {
				log.Println(product, err)
			}
		}
