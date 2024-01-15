# database - sql

---

## content

- [rmdb](#rmdb)
  - [sqlite](#sqlite)
  - [mysql](#mysql)
  - [postgesql](#postgesql)
  - [container](#container)
- [sql](#sql)
  - [DDL – Data Definition Language](#ddl-–-data-definition-language)
  - [DML – Data Manipulation Language](#dml-–-data-manipulation-language)
- [code](#code)
  - [sqlite example](#sqlite-example)
  - [mysql example](#mysql-example)
  - [postgresql example](#postgresql-example)

---

## rmdb

### sqlite

```bash
linux:~ # dnf install sqlite  # for rhel, feodra
linux:~ # apt install sqlite  # for debian, ubuntu

linux:~ $ sqlite3 db.sqlite3
```

```sql
sqlite> .help                     -- help
sqlite> .databases                -- show database
sqlite> .tables [<table_name>]    -- show table
sqlite> .schema [<table_name>]    -- show schema
sqlite> .show                     -- show config
sqlite> .quit
sqlite> .exit
```

### mysql

[MySQL Community Downloads](https://dev.mysql.com/downloads/)

```bash
linux:~ # dnf install community-mysql-server | mariadb  # for feodra
linux:~ # apt install mysql-server                      # for ubuntu

linux:~ # apt install lsb-release gnupg                     	# for debian
linux:~ # curl -LO https://dev.mysql.com/get/mysql-apt-config_0.8.29-1_all.deb
linux:~ # apt install ./mysql-apt-config_0.8.29-1_all.deb
linux:~ # apt update
linux:~ # apt install mysql-community-server

linux:~ # mysql -u root
```

```sql
mysql> CREATE DATABASE <database_name>;      -- create database
mysql> DROP DATABASE <database_name>;        -- delete database
mysql> SHOW DATABASES;                       -- list database
mysql> SHOW TABLES;                          -- list table
mysql> USE <database_name>                   -- use database
mysql> SHOW COLUMNS FROM <table_name>;
mysql> SHOW INDEX FROM <table_name>;
mysql> SHOW TABLE STATUS LIKE <table_name>;
```

### postgesql

```bash
linux:~ # dnf install postgresql-server   # for rhel, feodra
linux:~ # apt install postgresql          # for debian, ubuntu

linux:~ # psql
```

```sql
postgres=# CREATE DATABASE <database_name>;  -- create database
postgres=# DROP DATABASE <database_name>;    -- delete database
postgres=# \l                                -- list database
postgres=# \c <database_name>                -- use database
postgres=# SELECT current_database();        -- show current database
```

### container

- [mysql](https://hub.docker.com/_/mysql)
- [postgres](https://hub.docker.com/_/postgres)
- [redis](https://hub.docker.com/_/redis)

```bash
# mysql
linux:~ # docker run -d -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 -p 33060:33060 --name mysql mysql
linux:~ # docker exec -it mysql mysql -u root -p

# postgres
linux:~ # docker run -d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5432:5432 --name postgres postgres
linux:~ # docker exec -it postgres psql -U postgres

# redis
linux:~ # docker run -d -p 6379:6379 --name redis redis
linux:~ # docker exec -it redis redis-cli

# mongo
linux:~ # docker run -d -p 27017:27017 --name mongo mongo
linux:~ # docker exec -it mongo mongo
```

---

## sql

### DDL – Data Definition Language

```sql
-- sqlite
.databases

.tables

CREATE TABLE userinfo (
    uid INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(64) NULL,
    department VARCHAR(64) NULL,
    created DATE NULL
);

CREATE TABLE userdetail (
    uid INT(10) NULL,
    intro TEXT NULL,
    profile TEXT NULL,
    PRIMARY KEY (uid)
);

DROP TABLE userdetail;
DROP TABLE userinfo;
```

```sql
-- mysql
SHOW DATABASES;
CREATE DATABASE foo;
USE foo;
DROP DATABASE foo;

SHOW TABLES;

CREATE TABLE userinfo (
    uid INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(64) NULL,
    department VARCHAR(64) NULL,
    created DATE DEFAULT (CURRENT_DATE)
);

CREATE TABLE userdetail (
    uid INT(10) NOT NULL,
    intro TEXT NULL,
    profile TEXT NULL,
    PRIMARY KEY (uid)
);

DROP TABLE userdetail;
DROP TABLE userinfo;
```

```sql
-- postgre
\l
CREATE DATABASE foo;
\c foo
DROP DATABASE foo;

\d

CREATE TABLE userinfo (
    uid SERIAL PRIMARY KEY,
    username VARCHAR(64) NULL,
    department VARCHAR(64) NULL,
    created DATE DEFAULT (CURRENT_DATE)
);

CREATE TABLE userdetail (
    uid INT(10) NOT NULL,
    intro TEXT NULL,
    profile TEXT NULL,
    PRIMARY KEY (uid)
);

DROP TABLE userdetail;
DROP TABLE userinfo;
```

### DML – Data Manipulation Language

```sql
INSERT INTO userinfo (uid,username,department,created)
    VALUES (1, 'Paul', 'IT', '2024-01-01');
INSERT INTO userinfo (uid,username,department)
   VALUES (2, 'Allen', 'Account'),
   (3, 'Teddy', 'HR');

SELECT * FROM userinfo;
UPDATE userinfo SET created = '2024-01-01' WHERE uid = 1;
DELETE FROM userinfo WHERE uid = 4;
```

---

## code

### sqlite example

```go
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func createTable(db *sql.DB) {
	query1 := `CREATE TABLE IF NOT EXISTS userinfo(
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(64) NULL,
		department VARCHAR(64) NULL,
		created DATE DEFAULT CURRENT_DATE
	);`

	query2 := `CREATE TABLE IF NOT EXISTS userdetail (
		uid INT(10) NULL,
		intro TEXT NULL,
		profile TEXT NULL,
		PRIMARY KEY (uid)
	);`

	if _, err := db.Exec(query1); err != nil {
		panic(err)
	}

	if _, err := db.Exec(query2); err != nil {
		panic(err)
	}

}

func insertRow(db *sql.DB, username string, department string, created string) int64 {
	var stmt *sql.Stmt
	var err error
	var res sql.Result
	var id int64

	query := "INSERT INTO userinfo(username, department, created) values(?, ?, ?)"
	stmt, err = db.Prepare(query)
	checkErr(err)

	res, err = stmt.Exec(username, department, created)
	checkErr(err)

	id, err = res.LastInsertId()
	checkErr(err)

	// fmt.Println(id)
	return id
}

func updateRow(db *sql.DB, id int64) {
	var stmt *sql.Stmt
	var err error
	var res sql.Result

	query := "UPDATE userinfo SET username = ? WHERE uid = ?"

	stmt, err = db.Prepare(query)
	checkErr(err)

	res, err = stmt.Exec("Astaxie", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func queryRow(db *sql.DB) {
	var rows *sql.Rows
	var err error
	query := "SELECT * FROM userinfo"

	rows, err = db.Query(query)
	checkErr(err)

	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Printf("uid: %d, username: %s, department: %s, created: %v\n", uid, username, department, created)
	}
}

func deleteRow(db *sql.DB, id int64) {
	var stmt *sql.Stmt
	var err error
	var res sql.Result
	query := "DELETE FROM userinfo WHERE uid = ?"

	stmt, err = db.Prepare(query)
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)
}

func main() {
	var conn string
	var db *sql.DB
	var err error

	conn = "foo.db"
	db, err = sql.Open("sqlite3", conn)
	checkErr(err)
	defer db.Close()

	// create table
	createTable(db)

	// insert
	// id := insertRow(db)
	id := insertRow(db, "Paul", "IT", "2024-01-01")
	insertRow(db, "Allen", "Account", "")
	insertRow(db, "Teddy", "HR", time.Now().String())
	// query
	fmt.Println("--- after insert ---")
	queryRow(db)

	// update
	updateRow(db, id)

	// query
	fmt.Println("--- after update ---")
	queryRow(db)

	// delete
	deleteRow(db, id)

	// query
	fmt.Println("--- after delete ---")
	queryRow(db)
}
```

```bash
#!/bin/bash

test -f foo.db && rm -rf foo.db

sqlite3 foo.db << EOF
CREATE TABLE userinfo (
    uid INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(64) NULL,
    department VARCHAR(64) NULL,
    created DATE NULL
);
EOF

go run .

sqlite3 foo.db << EOF
.table
.schema
SELECT * FROM userinfo
EOF
```

### mysql example

```go
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

func createTable(db *sql.DB) {
	query1 := `CREATE TABLE IF NOT EXISTS userinfo(
		uid INT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(64) NULL,
		department VARCHAR(64) NULL,
		created DATE DEFAULT (CURRENT_DATE)
	);`

	query2 := `CREATE TABLE IF NOT EXISTS userdetail (
		uid INT(10),
		intro TEXT NULL,
		profile TEXT NULL,
		PRIMARY KEY (uid)
	);`

	if _, err := db.Exec(query1); err != nil {
		panic(err)
	}

	if _, err := db.Exec(query2); err != nil {
		panic(err)
	}

}

func insertRow(db *sql.DB, username string, department string, created string) int64 {
	var stmt *sql.Stmt
	var err error
	var res sql.Result
	var id int64

	query := "INSERT INTO userinfo(username, department, created) values(?, ?, ?)"
	stmt, err = db.Prepare(query)
	checkErr(err)

	if created == "" {
		created = time.Now().Format("2006-01-02")
	}

	res, err = stmt.Exec(username, department, created)
	checkErr(err)

	id, err = res.LastInsertId()
	checkErr(err)

	return id
}

func updateRow(db *sql.DB, id int64) {
	var stmt *sql.Stmt
	var err error
	var res sql.Result

	query := "UPDATE userinfo SET username = ? WHERE uid = ?"

	stmt, err = db.Prepare(query)
	checkErr(err)

	res, err = stmt.Exec("Astaxie", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func queryRow(db *sql.DB) {
	var rows *sql.Rows
	var err error
	query := "SELECT * FROM userinfo"

	rows, err = db.Query(query)
	checkErr(err)

	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Printf("uid: %d, username: %s, department: %s, created: %v\n", uid, username, department, created)
	}
}

func deleteRow(db *sql.DB, id int64) {
	var stmt *sql.Stmt
	var err error
	var res sql.Result
	query := "DELETE FROM userinfo WHERE uid = ?"

	stmt, err = db.Prepare(query)
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)
}
func main() {
	var conn string
	var db *sql.DB
	var err error

	conn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&readTimeout=%dms&writeTimeout=%dms&timeout=%dms", UserName, Password, Addr, Port, Database, 1000, 1000, 1000)

	db, err = sql.Open("mysql", conn)
	checkErr(err)
	defer db.Close()
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)

	// create table
	createTable(db)

	// insert
	id := insertRow(db, "Paul", "IT", "2024-01-01")
	insertRow(db, "Allen", "Account", "")
	insertRow(db, "Teddy", "HR", time.Now().Format("2006-01-02"))
	// query
	fmt.Println("--- after insert ---")
	queryRow(db)

	// update
	updateRow(db, id)

	// query
	fmt.Println("--- after update ---")
	queryRow(db)

	// delete
	deleteRow(db, id)

	// query
	fmt.Println("--- after delete ---")
	queryRow(db)
}
```

```bash
#!/bin/bash

username=root
password=password
db=foo

# create database
create_database() {
  docker exec -i mysql mysql -u$username -p$password << EOF
CREATE DATABASE $db;
EOF
}

# show table
show_table() {
  docker exec -i mysql mysql -u$username -p$password -D$db << EOF
SELECT * FROM userinfo;
EOF
}

# drop database
drop_database() {
  docker exec -i mysql mysql -u$username -p$password << EOF
DROP DATABASE $db;
EOF
}

###
### main
###

create_database

go run .
show_table

drop_database
```

### postgresql example

```go
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
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

func createTable(db *sql.DB) {
	query1 := `CREATE TABLE userinfo
	(
		uid serial NOT NULL,
		username character varying(100) NOT NULL,
		department character varying(500) NOT NULL,
		Created date,
		CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
	)
	WITH (OIDS=FALSE);`

	query2 := `CREATE TABLE userdetail
	(
		uid integer,
		intro character varying(100),
		profile character varying(100)
	)
	WITH(OIDS=FALSE);`

	if _, err := db.Exec(query1); err != nil {
		panic(err)
	}

	if _, err := db.Exec(query2); err != nil {
		panic(err)
	}

}

func insertRow(db *sql.DB, username string, department string, created string) int64 {
	var stmt *sql.Stmt
	var err error
	// var res sql.Result
	var id int64

	query := "INSERT INTO userinfo(username, department, created) VALUES($1, $2, $3) RETURNING uid"
	stmt, err = db.Prepare(query)
	checkErr(err)

	if created == "" {
		created = time.Now().Format("2006-01-02")
	}

	_, err = stmt.Exec(username, department, created)
	checkErr(err)

	// var id int64
	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", username, department, created).Scan(&id)
	checkErr(err)

	return id
}

func updateRow(db *sql.DB, id int64) {
	var stmt *sql.Stmt
	var err error
	var res sql.Result

	query := "UPDATE userinfo SET username = $1 WHERE uid = $2"

	stmt, err = db.Prepare(query)
	checkErr(err)

	res, err = stmt.Exec("Astaxie", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func queryRow(db *sql.DB) {
	var rows *sql.Rows
	var err error
	query := "SELECT * FROM userinfo"

	rows, err = db.Query(query)
	checkErr(err)

	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Printf("uid: %d, username: %s, department: %s, created: %v\n", uid, username, department, created)
	}
}

func deleteRow(db *sql.DB, id int64) {
	var stmt *sql.Stmt
	var err error
	var res sql.Result
	query := "DELETE FROM userinfo WHERE uid = $1"

	stmt, err = db.Prepare(query)
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)
}

func main() {
	var conn string
	var db *sql.DB
	var err error

	conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Addr, Port, UserName, Password, Database)

	db, err = sql.Open("postgres", conn)
	checkErr(err)
	defer db.Close()
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)

	// create table
	createTable(db)

	// insert
	id := insertRow(db, "Paul", "IT", "2024-01-01")
	insertRow(db, "Allen", "Account", "")
	insertRow(db, "Teddy", "HR", time.Now().Format("2006-01-02"))
	// query
	fmt.Println("--- after insert ---")
	queryRow(db)

	// update
	updateRow(db, id)

	// query
	fmt.Println("--- after update ---")
	queryRow(db)

	// delete
	deleteRow(db, id)

	// query
	fmt.Println("--- after delete ---")
	queryRow(db)
}
```

```bash
#!/bin/bash

username=postgres
db=foo

create_database() {
    docker exec -i postgres psql -U $username << EOF
CREATE DATABASE $db;
EOF
}

show_table() {
    docker exec -i postgres psql -U $username -d $db << EOF
SELECT * FROM userinfo;
EOF
}

drop_database() {
    docker exec -i postgres psql -U $username << EOF
DROP DATABASE $db;
EOF
}

# show table
show_table() {
  docker exec -i postgres psql -U $username -d $db << EOF
\d
\d userinfo
SELECT * FROM userinfo;
EOF
}

###
### main
###

create_database

go run .

show_table
drop_database
```
