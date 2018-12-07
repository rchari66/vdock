package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"misc/logger"
	"model"
	"net/http"
)

type GitHubService struct{}

// return github repos for given users
func (ghService *GitHubService) GetRepos(w http.ResponseWriter, r *http.Request) {
	userIds, ok := r.URL.Query()["userId"]

	if !ok || len(userIds[0]) < 1 {
		logger.Logger("Error : invalid user id")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, err := http.Get("https://api.github.com/users/" + userIds[0] + "/repos")
	if err != nil {
		logger.Logger(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		logger.Logger(err)
	}

	if res.StatusCode != http.StatusOK {
		logger.Logger(res.StatusCode)
		w.WriteHeader(res.StatusCode)
		return
	}

	repos := []model.Repo{}
	err = json.Unmarshal(body, &repos)

	if err != nil {
		logger.Logger(err)
	}

	b, err := json.Marshal(repos)
	if err != nil {
		logger.Logger(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(b))
}
