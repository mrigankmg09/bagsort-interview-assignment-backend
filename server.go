package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	publicPath := filepath.Join(parent, "bagsort-interview-assignment-frontend/public")
	fs := http.FileServer(http.Dir(filepath.Join(publicPath, "resources")))
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))
	fs = http.FileServer(http.Dir(publicPath))
	http.Handle("/", fs)
	http.HandleFunc("/api/date-diff", calculateDayDifference)
	port := ":3030"
	log.Fatal(http.ListenAndServe(port, nil))
}

func calculateDayDifference(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
			return
		}
		currentTime := time.Now()
		selectedDate, err := time.ParseInLocation("2006-01-02", string(reqBody), currentTime.Location())
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Invalid Date", 400)
			return
		}
		json.NewEncoder(w).Encode(math.Ceil(selectedDate.Sub(currentTime).Hours() / 24))
	} else {
		http.NotFound(w, r)
	}
}
