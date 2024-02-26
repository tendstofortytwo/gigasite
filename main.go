package main

import (
	"compress/gzip"
	"log"
	"net/http"
	"strings"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if enc := r.Header.Get("Accept-Encoding"); !strings.Contains(enc, "gzip") {
			http.Error(w, "no", http.StatusNotAcceptable)
			return
		}

		w.Header().Add("Content-Encoding", "gzip")
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		encoder := gzip.NewWriter(w)
		defer encoder.Close()

		written := 0
		content := "<b>hi</b><br>"
		for written < 1024*1024*1024 {
			encoder.Write([]byte(content))
			written += len(content)
		}
	})
	log.Fatal(http.ListenAndServe(":1024", handler))
}
