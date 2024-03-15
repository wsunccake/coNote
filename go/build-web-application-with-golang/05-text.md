# text

---

## content

- [xml](#xml)
  - [read xml](#read-xml)
  - [write xml](#write-xml)
- [json](#json)
  - [read json](#read-json)
  - [write json](#write-json)
- [regexp](#regexp)
- [template](#template)
  - [basic](#basic)
  - [format](#format)
  - [variable](#variable)
  - [condition](#condition)
  - [compare](#compare)
  - [loop](#loop)
  - [nested template](#nested-template)
  - [parse file](#parse-file)
- [file](#file)
  - [dir](#dir)
  - [write file](#write-file)
  - [read file](#read-file)
- [string](#string)
  - [string operate](#string-operate)
  - [string convert](#string-convert)

---

## xml

### read xml

```go
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}
```

```xml
<!-- servers.xml -->
<?xml version="1.0" encoding="utf-8"?>
<servers version="1">
    <server>
        <serverName>Shanghai_VPN</serverName>
        <serverIP>127.0.0.1</serverIP>
    </server>
    <server>
        <serverName>Beijing_VPN</serverName>
        <serverIP>127.0.0.2</serverIP>
    </server>
</servers>
```

### write xml

```go
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))

	os.Stdout.Write(output)
}
```

---

## json

### read json

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice

	file, err := os.Open("servers.json")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// data := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	json.Unmarshal([]byte(data), &s)
	fmt.Println(s)
}
```

```json
// servers.json
{
  "servers": [
    {
      "serverName": "Shanghai_VPN",
      "serverIP": "127.0.0.1"
    },
    {
      "serverName": "Beijing_VPN",
      "serverIP": "127.0.0.2"
    }
  ]
}
```

### write json

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP,omitempty"`
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	s.Servers = append(s.Servers, Server{ServerName: "HongKong_VPN", ServerIP: ""})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
```

---

## regexp

```go
package main

import (
	"fmt"
	"os"
	"regexp"
)

func IsIPv4(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func main() {
	ip := os.Args[1]

	if IsIPv4(ip) {
		fmt.Printf("%s is IPv4\n", ip)
	} else {
		fmt.Printf("%s is not IPv4\n", ip)
	}
}
```

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("https://go.dev/")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	// replace HTML tag uppercase to lowercase
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// remove style
	re, _ = regexp.Compile(`\<style[\S\s]+?\</style\>`)
	src = re.ReplaceAllString(src, "")

	// remove script
	re = regexp.MustCompile(`\<script[\S\s]+?\</script\>`)
	src = re.ReplaceAllString(src, "")

	// remove HTML tag
	re = regexp.MustCompile(`\<[\S\s]+?\>`)
	src = re.ReplaceAllString(src, "\n")

	// remove space
	re = regexp.MustCompile(`\s{2,}`)
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}
```

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"
	fmt.Println(a, len(a))

	re, _ := regexp.Compile("[a-z]{2,4}")

	// find to match first
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one), one)

	// find to match all
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)
	for _, v := range all {
		fmt.Println(string(v), v)
	}

	// find to match first then return begin and end
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", a[index[0]:index[1]], index)

	// find to match all then return begin and end
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)
	for _, v := range allindex {
		fmt.Println(a[v[0]:v[1]], v)
	}

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	// find to sub match
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v), v)
	}

	// like as FindIndex
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println("FindSubmatchIndex", submatchindex)
	for _, v := range submatchindex {
		fmt.Println(a[v:], v)
	}

	// like as FindAll FindAllSubmatch
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println("FindAllSubmatch", submatchall)
	for _, v := range submatchall {
		for _, vv := range v {
			fmt.Println(vv, string(vv))
		}
	}

	// like as FindAllIndex
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println("FindAllSubmatchIndex", submatchallindex)
	for _, v := range submatchallindex {
		fmt.Println(a[v[0]:v[1]], v)
	}
}
```

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	src := []byte(`
        call hello alice
        hello bob
        call hello eve
    `)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
```

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	lib := "/usr/lib/python2.6/site-packages/gtk-2.0/gconf.so"
	fmt.Println(lib)

	var r1 *regexp.Regexp = regexp.MustCompile(`(.*)\.(.*?)$`)
	var groups1 []string = r1.FindStringSubmatch(lib)
	fmt.Printf("%s <=> %s\n", groups1[1], groups1[2])
	fmt.Println(r1.ReplaceAllString(lib, "$1"))

	r2 := regexp.MustCompile(`(.*?)\.(.*)`)
	groups2 := r2.FindStringSubmatch(lib)
	fmt.Printf("%s <=> %s\n", groups2[1], groups2[2])
	fmt.Println(r2.ReplaceAllString(lib, "$1"))

	r3 := regexp.MustCompile(`(.*?)/(.*)`)
	groups3 := r3.FindStringSubmatch(lib)
	fmt.Printf("%s <=> %s\n", groups3[1], groups3[2])
	fmt.Println(r3.ReplaceAllString(lib, "$1"))

	r4 := regexp.MustCompile(`(.*)/(.*)`)
	groups4 := r4.FindStringSubmatch(lib)
	fmt.Printf("%s <=> %s\n", groups4[1], groups4[2])
	fmt.Println(r4.ReplaceAllString(lib, "$1"))
}
```

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	sentence := `wlan0 down AP wlan0 0 00:00:00:00:00:00 Wireless1
wlan1 down AP wlan1 0 00:00:00:00:00:00 Wireless2
wlan2 up AP wlan2 0 00:00:00:00:00:00 Wireless3
wlan36 up AP wlan36 1 00:00:00:00:00:00 Wireless13
wlan37 down AP wlan37 1 00:00:00:00:00:00 Wireless14
wlan38 up AP wlan38 1 00:00:00:00:00:00 Wireless15
wlan39 down AP wlan39 1 00:00:00:00:00:00 Wireless16`

	lazy_pattern := "wlan\\d+.*up[\\s\\S]+?Wireless\\d+"
	greedy_pattern1 := `wlan\d+.*up[\s\S]+Wireless\d+`
	greedy_pattern2 := `wlan\d+.*up[\s\S]*Wireless\d+`
	match, _ := regexp.MatchString(lazy_pattern, sentence)
	fmt.Println(match)

	r1, _ := regexp.Compile(lazy_pattern)
	fmt.Println(r1.MatchString(sentence))
	fmt.Println(r1.FindString(sentence))
	fmt.Println(r1.FindAllString(sentence, -1))

	r2, _ := regexp.Compile(greedy_pattern1)
	fmt.Println(r2.MatchString(sentence))
	fmt.Println(r2.FindString(sentence))
	fmt.Println(r2.FindAllString(sentence, -1))

	r3, _ := regexp.Compile(greedy_pattern2)
	fmt.Println(r3.MatchString(sentence))
	fmt.Println(r3.FindString(sentence))
	fmt.Println(r3.FindAllString(sentence, -1))
}
```

---

## template

### basic

```go
package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

type Man struct {
	Say  string
	Name string
}

func (my *Man) SayHello() string { //没参数
	return "world"
}

func (my *Man) SayYouName(name string) string { //有参数
	return "my name is : " + name
}

func main() {
	n := "world"
	t1, err := template.New("test").Parse("hello, {{.}}\n")
	if err != nil {
		panic(err)
	}
	err = t1.Execute(os.Stdout, n) 	// pass string
	if err != nil {
		panic(err)
	}

	t2 := template.New("fieldname example")
	t2, _ = t2.Parse("hello {{.UserName}}!")
	p := Person{UserName: "Astaxie"}
	t2.Execute(os.Stdout, p)			// pass struct

	m := &Man{Say: "hello", Name: "jo"}
	c := `{{$var1 := .Say}} {{$var2 := .SayHello}} {{$var3 := .SayYouName .Name}}
{{$var1}}
{{$var2}}
{{$var3}}
`
	t3 := template.New("test")			// assign variable
	t3, _ = t3.Parse(c)
	t3.Execute(os.Stdout, m)
}
```

### format

```go
package main

import (
	"os"
	"text/template"
)

func main() {
	// comment
	// trim
	const c = `{{/* comment */}}
{{- 123 -}}>{{- 45 }}<{{ 90 }}
{{- print 12 }}
{{ printf "%03d" 12 }}
{{ println 12 }}
{{ 12 | printf "%03d" }}
{{ 3 | printf "%d+%d=%d" 1 2 }}
`
	t := template.Must(template.New("tmpl").Parse(c))
	t.Execute(os.Stdout, nil)
}
```

### variable

```go
package main

import (
	"os"
	"text/template"
)

func main() {
	const c = `
{{- $userName := "MegaShow" -}} {{ printf "userName: %s" $userName }}
{{ $realName := $userName -}} {{ printf "real: %s" $realName }}
{{ $realName | printf "Hello, %s." }}
`

	t := template.Must(template.New("tmpl").Parse(c))
	t.Execute(os.Stdout, nil)
}

```

### condition

if-else

```go
package main

import (
	"os"
	"text/template"
)

func main() {
	c := `{{- $pipeline := true -}}
{{ if $pipeline }} A1 {{ end }}
{{ if $pipeline }} B1 {{ else }} B2 {{ end }}
{{ if $pipeline }} C1 {{ else if $pipeline }} C2 {{ end }}
{{ if $pipeline }} D1 {{ else }} D2 {{ if $pipeline }} D3 {{ end }}{{ end }}
{{ if $pipeline }} E1 {{ else if $pipeline }} E2 {{ else }} E3 {{ end }}
`
	t := template.New("test")
	t, _ = t.Parse(c)
	t.Execute(os.Stdout, nil)
}
```

### compare

```go
package main

import (
	"os"
	"text/template"
)

func main() {
	c := `
{{- /* false      => true  */ -}}
{{- /* true       => false */ -}}
{{- /* nil        => true  */ -}}
{{- /* ""         => true  */ -}}
{{- /* "MegaShow" => false */ -}}
{{ "{{ not . }}" }} => {{ not . }}

{{- /* X:true       Y:true      => true      */ -}}
{{- /* X:false      Y:true      => false     */ -}}
{{- /* X:""         Y:true      => ""        */ -}}
{{- /* X:"MegaShow" Y:"icytown" => "icytown" */}}
{{ "{{ and .X .Y }}" }}	=> {{ and .X .Y }}	=> if .X then .Y else .X

{{- /* X:false      Y:false     => false      */ -}}
{{- /* X:false      Y:true      => true       */ -}}
{{- /* X:false      Y:""        => ""         */ -}}
{{- /* X:"MegaShow" Y:"icytown" => "MegaShow" */}}
{{ "{{ or .X .Y }}" }}	=> {{ or .X .Y }}	=> if .X then .X else .Y

{{- /* X, Y both not empty or Z not empty  */}}
{{ "{{ or (and (not .X) (not .Y)) (not .Z) }}" }} => {{ or (and (not .X) (not .Y)) (not .Z) }}
{{- /* fetch X, Y, Z first not empty, if all empty to return "MegaShow" */}}
{{ "{{ or (or (or .X .Y) .Z) \"MegaShow\" }}" }} => {{ or (or (or .X .Y) .Z) "MegaShow" }}

{{ "{{ eq .X .Y }}" }} =>	{{ eq .X .Y }} => {{ "{{/* .X == .Y */}}" }}
{{ "{{ ne .X .Y }}" }} =>	{{ ne .X .Y }} => {{ "{{/* .X != .Y */}}" }}
{{ "{{ lt .X .Y }}" }} =>	{{ lt .X .Y }} => {{ "{{/* .X < .Y  */}}" }}
{{ "{{ le .X .Y }}" }} =>	{{ le .X .Y }} => {{ "{{/* .X <= .Y */}}" }}
{{ "{{ gt .X .Y }}" }} =>	{{ gt .X .Y }} => {{ "{{/* .X > .Y  */}}" }}
{{ "{{ ge .X .Y }}" }} =>	{{ ge .X .Y }} => {{ "{{/* .X >= .Y */}}" }}
`

	t := template.Must(template.New("tmpl").Parse(c))
	t.Execute(os.Stdout, nil)
}
```

### loop

with-end and range-end

```go
package main

import (
	"os"
	"text/template"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	c: = `hello {{ .UserName }}!
{{ range .Emails }}
an email {{ . }}
{{ else }}
no email
{{ end }}
{{ with .Friends }}
	{{ range . }}
my friend name is {{ .Fname }}
    {{ end }}
{{ end }}
`
	t := template.New("fieldname example")
	t, _ = t.Parse(c)
	t.Execute(os.Stdout, p)
}
```

```go
package main

import (
	"os"
	"text/template"
)

func mkSlice(args ...interface{}) []interface{} {
	return args
}

func main() {
	tmpl := `
{{- $slice := mkSlice "a" 5 "b" -}}
{{- range $slice -}}
     {{ . }}
{{ end -}}
`
	funcMap := map[string]interface{}{"mkSlice": mkSlice}
	t := template.New("demo").Funcs(template.FuncMap(funcMap))
	template.Must(t.Parse(tmpl))
	t.ExecuteTemplate(os.Stdout, "demo", nil)
}
```

```go
package main

import (
	"os"
	"text/template"
)

type User struct {
	Id   int
	Name string
}

func main() {
	const tmpl = `
{{- range . -}}
	{{ .Id }} - {{ .Name }}
{{ end -}}
`
	var myuserlist []User = []User{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}

	t := template.Must(template.New("tmpl").Parse(tmpl))

	t.Execute(os.Stdout, myuserlist)
}
```

### nested template

```go
package main

import (
	"os"
	"text/template"
)

func main() {

	const c = `{{ "{{- define \"header\" -}} header is {{ . }} ... {{- end -}}" }} =>
{{- define "header" -}} header is {{ . }} ... {{- end }}
{{ "{{ template \"header\" }}" }} => {{ template "header" }}
{{ "{{ template \"header\" . }}" }} => {{ template "header" . }}
{{ "{{ template \"header\" \"hello\" }}" }} => {{ template "header" "hello" }}

{{ "{{ block \"footer\" \"hello\" }} footer is {{ . }} ... {{ end }}" }} =>
{{- block "footer" "hello" }} footer is {{ . }} ... {{ end }}
{{ "{{ template \"footer\" }}" }}=> {{ template "footer" }}
`
	t := template.Must(template.New("tmpl").Parse(c))
	t.Execute(os.Stdout, nil)
	t.ExecuteTemplate(os.Stdout, "header", "Hello")
	t.ExecuteTemplate(os.Stdout, "footer", "Hello")
}
```

### parse file

```html
<!-- index.html -->
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta title="index" />
  </head>
  <body>
    <p>Hello {{ .UserName }}</p>
  </body>
</html>
```

```go
// main.go
package main

import (
	"html/template"
	"net/http"
)

var htmlTemplate = template.Must(template.ParseFiles("index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	htmlTemplate.Execute(w, map[string]any{"UserName": "MegaShow"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe(":8080", mux)
}
```

---

## file

### dir

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("astaxie", 0777)
	os.MkdirAll("astaxie/test1/test2", 0777)
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("astaxie")
}
```

### write file

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "astaxie.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}
}
```

### read file

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "asatxie.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
```

---

## string

### string operate

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) //foo, bar, baz

	fmt.Println(strings.Index("chicken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "dmr")) // -1

	fmt.Println("ba" + strings.Repeat("na", 2)) // banana

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]

	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))       // ["Achtung"]
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) // Fields are: ["foo" "bar" "baz"]
}
```

### string convert

```go
package main

import (
	"fmt"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
func main() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '單')
	fmt.Println(string(str))

	a1 := strconv.FormatBool(false)
	b1 := strconv.FormatFloat(123.23, 'g', 12, 64)
	c1 := strconv.FormatInt(1234, 10)
	d1 := strconv.FormatUint(12345, 10)
	e1 := strconv.Itoa(1023)
	fmt.Println(a1, b1, c1, d1, e1)

	a2, err := strconv.ParseBool("false")
	checkError(err)
	b2, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c2, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d2, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e2, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a2, b2, c2, d2, e2)
}
```
