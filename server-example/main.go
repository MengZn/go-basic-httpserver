package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RegRequest struct {
	Name string
}

type RegResponse struct {
	Info string
}

func regHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		var t RegRequest
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		greet := "hello " + t.Name
		info := RegResponse{Info: greet}
		if resp, err := json.Marshal(info); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			w.Write(resp)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/reg", regHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
