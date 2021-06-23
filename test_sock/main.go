package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	mydata := `{"data":{"region":"cna","app_id":null}}`
	fmt.Println(mydata)
	io.WriteString(w, mydata)
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":82", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
