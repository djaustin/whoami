package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Write([]byte(r.RemoteAddr))
	})
	http.ListenAndServe(":80", nil)
}
