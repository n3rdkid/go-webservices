package main

import "net/http"

// Using http.Handle
type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

// Using http.HandleFunc
func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar called"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "foo Called"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}
