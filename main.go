package main

import (
	"encoding/hex"
	"fmt"
	"net/http"
)

var store *URLStore = NewURLStore()

func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8080", nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := LongURL(r.FormValue("url"))
	if url == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, string(AddForm))
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := ShortURL(r.URL.Path[1:])
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, string(url), http.StatusFound)
}