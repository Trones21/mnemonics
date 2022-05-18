//go:build exclude

package main

import "net/http"

func exampleMain() {
	http.Handle("/foo", &fooHandler{Message: "foo call"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe("localhost:5000", nil)
}

type fooHandler struct {
	Message string
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar handler called"))
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(f.Message))
}
