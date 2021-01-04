package sqlcipher

import (
	"encoding/json"
	"fmt"
	"testing"

	"gorm.io/gorm"
)

type Code struct {
	Id     int64
	Code   *string
	Status *string
}

func (Code) TableName() string {
	return "code"
}

func (c Code) String() string {
	bts, _ := json.Marshal(c)
	return string(bts)
}

func TestNormal(t *testing.T) {

	db, _ := gorm.Open(Open("test1.db"), nil)

	db.Exec(`CREATE TABLE IF NOT EXISTS "code" (
	  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	  "code" TEXT(64) NOT NULL DEFAULT '',
	  "status" integer(11) NOT NULL DEFAULT 1
	);`)

	db.Exec(`INSERT INTO code(code,status) VALUES('123', 1)`)

	codes := make([]*Code, 0)
	err := db.Find(&codes).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(codes)
}

func TestEncrypt(t *testing.T) {

	//key := "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"
	//dbname := fmt.Sprintf("test.db?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", key)
	/**
		SQLCipher 4 默认值,好多软件打不开db文件
		PRAGMA cipher_page_size = 4096;
		PRAGMA kdf_iter = 256000;
		PRAGMA cipher_hmac_algorithm = HMAC_SHA512;
		PRAGMA cipher_kdf_algorithm = PBKDF2_HMAC_SHA512;
	 */
	dbname := "users.db"
	db, _ := gorm.Open(Open(dbname), nil)
	db.Exec(`
		PRAGMA key = 123456;
		PRAGMA cipher_page_size = 1024;
		PRAGMA kdf_iter = 4000;
		PRAGMA cipher_hmac_algorithm = HMAC_SHA1;
		PRAGMA cipher_kdf_algorithm = PBKDF2_HMAC_SHA1;
		PRAGMA cipher_use_hmac = OFF;
	`)

	rows, err := db.Raw(`SELECT * FROM users;`).Rows()
	if err != nil {
		panic(err)
	}

	fmt.Println(rows.Columns())


}



