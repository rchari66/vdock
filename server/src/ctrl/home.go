package ctrl

import (
	"net/http"
)

var (
	ide *IdeController = new(IdeController)
)

type HomeController struct{}

func (cf *HomeController) HomeAndIdeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./index.html")
	} else {
		ide.Ide(w, r)
	}
}
