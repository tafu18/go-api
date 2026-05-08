package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Selam! Go API burası. Proje başarıyla yayında.")
	})

	fmt.Println("Go sunucusu 8080 portunda çalışıyor...")
	http.ListenAndServe(":8080", nil)
}
