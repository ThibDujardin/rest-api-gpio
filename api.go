package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stianeikeland/go-rpio"
	"log"
	"net/http"
	"strconv"
	"time"
)

type test_struct struct {
	test string
}

func getHour(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	t := time.Now()

	err := json.NewEncoder(w).Encode(t.Format(time.RFC3339))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send okey")
}

// Normal post
func tryPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", len(r.PostForm))
	test := r.FormValue("test")
	fmt.Fprintf(w, "Name = %s\n", test)
	fmt.Println(test)
}

// POST with JSON
func test(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.test)
}

func GPIO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	pinPosition := vars["pos"]
	methodeCalling := vars["methode"]

	log.Printf("%s gpio number : %d",methodeCalling, pinPosition)
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	pos, err := strconv.Atoi(pinPosition)
	if err != nil {
		panic(err)
		return
	}
	defer rpio.Close()
	pin := rpio.Pin(pos)
	pin.Output()

	if (methodeCalling == "toogle"){
		pin.Toggle()
		err = json.NewEncoder(w).Encode("GPIO TOOGLE")
	}else if (methodeCalling == "open"){
		pin.High()
		err = json.NewEncoder(w).Encode("GPIO OPEN")
	}else if (methodeCalling == "close"){
		pin.Low()
		err = json.NewEncoder(w).Encode("GPIO CLOSE")
	}else {
		err = json.NewEncoder(w).Encode("ERROR BAD NAME METHODE - MUST BE open OR close OR toogle")
	}
 	if err != nil {
		panic(err)
		return
	}
	log.Println("send okey")
}

// our main function
func main() {
	// define new router
	router := mux.NewRouter()

	// handler
	router.HandleFunc("/api/hour", getHour).Methods("GET")

	//no JSON
	router.HandleFunc("/api/tryPost", tryPost).Methods("POST")

	//WITH JSON
	router.HandleFunc("/api/tryPostJSON", test).Methods("POST")

	router.HandleFunc("/api/{methode}/{pos}", GPIO).Methods("GET")

	// take care of fatal error
	log.Fatal(http.ListenAndServe(":8001", router))
}
