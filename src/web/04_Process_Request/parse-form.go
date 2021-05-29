package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// fmt.Fprintln(w, r.Form) // 包含GET和POST参数
	// fmt.Fprintln(w, r.PostForm) // 仅包含POST参数 application/x-www-form-urlencoded
	fmt.Fprintln(w, r.MultipartForm) // 仅包含POST参数 multipart/form-data
}
func processMultiPart(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm) // 仅包含POST参数 multipart/form-data
}

func formValue(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, "(1)", r.FormValue("hello"))
	fmt.Fprintln(w, "(2)", r.PostFormValue("hello"))
	fmt.Fprintln(w, "(3)", r.PostForm)
	fmt.Fprintln(w, "(4)", r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/processMultiPart", processMultiPart)
	http.HandleFunc("/formvalue", formValue)
	server.ListenAndServe()
}
