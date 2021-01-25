package main

import (
	"encoding/json"
	"net/http"
)

type profile struct {
	Name    string
	Hobbies []string
}

// heroku buildpacks:set heroku/go
// heroku create go-test-by-sai --buildpack heroku/go

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p := profile{"Sai", []string{"playing cricket", "programming"}}

	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
