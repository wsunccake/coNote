# cookie and session

---

## content

- [cookie](#cookie)
  - [chrome](#chrome)
  - [edge](#edge)
  - [safari](#safari)
  - [firefox](#firefox)
- [session](#session)
  - [gorilla/sessions](#gorillasessions)

---

## cookie

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:    "username",
		Value:   "johndoe",
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, &c)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s!", c.Value)
}

func deleteCookie(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:   "username",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cookie/set", setCookie)
	mux.HandleFunc("/cookie/get", getCookie)
	mux.HandleFunc("/cookie/del", deleteCookie)
	server := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
```

```bash
linux:~ $ curl -b "username=alex" http://127.0.0.1:9090/cookie/get

linux:~ $ curl -c cookie.txt http://127.0.0.1:9090/cookie/set
linux:~ $ cat cookie.txt
```

### chrome

```bash
# macosx
mac:~ $ sqlite3 ~/Library/Application Support/Google/Chrome/Default/Cookies

# linux
linux:~ $ sqlite3 ~/.config/google-chrome/Default/Cookies

sqlite> SELECT * FROM cookies WHERE name = "ClientId";
sqlite> SELECT * FROM cookies WHERE host_key LIKE '%google%' LIMIT 3;
```

---

## session

### gorilla/sessions

```go
// sessions.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
```

```bash
linux:~ $ curl http://localhost:8080/secret

# with session
linux:~ $ curl -I http://localhost:8080/login
linux:~ $ curl -b "cookie-name=...=" http://localhost:8080/secret

# with session
linux:~ $ curl -c session http://localhost:8080/login
linux:~ $ curl -b session http://localhost:8080/secret
```

(sessions)[https://github.com/gorilla/sessions]
