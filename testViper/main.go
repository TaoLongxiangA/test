package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/resumes", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "个人信息：\n")
		_, _ = fmt.Fprintf(w, "姓名：%s，\n性别：%s，\n年龄 %d!", Resume.Name, Resume.Gender, Resume.Age)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
