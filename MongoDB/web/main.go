package main

import (
	"flag"
	"log"
	"net/http"
)

func main(){
	var addr = flag.String("addr",":8081","webサイトのアドレス")
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/",http.StripPrefix("/",http.FileServer(http.Dir("public"))))
	log.Println("webサイトのアドレス:",*addr)
	http.ListenAndServe(*addr,mux)
}