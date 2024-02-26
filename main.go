package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	var buf bytes.Buffer
	encoder := gzip.NewWriter(&buf)
	written := 0
	content := "<b>hi</b><br>"
	for written < 1024*1024*1024 {
		encoder.Write([]byte(content))
		written += len(content)
	}
	encoder.Close()
	log.Printf("ready")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if enc := r.Header.Get("Accept-Encoding"); !strings.Contains(enc, "gzip") {
			http.Error(w, "no", http.StatusNotAcceptable)
			return
		}

		w.Header().Add("Content-Encoding", "gzip")
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		io.Copy(w, bytes.NewReader(buf.Bytes()))
	})
	log.Fatal(http.ListenAndServe(":1024", handler))
}
