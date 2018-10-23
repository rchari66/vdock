package ctrl

import (
	"misc/logger"
	"net/http"
	"os/exec"
)

type PublishController struct{}

func (cf *PublishController) PublishHandler(w http.ResponseWriter, r *http.Request) {

	// check if configuration is not set; then return with error
	msgs, ok := r.URL.Query()["commitMessage"]

	if !ok || len(msgs) < 1 || len(msgs[0]) < 1 {
		logger.Logger("Error: Invalid name")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := exec.Command("api_hops_publish", msgs[0]).Output()

	if err != nil {
		logger.Logger("Error: unable to publish to github")
		logger.Logger(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		logger.Logger(string(out))
		logger.Logger("Successfully published on to github ")
	}
}
