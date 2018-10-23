package ctrl

import (
	"fmt"
	"net/http"
)

type ThemeController struct{}

func (cf *ThemeController) GetThemes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
