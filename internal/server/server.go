package server

import (
	"fmt"
	"net/http"
	"strings"
)

type storage interface {
	Set(mType, mName, mValue string)
}

func Serve(addr string, strg storage) error {
	server := http.Server{
		Addr: addr,
	}

	http.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.Split(r.RequestURI[8:], "/")
		strg.Set(urlParts[0], urlParts[1], urlParts[2])

		fmt.Println(strg)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	return server.ListenAndServe()
}
