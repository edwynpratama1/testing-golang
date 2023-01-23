package Configs

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func DatabaseOracle(mobileNumber string) string {

	db, err := sql.Open("godror", "ssldb/ssldb123@10.22.103.79:1521/ORCL")

	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer db.Close()

	fmt.Println("mobileNumber DB : ", mobileNumber)
	rows, err := db.Query("select rsa_user_name from app_user where login_name='" + mobileNumber + "'")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return ""
	}
	defer rows.Close()

	var cif string

	for rows.Next() {

		rows.Scan(&cif)
	}
	fmt.Printf("cif : %s\n", cif)

	return cif
}
