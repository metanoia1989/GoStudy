package main

import (
	"fmt"
	"io/ioutil"
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

func processFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

// http.Request.FormFile 处理单个文件，返回给定键的第一个值
func singleFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/processMultiPart", processMultiPart)
	http.HandleFunc("/processFile", processFile)
	http.HandleFunc("/singleFile", singleFile)
	http.HandleFunc("/formvalue", formValue)
	server.ListenAndServe()
}
