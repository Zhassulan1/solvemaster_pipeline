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
	pullCMD := exec.Command("git", "-C", "/home/ubuntu/solvemaster_pipeline/", "pull")
	supervisorCMD := exec.Command("supervisorctl", "restart", "uvi:uvicorn")

	pullOutput, pullERR := pullCMD.CombinedOutput()
	if pullERR != nil {
		// if there was any error, print it here
		err := writeJSON(w, 500, envelope{
			"git_pull_err":    pullERR,
			"git_pull_output": pullOutput,
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
			"git_pull_output":   pullOutput,
			"supervisor_error":  supervisorERR,
			"supervisor_output": supervisorOutput,
		}, nil)

		if err != nil {
			w.WriteHeader(500)
		}
		return
	}

	writeJSON(w, 200, envelope{
		"git_pull_output":   pullOutput,
		"supervisor_output": supervisorOutput,
	}, nil)
	return
}
