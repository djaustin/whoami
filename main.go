package main

import (
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		remoteIP := strings.Split(r.RemoteAddr, ":")[0]
		rw.Write([]byte(remoteIP))
	})
	http.ListenAndServe(":80", nil)
}
