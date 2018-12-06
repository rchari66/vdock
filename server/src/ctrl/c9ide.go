package ctrl

import (
	"fmt"
	"misc/logger"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type IdeController struct{}

// init cloud9 ide proxy
var ideProxy = initIdeProxy()

func initIdeProxy() *httputil.ReverseProxy {
	logger.Logger("Code proxy initializing.. ")

	// Url for code editor
	codeUrl, _ := url.Parse("http://localhost:8282/")

	codeDirectory := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", codeUrl.Host)
		req.URL.Scheme = "http"
		req.URL.Host = codeUrl.Host
	}

	ideProxy := &httputil.ReverseProxy{Director: codeDirectory}
	return ideProxy
}

func (ic *IdeController) Ide(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("ALLOW_PUBLIC_ACCESS") != "YES" && !strings.Contains(r.Host, "localhost:") {
		fmt.Fprintf(w, " Sorry Access is blocked!! ")
		return
	}
	// server ide traffic
	ideProxy.ServeHTTP(w, r)
}
