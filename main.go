package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Println("首页访问")
	t, err := template.ParseFiles("./view/index.html")
	if err != nil {
		fmt.Println("读取文件失败", err)
	}
	t.Execute(w, nil)

}

func main() {

	fmt.Println("Go Application & Ant Design Pro")

	http.HandleFunc("/index.html", index)

//	http.HandleFunc("/user/login", index)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Web 服务启动失败", err)
	}

}
