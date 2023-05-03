package pages

import (
	"io/ioutil"
	"net/http"
)

func Raw(w http.ResponseWriter, r *http.Request) {
	// read file dari people.json
	data, err := ioutil.ReadFile("people.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
