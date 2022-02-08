package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	s := strings.Split(path, "/")
	p := make([]string, len(s))

	l := 0
	for _, e := range s {
		switch {
		case e == ".":
			continue
		case e == "":
			continue
		case e == "..":
			if l == 0 {
				continue
			}
			l--
		default:
			p[l] = e
			l++
		}
	}

	r := ""
	if l > 0 {
		r = strings.Join(p[0:l], "/")
	}

	return "/" + r
}

func main() {
	i := "/home/"
	o := "/home"
	if simplifyPath(i) != o {
		fmt.Println(o)
	}

	i = "/../"
	o = "/"
	if simplifyPath(i) != o {
		fmt.Println(o)
	}

	i = "/home//foo/"
	o = "/home/foo"
	if simplifyPath(i) != o {
		fmt.Println(o)
	}

	i = "/a/./b/../../c/"
	o = "/c"
	if simplifyPath(i) != o {
		fmt.Println(i)
	}

	i = "/a/../../b/../c//.//"
	o = "/c"
	if simplifyPath(i) != o {
		fmt.Println(i)
	}

}
