# GORM Sqlite Driver

![CI](https://github.com/go-gorm/sqlite/workflows/CI/badge.svg)

## USAGE

### 主要用来适配gorm和go-sqlcipher的

```go
import (
    "github.com/zhaobingss/sqlite"
    "gorm.io/gorm"
    "fmt"
    "testing"
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

func Test(t *testing.T) {
    key := "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"
    dbname := fmt.Sprintf("test.db?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", key)
    
    db, _ := gorm.Open(sqlite.Open(dbname), nil)
    
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

```


