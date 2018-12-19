package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	filename := flag.String("f", ".", "the file to serve")
	addr := flag.String("http", ":9090", "the addr to serve on")
	flag.Parse()

	log.SetFlags(0)

	if _, err := os.Stat(*filename); os.IsNotExist(err) {
		log.Fatalf("file %q does not exist", *filename)
	} else if err != nil {
		log.Fatalf("filesystem error checking %q: %v", *filename, err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/health" {
			return
		}

		http.ServeFile(w, r, *filename+r.URL.Path)
	})

	log.Println("listening on", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
