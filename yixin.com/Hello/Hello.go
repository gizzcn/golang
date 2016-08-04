package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"yixin.com/Hello/mysqldb"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

type page struct {
	title string
	body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]

	//p2 := loadPage(file)

	fmt.Fprint(w, getUser(id))
}

func loadPage(file string) *page {
	filename := "./web/" + file + ".html"
	body, _ := ioutil.ReadFile(filename)

	return &page{title: file, body: body}
}

func getUser(id string) string {
	return mysqldb.SqlGet("select * from table1 where id = " + id)
}
