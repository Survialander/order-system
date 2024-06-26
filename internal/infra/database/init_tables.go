package database

import "database/sql"

func InitDbTables(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS orders")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	if err != nil {
		return err
	}

	return nil
}
