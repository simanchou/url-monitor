package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":99", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	var cs []string
	fmt.Println("Header parameter")
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			fmt.Printf("%s\t=>\t%s\n", k, v[0])
			cs = append(cs, fmt.Sprintf("%40s\t=>\t%s", k, v[0]))
		}
	}

	fmt.Println()
	fmt.Println("Form parameter")
	r.ParseForm()
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			fmt.Printf("%s\t=>\t%s\n", k, v[0])
			cs = append(cs, fmt.Sprintf("%40s\t=>\t%s", k, v[0]))
		}
	}

	w.WriteHeader(http.StatusOK)
	csfs := strings.Join(cs, "\n")
	if (r.Form.Get("user") == "admin") && (r.Form.Get("pass") == "123456") {
		w.Write([]byte("hello, login successful\n"))
		w.Write([]byte(csfs))
	} else {
		w.Write([]byte("hello, login fail\n"))
		w.Write([]byte(csfs))
	}

}
