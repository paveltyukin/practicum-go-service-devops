package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/paveltyukin/practicum-go-service-devops/internal"
	"github.com/paveltyukin/practicum-go-service-devops/pkg"
)

var m = MemStorage{metrics: internal.Metrics{}}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.RequestURI[8:], "/")

	if !pkg.StringInSlice(urlParts[0], []string{"gauge", "counter"}) {
		panic("Metrics error types")
	}

	m.Set(urlParts[1], urlParts[2])

	fmt.Println(m)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/update/", HelloWorld)
	server := &http.Server{
		Addr: "127.0.0.1:8080",
	}

	log.Fatal(server.ListenAndServe())
}
