package ctrl

import (
	"misc/logger"
	"net/http"
	"os/exec"
	"path"
)

type ExecController struct{}

// TODO: enum for valid cmds
// var valid_cmds = []string{"newpost"}

// Executes given command with parameters passed.
// TODO: Use mux routing frameworkd
func (cf *ExecController) ExecHandler(w http.ResponseWriter, r *http.Request) {
	cmd := path.Base(r.URL.Path)

	switch cmd {
	case "newpost":
		// create a new post
		// TODO: Add Syncronization (using channels) while creating post
		handleNewPost(w, r)
	default:
		// exit if the cmd is not valid
		logger.Logger("Error: " + cmd + " is an invalid cmd or parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handleNewPost(w http.ResponseWriter, r *http.Request) {
	names, ok := r.URL.Query()["name"]

	if !ok || len(names) < 1 || len(names[0]) < 1 {
		logger.Logger("Error: Invalid name")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := exec.Command("api_hops_newpost", names[0]+".md").Output()

	if err != nil {
		logger.Logger("Error: while creating post")
		logger.Logger(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		logger.Logger(out)
		logger.Logger("Successfully create post ")
	}
}

func contains(array []string, item string) bool {
	for _, a := range array {
		if a == item {
			return true
		}
	}
	return false
}
