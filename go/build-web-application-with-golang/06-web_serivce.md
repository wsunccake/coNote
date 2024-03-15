# web service

---

## content

- [socket](#socket)
  - [c](#c)
  - [cpp](#cpp)
  - [python](#python)
  - [go](#go)
  - [command](#command)
- [websocket](#websocket)
  - [net/websocket](#netwebsocket)
  - [gorilla/websocket](#gorillawebsocket)
- [restful](#restful)
  - [julienschmidt/httprouter](#julienschmidthttprouter)
  - [gorilla/mux](#gorillamux)
- [rpc](#rpc)
  - [http](#http)
  - [tcp](#tcp)
  - [json](#json)
- [protobuf](#protobuf)
- [grpc](#grpc)

---

## socket

UDS / Unix Domain Socket

### c

```c
// server.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/un.h>

int main() {
    int server_socket;
    int client_socket;
    struct sockaddr_un server_addr;
    struct sockaddr_un client_addr;

    int result;

    server_socket = socket(AF_UNIX, SOCK_STREAM, 0);

    server_addr.sun_family = AF_UNIX;
    strcpy(server_addr.sun_path, "/tmp/socket");

    int slen = sizeof(server_addr);

    bind(server_socket, (struct sockaddr *) &server_addr, slen);

    listen(server_socket, 5);

    while(1){
        char ch;
        int clen = sizeof(client_addr);
        client_socket = accept(server_socket, (struct sockaddr *) &client_addr, &clen);
        read(client_socket, &ch, 1);
        printf("\nServer: I recieved %c from client!\n", ch);
        ch++;
        write(client_socket, &ch, 1);
        close(client_socket);
    }

    exit(0);
}
```

```c
// client.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/un.h>

int main() {
    int server_socket;
    struct sockaddr_un server_addr;
    int connection_result;

    char ch='C';

    server_socket = socket(AF_UNIX, SOCK_STREAM, 0);

    server_addr.sun_family = AF_UNIX;
    strcpy(server_addr.sun_path, "/tmp/socket");

    connection_result = connect(server_socket, (struct sockaddr *)&server_addr, sizeof(server_addr));

    if (connection_result == -1) {
        perror("Error:");
        exit(1);
    }

    write(server_socket, &ch, 1);
    read(server_socket, &ch, 1);
    printf("Client: I recieved %c from server!\n", ch);
    close(server_socket);
    exit(0);
}
```

### cpp

```cpp
// server.cpp
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <sys/types.h>

static const char *socket_path = "/tmp/socket";
static const unsigned int nIncomingConnections = 5;

int main()
{
	// create server side
	int s = 0;
	int s2 = 0;
	struct sockaddr_un local, remote;
	int len = 0;

	s = socket(AF_UNIX, SOCK_STREAM, 0);
	if (-1 == s)
	{
		printf("Error on socket() call \n");
		return 1;
	}

	local.sun_family = AF_UNIX;
	strcpy(local.sun_path, socket_path);
	unlink(local.sun_path);
	len = strlen(local.sun_path) + sizeof(local.sun_family);
	if (bind(s, (struct sockaddr *)&local, len) != 0)
	{
		printf("Error on binding socket \n");
		return 1;
	}

	if (listen(s, nIncomingConnections) != 0)
	{
		printf("Error on listen call \n");
	}

	bool bWaiting = true;
	while (bWaiting)
	{
		unsigned int sock_len = 0;
		printf("Waiting for connection.... \n");
		if ((s2 = accept(s, (struct sockaddr *)&remote, &sock_len)) == -1)
		{
			printf("Error on accept() call \n");
			return 1;
		}

		printf("Server connected \n");

		int data_recv = 0;
		char recv_buf[100];
		char send_buf[200];
		do
		{
			memset(recv_buf, 0, 100 * sizeof(char));
			memset(send_buf, 0, 200 * sizeof(char));
			data_recv = recv(s2, recv_buf, 100, 0);
			if (data_recv > 0)
			{
				printf("Data received: %d : %s \n", data_recv, recv_buf);
				strcpy(send_buf, "Got message: ");
				strcat(send_buf, recv_buf);

				if (strstr(recv_buf, "quit") != 0)
				{
					printf("Exit command received -> quitting \n");
					bWaiting = false;
					break;
				}

				if (send(s2, send_buf, strlen(send_buf) * sizeof(char), 0) == -1)
				{
					printf("Error on send() call \n");
				}
			}
			else
			{
				printf("Error on recv() call \n");
			}
		} while (data_recv > 0);

		close(s2);
	}

	return 0;
}
```

```cpp
// client.cpp
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <sys/types.h>

static const char *socket_path = "/tmp/socket";
static const unsigned int s_recv_len = 200;
static const unsigned int s_send_len = 100;

int main()
{
	int sock = 0;
	int data_len = 0;
	struct sockaddr_un remote;
	char recv_msg[s_recv_len];
	char send_msg[s_send_len];

	memset(recv_msg, 0, s_recv_len * sizeof(char));
	memset(send_msg, 0, s_send_len * sizeof(char));

	if ((sock = socket(AF_UNIX, SOCK_STREAM, 0)) == -1)
	{
		printf("Client: Error on socket() call \n");
		return 1;
	}

	remote.sun_family = AF_UNIX;
	strcpy(remote.sun_path, socket_path);
	data_len = strlen(remote.sun_path) + sizeof(remote.sun_family);

	printf("Client: Trying to connect... \n");
	if (connect(sock, (struct sockaddr *)&remote, data_len) == -1)
	{
		printf("Client: Error on connect call \n");
		return 1;
	}

	printf("Client: Connected \n");

	while (printf(">"), fgets(send_msg, s_send_len, stdin), !feof(stdin))
	{
		if (send(sock, send_msg, strlen(send_msg) * sizeof(char), 0) == -1)
		{
			printf("Client: Error on send() call \n");
		}
		memset(send_msg, 0, s_send_len * sizeof(char));
		memset(recv_msg, 0, s_recv_len * sizeof(char));

		if ((data_len = recv(sock, recv_msg, s_recv_len, 0)) > 0)
		{
			printf("Client: Data received: %s \n", recv_msg);
		}
		else
		{
			if (data_len < 0)
			{
				printf("Client: Error on recv() call \n");
			}
			else
			{
				printf("Client: Server socket closed \n");
				close(sock);
				break;
			}
		}
	}

	printf("Client: bye! \n");

	return 0;
}
```

### python

```py
# server.py
import socket

class SocketServer:
    def __init__(self):
        # tcp socket
        # server_address = ('127.0.0.1', 9999)
        # socket_family = socket.AF_INET
        # socket_type = socket.SOCK_STREAM

        # unix domain sockets
        server_address = '/tmp/socket'
        socket_family = socket.AF_UNIX
        socket_type = socket.SOCK_STREAM

        self.sock = socket.socket(socket_family, socket_type)
        self.sock.bind(server_address)
        self.sock.listen(1)
        print(f"listening on '{server_address}'.")
        pass

    def wait_and_deal_client_connect(self):
        while True:
            connection, client_address = self.sock.accept()
            data = connection.recv(1024)
            print(f"recv data from client '{client_address}': {data.decode()}")
            connection.sendall("hello client".encode())

    def __del__(self):
        self.sock.close()

if __name__ == "__main__":
    socket_server_obj = SocketServer()
    socket_server_obj.wait_and_deal_client_connect()
```

```py
# client.py
import socket

class SocketClient:
    def __init__(self):
        pass

    def connect_to_server(self):
        # tcp socket
        # server_address = ('127.0.0.1', 9999)
        # socket_family = socket.AF_INET
        # socket_type = socket.SOCK_STREAM

        # unix domain socket
        server_address = '/tmp/socket'
        socket_family = socket.AF_UNIX
        socket_type = socket.SOCK_STREAM

        sock = socket.socket(socket_family, socket_type)
        sock.connect(server_address)
        sock.sendall("hello server".encode())
        data = sock.recv(1024)
        print(f"recv data from server '{server_address}': {data.decode()}")
        sock.close()

if __name__ == "__main__":
    socket_client_obj = SocketClient()
    socket_client_obj.connect_to_server()
```

### go

```go
// server.go
package main

import (
	"io"
	"log"
	"net"
	"os"
)

const SockAddr = "/tmp/socket"

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	io.Copy(c, c)
	c.Close()
}

func main() {
	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(conn)
	}
}
```

```go
// client.go
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf[:])
	if err != nil {
		return
	}
	println("Client got:", string(buf[0:n]))
}

func main() {
	c, err := net.Dial("unix", "/tmp/socket")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go reader(c)
	_, err = c.Write([]byte("hi"))
	if err != nil {
		log.Fatal("write error:", err)
	}
	reader(c)
	time.Sleep(100 * time.Millisecond)
}
```

### netcat

```bash
# send data to socket
linux:~ # nc -U /tmp/socket
```

```bash
# receive / server
server:~ # nc -l -p 1234

# send / client
client:~ # nc <server ip> 1234
```

### socat

```bash
# read local file
linux:~ $ socat - /etc/sysctl.conf

# write local file
linux:~ $ echo "Hello" | socat - /tmp/hello.txt
```

---

## websocket

### net/websocket

```go
// websocket.go
package main

import (
	"fmt"
	"html/template"

	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("receive failed:", err)
			break
		}

		fmt.Println("reveived from client: " + reply)
		msg := reply
		fmt.Println("send to client:" + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("send failed:", err)
			break
		}
	}
}

func web(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("websockets.html")
		t.Execute(w, nil)
	} else {
		fmt.Println(r.PostFormValue("username"))
		fmt.Println(r.PostFormValue("password"))
	}

}

func main() {
	http.Handle("/websocket", websocket.Handler(Echo))
	http.HandleFunc("/web", web)

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
```

```html
<!-- websocket.html -->
<!DOCTYPE html>

<html>
  <head>
    <meta charset="utf-8" />
    <title>go websocket</title>
  </head>

  <body>
    <script type="text/javascript">
      var sock = null;

      var wsuri = "ws://127.0.0.1:1234/websocket";

      window.onload = function () {
        console.log("onload");

        sock = new WebSocket(wsuri);

        sock.onopen = function () {
          console.log("connected to " + wsuri);
          output.innerHTML += "Status: Connected\n";
        };

        sock.onclose = function (e) {
          console.log("connection closed (" + e.code + ")");
          output.innerHTML += "connection closed: " + e.data + "\n";
        };

        sock.onmessage = function (e) {
          console.log("message received: " + e.data);
          output.innerHTML += "message received: " + e.data + "\n";
        };
      };

      function send() {
        var msg = document.getElementById("message").value;
        sock.send(msg);
      }
    </script>

    <h1>WebSocket Echo Test</h1>

    <form>
      <p>Message: <input id="message" type="text" value="Hello, world!" /></p>
    </form>

    <button onclick="send();">Send Message</button>
    <pre id="output"></pre>
  </body>
</html>
```

### gorilla/websocket

```go
// websockets.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
```

```html
<!-- websockets.html -->
<input id="input" type="text" />
<button onclick="send()">Send</button>
<pre id="output"></pre>
<script>
  var input = document.getElementById("input");
  var output = document.getElementById("output");
  //   var socket = new WebSocket("ws://127.0.0.1:8080/echo");
  var socket = new WebSocket("ws://" + document.location.host + "/echo");

  socket.onopen = function () {
    output.innerHTML += "Status: Connected\n";
  };

  socket.onmessage = function (e) {
    output.innerHTML += "Server: " + e.data + "\n";
  };

  function send() {
    socket.send(input.value);
    input.value = "";
  }
</script>
```

---

## restful

### julienschmidt/httprouter

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func getuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are modify user %s", uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}

func adduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// uid := r.FormValue("uid")
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are add user %s", uid)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	router.GET("/user/:uid", getuser)
	router.POST("/adduser/:uid", adduser)
	router.DELETE("/deluser/:uid", deleteuser)
	router.PUT("/moduser/:uid", modifyuser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
```

```bash
linux:~ $ curl -X GET http://127.0.0.1:8080
linux:~ $ curl -X GET http://127.0.0.1:8080/hello/go
linux:~ $ curl -X POST http://127.0.0.1:8080/adduser/go
linux:~ $ curl -X DELETE http://127.0.0.1:8080/deluser/go
linux:~ $ curl -X PUT http://127.0.0.1:8080/moduser/go
```

### gorilla/mux

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, post := range posts {
		if post.ID == id {
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post.ID = len(posts) + 1
	posts = append(posts, post)
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedPost Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, post := range posts {
		if post.ID == id {
			posts[i].Title = updatedPost.Title
			posts[i].Body = updatedPost.Body
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	// Seed data
	posts = append(posts, Post{ID: 1, Title: "My first post", Body: "This is the content of my first post."})

	router.HandleFunc("/api/posts", getPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/api/posts", createPost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
```

```bash
linux:~ $ curl -X GET http://127.0.0.1:8000/api/posts
linux:~ $ curl -X GET http://127.0.0.1:8000/api/posts/1
linux:~ $ curl -X POST -d '{"title": "abc", "body": "xyz"}' http://127.0.0.1:8000/api/posts
linux:~ $ curl -X PUT -d '{"id":2,"title": "ABC", "body": "xyz"}' http://127.0.0.1:8000/api/posts
linux:~ $ curl -X DELETE http://127.0.0.1:8000/api/2
```

---

## rpc

### http

```go
// http_server.go
package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
```

```go
// http_client.go
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
```

```bash
# for server
linux:~ $ ./http_server

# for client
linux:~ $ ./http_client 127.0.0.1
```

### tcp

```go
// tcp_server.go
package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {=
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
```

```go
// tcp_client.go
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}
	service := os.Args[1]

	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
```

```bash
# for server
linux:~ $ ./tcp_server

# for client
linux:~ $ ./tcp_client 127.0.0.1:1234
```

### json

```go
// json_server.go
package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
```

```go
// json_client.go
package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		log.Fatal(1)
	}
	service := os.Args[1]

	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
```

```bash
# for server
linux:~ $ ./json_server

# for client
linux:~ $ ./json_client 127.0.0.1:1234
```

---

## [protobuf](https://github.com/protocolbuffers/protobuf)

```bash
# protocol buffer compiler
# for debian, ubuntu
linux:~ # apt install -y protobuf-compiler

# for binary
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
PB_VER=25.2
# PB_VER=3.20.3
linux:~ # curl -LO $PB_REL/download/v${PB_VER}/protoc-${PB_VER}-linux-x86_64.zip
linux:~ # unzip protoc-25.2-linux-x86_64.zip -d /usr/local/protoc-${PB_VER}
linux:~ # ln -s /usr/local/protoc-${PB_VER}/bin/protoc /usr/local/bin/protoc

linux:~ # protoc --version

# go plugin for protocol compiler
linux:~ $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
linux:~ $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

linux:~ $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
linux:~ $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

```go
// student.pd
syntax = "proto3";

option go_package = "example.com/student";

message Student {
  string name = 1;
  int32  age = 2;
  string gender = 3;
  int32  number = 4;
}
```

```bash
linux:~ $ ls student.pd
linux:~ $ protoc --go_out=. *.proto
linux:~ $ ls example.com/student/student.pb.go
linux:~ $ cd example.com/student
linux:~/example.com/student$ go mod init example.com/student
linux:~/example.com/student$ go mod tidy
```

```go
package main

import (
	"log"

	pb "example.com/student"
	"google.golang.org/protobuf/proto"
)

func main() {
	s := &pb.Student{
		Name:   "Peng Jie",
		Age:    24,
		Gender: "Male",
		Number: 99,
	}

	log.Println(
		s.GetName(),
		s.GetAge(),
		s.GetGender(),
		s.GetNumber(),
	)

	data, err := proto.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)

	ss := &pb.Student{}
	err = proto.Unmarshal(data, ss)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ss)
}
```

```bash
# redirect to local package
linux:~/demo $ go mod edit -replace example.com/student=../example.com/student
linux:~/demo $ go mod tidy
```

---

## [grpc](https://github.com/grpc/grpc)

```go
// helloworld.pb
syntax = "proto3";

option go_package = "example.com/helloworld";
package helloworld;

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

```bash
linux:~ $ ls helloworld.pb
linux:~ $ protoc \
 --go_out=. [--go_opt=<go-opt>] \
 --go-grpc_out=. [--go-grpc_opt=<go-grpc_opt>] \
 *.proto
linux:~ $ ls example.com/helloworld/helloworld.pb.go
linux:~ $ ls example.com/helloworld/helloworld_grpc.pb.go
linux:~ $ cd example.com/helloworld
linux:~/example.com/helloworld $ go mod init example.com/helloworld
linux:~/example.com/helloworld $ go mod tidy
```

```go
// grpc_server.go
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "example.com/helloworld"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

```go
// grpc_client.go
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "example.com/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
```
