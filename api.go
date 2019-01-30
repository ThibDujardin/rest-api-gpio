package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stianeikeland/go-rpio"
	"log"
	"net/http"
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

func openGPIO21(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin := rpio.Pin(21)
	pin.Output()

	err = json.NewEncoder(w).Encode("GPIOOPEN")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send okey")
}

func closeGPIO21(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin := rpio.Pin(21)
	pin.Output()

	pin.PullUp()
	err = json.NewEncoder(w).Encode("GPIOOPEN")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send okey")
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

	router.HandleFunc("/api/open21", openGPIO21).Methods("GET")

	router.HandleFunc("/api/close21", closeGPIO21).Methods("GET")

	// take care of fatal error
	log.Fatal(http.ListenAndServe(":8001", router))
}
