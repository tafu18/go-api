package main

import (
	"fmt"
	"net/http"
)

func main() {
	// İleride burada internal/handler katmanını çağıracaksınız
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Profesyonel Go Yapısı Aktif! \n\nKatmanlar: cmd, internal (handler, service, repository, model)")
	})

	fmt.Println("Sunucu 8080 portunda, profesyonel dizin yapısıyla çalışıyor...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
