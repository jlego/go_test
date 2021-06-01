package main

//go:generate statik -src=./public -include=*.jpg,*.txt,*.html,*.css,*.js,*.ico,*.png,*.json
import (
	"log"
	"net/http"

	_ "github.com/jlego/go_test/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	http.ListenAndServe(":6003", nil)
}
