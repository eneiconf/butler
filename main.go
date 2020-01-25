package main

import (
	"net/http"
)


func main() {
	http.Handle("/ambassadors/img/", http.StripPrefix("/ambassadors/img/", http.FileServer(http.Dir("./assets/images/ambassadors"))))
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./assets/jsons/"))))
	http.ListenAndServe(":8000", nil)
}
