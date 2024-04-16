# web server

---

## content

- [simple](#simple)

---

## simple

### hello

```bash
linux:~ $ mkdir demo
linux:~ $ cd demo
linux:~/demo $ go mod init example.com/demo

linux:~/demo $ vi hello.go

linux:~/demo $ go run .
linux:~/demo $ go run hello.go
```

```go
// hello.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go\n")
}

func sleep(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(5)

	fmt.Fprintf(w, "start to sleep %d s\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Fprintf(w, "end to sleep\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/sleep", sleep)
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
linux:~ $ curl http://127.0.0.1:9090

linux:~ $ curl http://127.0.0.1:9090/sleep
linux:~ $ seq 10 | xargs -i sh -c "echo {} && curl http://127.0.0.1:9090/sleep &"
```

### static

```go
// static.go
package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

```html
<!-- ./static/index.html -->
<!DOCTYPE html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, shrink-to-fit=no"
    />

    <!-- Bootstrap CSS -->
    <link
      rel="stylesheet"
      href="https://cdn.jsdelivr.net/npm/bootstrap@4.3.1/dist/css/bootstrap.min.css"
      integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"
      crossorigin="anonymous"
    />

    <title>Hello, world!</title>
  </head>
  <body>
    <h1>Hello, world!</h1>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script
      src="https://code.jquery.com/jquery-3.3.1.slim.min.js"
      integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo"
      crossorigin="anonymous"
    ></script>
    <script
      src="https://cdn.jsdelivr.net/npm/popper.js@1.14.7/dist/umd/popper.min.js"
      integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1"
      crossorigin="anonymous"
    ></script>
    <script
      src="https://cdn.jsdelivr.net/npm/bootstrap@4.3.1/dist/js/bootstrap.min.js"
      integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"
      crossorigin="anonymous"
    ></script>
  </body>
</html>
```

```bash
linux:~ $ curl http://127.0.0.1:9090
```

---

## form

### data

```go
// form1.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	username := "go"
	if _, err := r.Form["username"]; err {
		username = r.Form["username"][0]
	}

	fmt.Fprintf(w, "hello %s\n", username)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", sayHello)
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
linux:~ $ curl http://127.0.0.1:9090/hello
linux:~ $ curl http://127.0.0.1:9090/hello/
linux:~ $ curl http://127.0.0.1:9090/hello?username=happy
```

### login

```go
// form2.go
package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		crutime := time.Now().Unix()
		token := fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(crutime, 10))))

		t, err := template.ParseFiles("login.html")
		if err != nil {
			panic(err)
		}

		err = t.Execute(w, token)
		if err != nil {
			panic(err)
		}

	} else {
		r.ParseForm()
		username := template.HTMLEscapeString(r.Form.Get("username"))
		password := template.HTMLEscapeString(r.Form.Get("password"))
		token := r.Form.Get("token")
		if token != "" {
			// validate token
			if t, err := template.ParseFiles("success.html"); err == nil {
				if err := t.Execute(w, map[string]string{"UserName": username, "Password": password}); err != nil {
					panic(err)
				}
			}

		} else {
			// handle missing token error
			if t, err := template.ParseFiles("fail.html"); err == nil {
				t.Execute(w, nil)
			}
		}

		fmt.Println("username length:", len(username))
		fmt.Println("username:", username)
		fmt.Println("password:", password)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)

	server := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
```

```html
<!-- login.html -->
<html>
  <head>
    <title></title>
  </head>
  <body>
    <form action="/login" method="post">
      USERNAME:<input type="text" name="username" /> PASSWORD:<input
        type="password"
        name="password"
      />
      <input type="hidden" name="token" value="{{.}}" />
      <input type="submit" value="登入" />
    </form>
  </body>
</html>
```

```html
<!-- success.html -->
<html>
  <head>
    <title></title>
  </head>
  <body>
    <p style="color: green">username: {{.UserName}}</p>
    <p style="color: red">password: {{.Password}}</p>
  </body>
</html>
```

```html
<!-- fail.html -->
<html>
  <head>
    <title></title>
  </head>
  <body>
    <p style="color: red">fail to login</p>
    <a href="/login">login page</a>
  </body>
</html>
```

### upload and download

```go
// form2.go
package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 2 * 1024 * 1024 // 2 mb
var uploadPath = os.TempDir()         // /tmp

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, nil)
		return
	}

	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		log.Printf("could not parse multipart form: %v\n", err)
		renderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
		return
	}

	// parse and validate file and post parameters
	file, fileHeader, err := r.FormFile("uploadFile")
	if err != nil {
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// get and print out file size
	fileSize := fileHeader.Size
	log.Printf("file size (bytes): %v\n", fileSize)

	// validate file size
	if fileSize > maxUploadSize {
		renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}

	// check file type, `DetectContentType` only needs the first 512 bytes
	detectedFileType := http.DetectContentType(fileBytes)
	switch detectedFileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return
	}

	fileName := randToken(12)
	fileEndings, err := mime.ExtensionsByType(detectedFileType)
	if err != nil {
		renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
		return
	}

	newFileName := fileName + fileEndings[0]
	newPath := filepath.Join(uploadPath, newFileName)
	log.Printf("fileType: %s, file: %s\n", detectedFileType, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// write file bytes and check for errors
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("SUCCESS - use /files/%v to access the file", newFileName)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadFile)

	fs := http.FileServer(http.Dir(uploadPath))
	mux.Handle("/files/", http.StripPrefix("/files", fs))

	server := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
```

```html
<!-- upload.html -->
<html>
  <head>
    <title>Upload file</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="http://localhost:9090/upload"
      method="post"
    >
      <input type="file" name="uploadFile" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>
```
