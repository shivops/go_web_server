package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello,World!!!")
    })


	fmt.Printf("Hello all! This is my first go web server!!!\n")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
