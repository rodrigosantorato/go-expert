package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Album struct {
	Name    string   `json:"name"`
	Band    string   `json:"band"`
	Members []string `json:"members"`
}

func main() {
	var a *Album

	http.HandleFunc("/", handler)
	go http.ListenAndServe(":8080", nil)

	res, err := http.Get("http://0.0.0.0:8080/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&a)
	if err != nil {
		panic(err)
	}

	m, err := json.Marshal(a)
	fmt.Println(string(m))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	a := &Album{
		Name:    "War of Being",
		Band:    "TesseracT",
		Members: []string{"Alec", "Jay", "James", "Amos", "Daniel"},
	}

	err := json.NewEncoder(w).Encode(a)
	if err != nil {
		panic(err)
	}
}
