package server

import (
	"fmt"
	"net/http"
	"strings"
)

type storage interface {
	Set(mType, mName, mValue string)
}

func createUpdateHandler(strg storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.Split(r.RequestURI[8:], "/")
		strg.Set(urlParts[0], urlParts[1], urlParts[2])

		fmt.Println(strg)
		w.WriteHeader(http.StatusOK)
	}
}

func createNotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func Serve(addr string, strg storage) error {
	server := http.Server{
		Addr: addr,
	}

	http.HandleFunc("/update/", createUpdateHandler(strg))

	http.HandleFunc("/", createNotFoundHandler)

	return server.ListenAndServe()
}
