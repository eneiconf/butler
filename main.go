package main

import (
	a "butler/internal/ambassador"
	h "butler/internal/homepage"
	s "butler/internal/speakers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

func getSpeakers(w http.ResponseWriter, r *http.Request) {
	var speakers s.Speakers

	jsonFile, err := os.Open("./assets/jsons/speakers.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, e := ioutil.ReadAll(jsonFile)
	if e != nil {
		fmt.Println(e)
	}
	json.Unmarshal(byteValue, &speakers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(speakers)

}

func getAmbassadors(w http.ResponseWriter, r *http.Request) {
	var ambassadors a.Ambassadors

	jsonFile, err := os.Open("./assets/jsons/ambassadors.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, e := ioutil.ReadAll(jsonFile)
	if e != nil {
		fmt.Println(e)
	}
	json.Unmarshal(byteValue, &ambassadors)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ambassadors)

}

func getHomePage(w http.ResponseWriter, r *http.Request) {
	var homepage h.HomePage

	jsonFile, err := os.Open("./assets/jsons/homepage.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, e := ioutil.ReadAll(jsonFile)
	if e != nil {
		fmt.Println(e)
	}
	json.Unmarshal(byteValue, &homepage)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(homepage)

}

func getAmbassadorImg(w http.ResponseWriter, r *http.Request) {
	name := (mux.Vars(r))["ambassador"]
	localURL := "./assets/images/ambassadors/" + name
	imgfile, err := os.Open(localURL)
	if err != nil {
		fmt.Println(err)
	}
	defer imgfile.Close()
	byteValue, e := ioutil.ReadAll(imgfile)
	if e != nil {
		fmt.Println(e)
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(byteValue)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "How may I serve you?")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/ambassadors/", getAmbassadors).Methods("GET")
	router.HandleFunc("/ambassadors/img/{ambassador}", getAmbassadorImg).Methods("GET")
	router.HandleFunc("/homepage/", getHomePage).Methods("GET")
	router.HandleFunc("/speakers/", getSpeakers).Methods("GET")
	http.ListenAndServe(":8000", router)
}
