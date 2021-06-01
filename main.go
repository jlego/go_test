package main

//go:generate statik -src=./public -include=*.jpg,*.txt,*.html,*.css,*.js,*.ico,*.png,*.json
import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	_ "github.com/jlego/go_test/statik"
	"github.com/rakyll/statik/fs"
)

type handle struct {
	host string
	port string
	path string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + this.host + ":" + this.port)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// interfaces, err := net.Interfaces()
	// if err != nil {
	// 	panic("Error:" + err.Error())
	// }
	// for _, inter := range interfaces {
	// 	fmt.Println(inter.HardwareAddr)
	// }
	h := &handle{host: "0.0.0.0", port: "1337"}
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	http.Handle("/auth/local", h)
	http.ListenAndServe(":6003", nil)
}
