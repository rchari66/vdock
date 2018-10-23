package service

import (
	"misc/logger"
	"model"
	"net/http"
	"os/exec"
)

type ConfigService struct{}

// return github repos for given users
func (ghService *ConfigService) UpdateConfig(cnfConfig model.ConfigData, w http.ResponseWriter, r *http.Request) {
	// move current blog & site repos to bkp folder.
	// Checkout new repos to /ws/
	logger.Logger("Cloning blogRepo from github")

	out, err := exec.Command("api_hops_cloneBlog", cnfConfig.BlogRepo).Output()

	if err != nil {
		logger.Logger("Error: while clone blog repository")
		logger.Logger(err)
		w.WriteHeader(http.StatusBadRequest)
		cleanup(cnfConfig)
		return
	} else {
		logger.Logger(out)
		logger.Logger("Successfully cloned blog repository")
	}

	// store git credentials globally
	logger.Logger("Storing git creds globally")
	out, err = exec.Command("_api_hops_save_git_creds", cnfConfig.UserId, cnfConfig.Password, cnfConfig.Email).Output()

	if err != nil {
		logger.Logger("Error: unable to store git creds")
		logger.Logger(err)
		w.WriteHeader(http.StatusBadRequest)
		cleanup(cnfConfig)
		return
	} else {
		logger.Logger(out)
		logger.Logger("Successfully stored/cached git creds")
	}

	// TODO: seperate boilerplate code
	logger.Logger("Cloning siteRepo from github")
	out, err = exec.Command("api_hops_cloneSite", cnfConfig.SiteRepo).Output()

	if err != nil {
		logger.Logger("Error: while clone site repository")
		logger.Logger(err)
		w.WriteHeader(http.StatusBadRequest)
		cleanup(cnfConfig)
		return
	} else {
		logger.Logger(out)
		logger.Logger("Successfully cloned site repository")
	}

	// startHugoServer()
	logger.Logger("Starting hugo... ")

	startCmd := exec.Command("api_hops_start_hugo")
	startCmd.Start()

	err = startCmd.Wait()
	if err != nil {
		logger.Logger("Error: starting hugo server")
		logger.Logger(err)
		w.WriteHeader(http.StatusBadRequest)
		cleanup(cnfConfig)
		return
	}
}

func cleanup(cnfConfig model.ConfigData) {
	// Clean repos and reset cnfConfig data
	// Also invalidate github credentials
}
