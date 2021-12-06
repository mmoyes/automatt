package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("---\nAutomatt\n---")
	router := mux.NewRouter()
	router.HandleFunc("/api/{category}", TestHandler)
	router.HandleFunc("/api/{category}/{cmd}", TestHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8800",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Listening on %v", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Printf("Vars: %v\n", vars)

	if vars["category"] == "health" {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		fmt.Printf("\tResponse: %v\n", "OK")
		w.WriteHeader(http.StatusOK)
		return
	}

	if vars["category"] == "kube" {

		if vars["cmd"] == "BE-reset" {
			c, err := RunCmd("/bin/bash", "./BE-reset.sh")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error: %v\n", err.Error())
				fmt.Fprintf(w, "%v\n", c)
			}
			w.WriteHeader(http.StatusOK)
			fmt.Printf("\tResponse: %v\n", c)
			return
		}

		fmt.Printf("\tResponse: %v\n", "OK")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK) // TODO error handle
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
	fmt.Printf("Vars: %v\n", vars)
}

func RunCmd(cmd string, args ...string) (string, error) {
	e := exec.Command(cmd, args...)
	var outb, errb bytes.Buffer
	e.Stdout = &outb
	e.Stderr = &errb
	err := e.Run()
	if err != nil {
		return errb.String(), err
	}
	return outb.String(), nil
}
