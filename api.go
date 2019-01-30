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

func openGPIO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	pinPosition := vars["pos"]

	fmt.Println("opening gpio number : %d", pinPosition)
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
	pin.High()
	err = json.NewEncoder(w).Encode("GPIOOPEN")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send okey")
}

func closeGPIO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	pinPosition := vars["pos"]

	fmt.Println("closing gpio number : %d", pinPosition)
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
	pin.Low()
	err = json.NewEncoder(w).Encode("GPIO CLOSE")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send okey")
}


func toogleGPIO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	pinPosition := vars["pos"]

	fmt.Println("toogle gpio number : %d", pinPosition)
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
	pin.Toggle()
	err = json.NewEncoder(w).Encode("GPIO CLOSE")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send okey")
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
	pin.High()
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
	pin.Low()
	err = json.NewEncoder(w).Encode("GPIOclose")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send okey")
}


func toogleGPIO21(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin := rpio.Pin(21)
	pin.Output()
	pin.Toggle()
	err = json.NewEncoder(w).Encode("GPIOClose")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send okey")
}


func GPIO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	pinPosition := vars["pos"]
	methodeCalling := vars["methode"]

	log.Println("%S gpio number : %d",methodeCalling, pinPosition)
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
		err = json.NewEncoder(w).Encode("GPIO CLOSE")
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

	router.HandleFunc("/api/open21", openGPIO21).Methods("GET")

	router.HandleFunc("/api/open/{pos}", openGPIO).Methods("GET")
	router.HandleFunc("/api/close/{pos}", closeGPIO).Methods("GET")
	router.HandleFunc("/api/toogle/{pos}", toogleGPIO).Methods("GET")
	router.HandleFunc("/api/{methode}/{pos}", toogleGPIO).Methods("GET")


	router.HandleFunc("/api/close21", closeGPIO21).Methods("GET")
	router.HandleFunc("/api/toogle21", toogleGPIO21).Methods("GET")

	// take care of fatal error
	log.Fatal(http.ListenAndServe(":8001", router))
}
