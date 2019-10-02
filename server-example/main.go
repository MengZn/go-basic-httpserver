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

type CheckInRequest struct {
	Book string
}
type CheckInResponse RegResponse

type CheckOutRequest CheckInRequest
type CheckOutResponse RegResponse

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

func checkOutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		var t CheckInRequest
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		detail := "You checkout a " + t.Book + " book"
		info := CheckInResponse{Info: detail}
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

func checkInHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodDelete:
		var t CheckOutRequest
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		detail := "You checkin a " + t.Book + " book"
		info := CheckOutResponse{Info: detail}
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
	http.HandleFunc("/checkout", checkOutHandler)
	http.HandleFunc("/checkin", checkInHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
