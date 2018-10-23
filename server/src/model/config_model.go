package model

// Configure
type ConfigData struct {
	UserId   string `json:"userId"`
	BlogRepo string `json:"blogRepo"`
	SiteRepo string `json:"siteRepo"`
	Password string `json:"password"`
	Email    string `json:"email"`
	// BlogBranch string `json:"blogBranch"`
	// SiteBranch string `json:"siteBranch"`
}
