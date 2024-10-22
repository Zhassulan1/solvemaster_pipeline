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
	pullCMD := exec.Command("git", "-C", "/home/ubuntu/solvmaster_backend/", "pull")
	supervisorCMD := exec.Command("supervisorctl", "restart", "uvi:uvicorn")

	pullOutput, pullERR := pullCMD.CombinedOutput()
	if pullERR != nil {
		// if there was any error, print it here
		err := writeJSON(w, 500, envelope{
			"git_pull_err":    pullERR.Error(),
			"git_pull_output": string(pullOutput),
		}, nil)
		if err != nil {
			w.WriteHeader(500)
		}
		return
	}

	supervisorOutput, supervisorERR := supervisorCMD.CombinedOutput()
	if supervisorERR != nil {
		// if there was any error, print it here
		err := writeJSON(w, 500, envelope{
			"git_pull_output":   string(pullOutput),
			"supervisor_error":  supervisorERR.Error(),
			"supervisor_output": string(supervisorOutput),
		}, nil)

		if err != nil {
			w.WriteHeader(500)
		}
		return
	}

	writeJSON(w, 200, envelope{
		"git_pull_output":   string(pullOutput),
		"supervisor_output": string(supervisorOutput),
	}, nil)
	return
}
