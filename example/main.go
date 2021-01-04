package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mutecomm/go-sqlcipher/v4"
)

func main() {
	dsn := "E:\\project\\mine\\go\\gorm-sqlcipher\\users.db"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		fmt.Println("sql.Open()", err)
		return
	}
	defer db.Close()

	s := `
	PRAGMA key = 123456;
	PRAGMA cipher_page_size = 1024;
	PRAGMA kdf_iter = 4000;
	PRAGMA cipher_hmac_algorithm = HMAC_SHA1;
	PRAGMA cipher_kdf_algorithm = PBKDF2_HMAC_SHA1;
	PRAGMA cipher_use_hmac = OFF;
	`

	_, err = db.Exec(s)
	if err != nil {
		fmt.Println(s, err)
		return
	}

	c := "CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY, `name` char, `password` chart, UNIQUE(`name`));"
	_, err = db.Exec(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	d := "INSERT INTO `users` (name, password) values('xeodou1', 123456);"
	_, err = db.Exec(d)
	if err != nil {
		fmt.Println(err)
		return
	}

	e := "select name, password from users where name='xeodou';"
	rows, err := db.Query(e)
	if err != nil {
		fmt.Println("db.Query() ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var password string
		rows.Scan(&name, &password)
		fmt.Print("{\"name\":\"" + name + "\", \"password\": \"" + password + "\"}")
	}
	rows.Close()
}
