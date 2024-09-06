package main

import (
	"net/http"
	"os/exec"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is working"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sh", "./update.sh")

	_, err := cmd.CombinedOutput()

	if err != nil {
		// if there was any error, print it here
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("OK"))
	// otherwise, print the output from running the command
	// fmt.Println("Output: ", string(out))
}
