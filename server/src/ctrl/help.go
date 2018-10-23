package ctrl

import (
	"fmt"
	"net/http"
)

type HelpController struct{}

func (cf *HelpController) GetHelp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
