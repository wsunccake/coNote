# database - orm

---

## content

- [gorm](#gorm)
  - [gorm - sqlite](#gorm---sqlite)
  - [gorm - mysql](#gorm---mysql)
  - [gorm - postgresql](#gorm---postgresql)
- [beego](#beego)
  - [beego - sqlite](#beego---sqlite)
  - [beego - mysql](#beego---mysql)
  - [beego - postgresql](#beego---postgresql)

---

## gorm

### gorm - sqlite

```go
package main
import (
	"database/sql"
	"fmt"
	"time"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
type Userinfo struct {
	Uid        uint      `gorm:"column:uid;primaryKey;autoIncrement"`
	Username   string    `gorm:"column:username"`
	Department string    `gorm:"column:department"`
	Created    time.Time `gorm:"column:created"`
}
type Tabler interface {
	TableName() string
}
func (Userinfo) TableName() string {
	return "userinfo"
}
func main() {
	var conn string
	var db *gorm.DB
	var err error
	var sqlDb *sql.DB
	conn = "foo.db"
	db, err = gorm.Open(sqlite.Open(conn), &gorm.Config{})
	checkErr(err)
	sqlDb, err = db.DB()
	checkErr(err)
	defer sqlDb.Close()
	// create
	db.Create(&Userinfo{Username: "jinzhu", Department: "Sales", Created: time.Now()})
	// read
	var userinfo1, userinfo2, userinfo3 Userinfo
	result := db.First(&userinfo1, 1)
	checkErr(result.Error)
	fmt.Printf("userinfo: %+v\n", userinfo1)
	db.First(&userinfo2, "uid = ?", 3)
	fmt.Printf("userinfo: %+v\n", userinfo2)
	db.Last(&userinfo3)
	fmt.Printf("userinfo: %+v\n", userinfo3)
	var userinfos []Userinfo
	db.Find(&userinfos)
	fmt.Printf("userinfos: %+v\n", userinfos)
	for _, v := range userinfos {
		fmt.Printf("userinfos: %+v\n", v)
	}
	// update
	db.Model(&userinfo3).Update("department", "Marketing")
	db.Model(&userinfo3).Updates(Userinfo{Username: "George", Department: "Finance"})
	// delete
	db.Delete(&userinfo3)
}
```

### gorm - mysql

```go
package main
import (
	"database/sql"
	"fmt"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
const (
	UserName     string = "root"
	Password     string = "password"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "foo"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Userinfo struct {
	Uid        uint      `gorm:"column:uid;primaryKey;autoIncrement"`
	Username   string    `gorm:"column:username"`
	Department string    `gorm:"column:department"`
	Created    time.Time `gorm:"column:created"`
}

type Tabler interface {
	TableName() string
}

func (Userinfo) TableName() string {
	return "userinfo"
}

func main() {
	var conn string
	var db *gorm.DB
	var err error
	var sqlDb *sql.DB
	conn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&readTimeout=%dms&writeTimeout=%dms&timeout=%dms", UserName, Password, Addr, Port, Database, 1000, 1000, 1000)
	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{})
	checkErr(err)
	sqlDb, err = db.DB()
	checkErr(err)
	defer sqlDb.Close()
	sqlDb.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	sqlDb.SetMaxOpenConns(MaxOpenConns)
	sqlDb.SetMaxIdleConns(MaxIdleConns)
	// create
	db.Create(&Userinfo{Username: "jinzhu", Department: "Sales", Created: time.Now()})
	// read
	var userinfo1, userinfo2, userinfo3 Userinfo
	result := db.First(&userinfo1, 1)
	checkErr(result.Error)
	fmt.Printf("userinfo: %+v\n", userinfo1)
	db.First(&userinfo2, "uid = ?", 3)
	fmt.Printf("userinfo: %+v\n", userinfo2)
	db.Last(&userinfo3)
	fmt.Printf("userinfo: %+v\n", userinfo3)
	var userinfos []Userinfo
	db.Find(&userinfos)
	fmt.Printf("userinfos: %+v\n", userinfos)
	for _, v := range userinfos {
		fmt.Printf("userinfos: %+v\n", v)
	}
	// update
	db.Model(&userinfo3).Update("department", "Marketing")
	db.Model(&userinfo3).Updates(Userinfo{Username: "George", Department: "Finance"})
	// delete
	db.Delete(&userinfo3)
}
```

### gorm - postgresql

```go
package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	UserName     string = "postgres"
	Password     string = "password"
	Addr         string = "127.0.0.1"
	Port         int    = 5432
	Database     string = "foo"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Userinfo struct {
	Uid        uint      `gorm:"column:uid;primaryKey;autoIncrement"`
	Username   string    `gorm:"column:username"`
	Department string    `gorm:"column:department"`
	Created    time.Time `gorm:"column:created"`
}

type Tabler interface {
	TableName() string
}

func (Userinfo) TableName() string {
	return "userinfo"
}

func main() {
	var conn string
	var db *gorm.DB
	var err error
	var sqlDb *sql.DB

	conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Addr, Port, UserName, Password, Database)
	db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	checkErr(err)

	sqlDb, err = db.DB()
	checkErr(err)
	defer sqlDb.Close()
	sqlDb.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	sqlDb.SetMaxOpenConns(MaxOpenConns)
	sqlDb.SetMaxIdleConns(MaxIdleConns)

	// create
	db.Create(&Userinfo{Username: "jinzhu", Department: "Sales", Created: time.Now()})

	// read
	var userinfo1, userinfo2, userinfo3 Userinfo
	result := db.First(&userinfo1, 1)
	checkErr(result.Error)
	fmt.Printf("userinfo: %+v\n", userinfo1)
	db.First(&userinfo2, "uid = ?", 3)
	fmt.Printf("userinfo: %+v\n", userinfo2)
	db.Last(&userinfo3)
	fmt.Printf("userinfo: %+v\n", userinfo3)
	var userinfos []Userinfo
	db.Find(&userinfos)
	fmt.Printf("userinfos: %+v\n", userinfos)
	for _, v := range userinfos {
		fmt.Printf("userinfos: %+v\n", v)
	}

	// update
	db.Model(&userinfo3).Update("department", "Marketing")
	db.Model(&userinfo3).Updates(Userinfo{Username: "George", Department: "Finance"})

	// delete
	db.Delete(&userinfo3)
}
```

---

## beego

### beego - sqlite

```go
package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type Userinfo struct {
	Uid        uint      `orm:"pk;auto;column(uid)"`
	Username   string    `orm:"null;size(64);column(username)"`
	Department string    `orm:"null;size(64);column(department)"`
	Created    time.Time `orm:"column(created)"`
}

func init() {
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	conn := "foo.db"

	orm.RegisterDataBase("default", "sqlite3", conn, MaxOpenConns)
	orm.RegisterModel(new(Userinfo))
	//orm.RegisterModel(new(Userinfo), new(Profile), new(Post))

	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	userinfo := Userinfo{
		Username:   "slene",
		Department: "RD",
		Created:    time.Now(),
	}

	// insert
	id, err := o.Insert(&userinfo)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	userinfo.Username = "astaxie"
	num, err := o.Update(&userinfo)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read
	u := Userinfo{Uid: userinfo.Uid}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
```

### beego - mysql

```go
package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Userinfo struct {
	Uid        uint      `orm:"pk;auto;column(uid)"`
	Username   string    `orm:"null;size(64);column(username)"`
	Department string    `orm:"null;size(64);column(department)"`
	Created    time.Time `orm:"column(created)"`
}

const (
	UserName     string = "root"
	Password     string = "password"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "foo"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&readTimeout=%dms&writeTimeout=%dms&timeout=%dms", UserName, Password, Addr, Port, Database, 1000, 1000, 1000)

	orm.RegisterDataBase("default", "mysql", conn, MaxOpenConns)
	orm.RegisterModel(new(Userinfo))
	//orm.RegisterModel(new(Userinfo), new(Profile), new(Post))

	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	userinfo := Userinfo{
		Username:   "slene",
		Department: "RD",
		Created:    time.Now(),
	}

	// insert
	id, err := o.Insert(&userinfo)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	userinfo.Username = "astaxie"
	num, err := o.Update(&userinfo)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read
	u := Userinfo{Uid: userinfo.Uid}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
```

### beego - postgresql

```go
package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

type Userinfo struct {
	Uid        uint      `orm:"pk;auto;column(uid)"`
	Username   string    `orm:"null;size(64);column(username)"`
	Department string    `orm:"null;size(64);column(department)"`
	Created    time.Time `orm:"column(created)"`
}

const (
	UserName     string = "postgres"
	Password     string = "password"
	Addr         string = "127.0.0.1"
	Port         int    = 5432
	Database     string = "foo"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func init() {
	orm.RegisterDriver("postgres", orm.DR_Postgres)

	conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Addr, Port, UserName, Password, Database)	orm.RegisterDataBase("default", "postgres", "user=postgres password=zxxx dbname=test host=127.0.0.1 port=5432 sslmode=disable")
	orm.RegisterDataBase("default", "postgres", conn, MaxOpenConns)
	orm.RegisterModel(new(Userinfo))
	//orm.RegisterModel(new(Userinfo), new(Profile), new(Post))

	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	userinfo := Userinfo{
		Username:   "slene",
		Department: "RD",
		Created:    time.Now(),
	}

	// insert
	id, err := o.Insert(&userinfo)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	userinfo.Username = "astaxie"
	num, err := o.Update(&userinfo)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read
	u := Userinfo{Uid: userinfo.Uid}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
```
