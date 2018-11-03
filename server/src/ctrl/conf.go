package ctrl

import (
	"encoding/json"
	"misc/logger"
	"model"
	"net/http"
	"path"
	"service"
)

type ConfController struct{}

var (
	ghService  *service.GitHubService = new(service.GitHubService)
	cnfService *service.ConfigService = new(service.ConfigService)

	cnfData model.ConfigData
)

func (cf *ConfController) ConfigHandler(w http.ResponseWriter, r *http.Request) {
	resource := path.Base(r.URL.Path)

	logger.Logger("method: " + r.Method)
	switch resource {
	// GET: /getRepos?userId=<github-userId>
	case "getRepos":
		// w.Header().Set("Content-Type", "text/html")
		logger.Logger("Getting git repos")
		ghService.GetRepos(w, r)

	case "updateConfig":
		// POST: /updateConfig
		// Update config
		logger.Logger("Updating configuration")

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&cnfData)
		if err != nil {
			logger.Logger(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		cnfService.UpdateConfig(cnfData, w, r)

		logger.Logger("Updated configuration!")

	default:
		logger.Logger("Error: invalid resource requested")
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
