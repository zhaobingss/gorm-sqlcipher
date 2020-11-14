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

	key := "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"
	dbname := fmt.Sprintf("test.db?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", key)

	db, _ := gorm.Open(Open(dbname), nil)

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



