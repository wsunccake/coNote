# security and encryption

## csrf

Cross-Site Request Forgery

### gorilla/csrf

`common use-case: HTML form`

```go
// vulnerable.go
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var form = `
<html>
<head>
<title>Sign Up!</title>
</head>
<body>
<form method="POST" action="/signup/post" accept-charset="UTF-8">
<input type="text" name="name">
<input type="text" name="email">
<input type="submit" value="Sign up!">
</form>
</body>
</html>
`

var t = template.Must(template.New("signup_form.tmpl").Parse(form))

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", ShowSignupForm)
	r.HandleFunc("/signup/post", SubmitSignupForm)

	http.ListenAndServe(":8000", r)
}

func ShowSignupForm(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "signup_form.tmpl", nil)
}

func SubmitSignupForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "%v\n", r.PostForm)
}
```

```go
// prevention.go
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

var form = `
<html>
<head>
<title>Sign Up!</title>
</head>
<body>
<form method="POST" action="/signup/post" accept-charset="UTF-8">
<input type="text" name="name">
<input type="text" name="email">
<!--
    The default template tag used by the CSRF middleware .
    This will be replaced with a hidden <input> field containing the
    masked CSRF token.
-->
{{ .csrfField }}
<input type="submit" value="Sign up!">
</form>
</body>
</html>
`

var t = template.Must(template.New("signup_form.tmpl").Parse(form))

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", ShowSignupForm)
	r.HandleFunc("/signup/post", SubmitSignupForm)

	http.ListenAndServe(":8000",
		csrf.Protect([]byte("32-byte-long-auth-key"))(r))
}

func ShowSignupForm(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "signup_form.tmpl", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	})
}

func SubmitSignupForm(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
	fmt.Fprintf(w, "%v\n", r.PostForm)
}
```

```bash
#!/bin/bash

SERVER=127.0.0.1:8000
NAME=abc
EMAIL=abc@email.com

COOKIE=${COOKIE:=$(mktemp).cookie}
echo "COOKIE: $COOKIE"

uri_encode() {
    local code=$1
    echo ${code} | sed \
        -e 's| |%20|g' \
        -e 's|!|%21|g' \
        -e 's|#|%23|g' \
        -e 's|\$|%24|g' \
        -e 's|%|%25|g' \
        -e 's|&|%26|g' \
        -e "s|'|%27|g" \
        -e 's|(|%28|g' \
        -e 's|)|%29|g' \
        -e 's|*|%2A|g' \
        -e 's|+|%2B|g' \
        -e 's|,|%2C|g' \
        -e 's|/|%2F|g' \
        -e 's|:|%3A|g' \
        -e 's|;|%3B|g' \
        -e 's|=|%3D|g' \
        -e 's|?|%3F|g' \
        -e 's|@|%40|g' \
        -e 's|\[|%5B|g' \
        -e 's|]|%5D|g'
}

if [[ -z "${USER_TOKEN}" ]]; then
    USER_TOKEN=$(curl -c $COOKIE -s http://${SERVER}/signup \
    | grep gorilla.csrf.Token \
    | sed s/.\*value=\"// \
    | sed s/\".\*//)
fi
TOKEN=$(uri_encode $USER_TOKEN)
echo "USER_TOKEN: ${USER_TOKEN} => ${TOKEN}"

curl -b $COOKIE -s http://${SERVER}/signup/post -d "name=${NAME}&email=${EMAIL}"
curl -b $COOKIE -s http://${SERVER}/signup/post -d "name=${NAME}&email=${EMAIL}&gorilla.csrf.Token=${USER_TOKEN}"
curl -b $COOKIE -s http://${SERVER}/signup/post \
    -d "name=${NAME}" \
    -d "email=${EMAIL}" \
    -H "X-CSRF-Token: ${USER_TOKEN}"

cat $COOKIE
rm -fv $COOKIE
```

`api backend`

```go
// prevention.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

type User struct {
	Id   int
	Name string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/token", GetToken).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/user", GetUser).Methods("POST")

	CSRFMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.RequestHeader("X-CSRF-Token"))
	http.ListenAndServe(":8000", CSRFMiddleware(r))
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", csrf.Token(r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{1, "Daniel"}
	b, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(b)
}
```

```bash
#!/bin/bash

SERVER=127.0.0.1:8000

COOKIE=${COOKIE:=$(mktemp).cookie}
echo "COOKIE: $COOKIE"

curl -b $COOKIE -s http://${SERVER}/api/user/1

if [[ -z "${USER_TOKEN}" ]]; then
    USER_TOKEN=$(curl -c $COOKIE -s http://${SERVER}/token)
fi
echo "USER_TOKEN: ${USER_TOKEN}"

curl -X POST -b http://${SERVER}/api/user
curl -X POST -b $COOKIE http://${SERVER}/api/user
curl -X POST -H "X-CSRF-Token: ${USER_TOKEN}" http://${SERVER}/api/user
curl -X POST -b $COOKIE -H "X-CSRF-Token: ${USER_TOKEN}" http://${SERVER}/api/user

rm -fv $COOKIE
```

---

## xss

Cross-Site Scripting

```go
// vulnerable.go
package main

import (
	"io"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get("param"))
}

func main() {
	http.HandleFunc("/", server)
	http.ListenAndServe(":5000", nil)
}
```

```go
// prevention.go
package main

import (
	"io"
	"net/http"
	"text/template"
)

func server(w http.ResponseWriter, r *http.Request) {
	encodedParam := template.HTMLEscapeString(r.URL.Query().Get("param"))
	io.WriteString(w, encodedParam)
}

func main() {
	http.HandleFunc("/", server)
	http.ListenAndServe(":5000", nil)
}
```

```bash
linux:~ $ curl http://127.0.0.1:5000/?param=123
linux:~ $ curl http://127.0.0.1:5000/?param=<script>alert(1)</script>
```

```go
// vulnerable.go
package main

import (
	"net/http"
	"text/template"
)

func server(w http.ResponseWriter, r *http.Request) {
	error := r.URL.Query().Get("param")
	tmpl := template.New("error")
	tmpl, _ = tmpl.Parse(`{{define "T"}}{{.}}{{end}}`)
	tmpl.ExecuteTemplate(w, "T", error)
}

func main() {
	http.HandleFunc("/", server)
	http.ListenAndServe(":5000", nil)
}
```

```go
// prevention.go
package main

import (
	"html/template"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {
	error := r.URL.Query().Get("param")
	tmpl := template.New("error")
	tmpl, _ = tmpl.Parse(`{{define "T"}}{{.}}{{end}}`)
	tmpl.ExecuteTemplate(w, "T", error)
}

func main() {
	http.HandleFunc("/", server)
	http.ListenAndServe(":5000", nil)
}
```

[Cross-Site Scripting in Go Lang](https://knowledge-base.secureflag.com/vulnerabilities/cross_site_scripting/cross_site_scripting_go_lang.html)

---

## sql injection

### mysql

```sql
CREATE DATABASE foo;
USE foo;
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
```

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

func main() {
	var conn string
	var db *sql.DB

	conn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&readTimeout=%dms&writeTimeout=%dms&timeout=%dms", UserName, Password, Addr, Port, Database, 1000, 1000, 1000)
	db, _ = sql.Open("mysql", conn)
	defer db.Close()

	// userId := "1"             // normal
	userId := "123 OR 1=1;--" // vulnerable

	query := "SELECT * FROM userinfo WHERE uid = " + userId // vulnerable
	rows, _ := db.Query(query)                              // vulnerable

	// stmt, _ := db.Prepare("SELECT * FROM userinfo WHERE uid = ?") // prevention
	// rows, _ := stmt.Query(userId)                                 // prevention

	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		rows.Scan(&uid, &username, &department, &created)
		// checkErr(err)
		fmt.Printf("uid: %d, username: %s, department: %s, created: %v\n", uid, username, department, created)
	}
}
```

[SQL Injection in Go Lang](https://knowledge-base.secureflag.com/vulnerabilities/sql_injection/sql_injection_go_lang.html)

### gorm

[GORM Security](https://gorm.io/docs/security.html)

---

## ref

[SecureFlag Knowledge Base](https://knowledge-base.secureflag.com/)
[WSTG - Latest](https://owasp.org/www-project-web-security-testing-guide/latest/)
