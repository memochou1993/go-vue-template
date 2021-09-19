package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed ui/dist
var ui embed.FS

func main() {
	stripped, err := fs.Sub(ui, "ui/dist")
	if err != nil {
		log.Panic(err)
	}
	http.Handle("/", http.FileServer(http.FS(stripped)))

	log.Panic(http.ListenAndServe(":8000", nil))
}
