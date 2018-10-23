package ctrl

import (
	"misc/logger"
	"net/http"
	"os"
)

var (
	home    *HomeController    = new(HomeController)
	conf    *ConfController    = new(ConfController)
	publish *PublishController = new(PublishController)
	execute *ExecController    = new(ExecController)
	help    *HelpController    = new(HelpController)
	theme   *ThemeController   = new(ThemeController)
)

func Setup() {

	logger.Logger("Configuring controllers..")

	// Server static content like css, js, and images
	http.HandleFunc("/public/", serveStaticFile)

	// Api for both home apge and cloud9 ide
	http.HandleFunc("/", home.HomeAndIdeHandler)

	// Api to execute commands
	http.HandleFunc("/api/v1/exec/", execute.ExecHandler)

	http.HandleFunc("/api/v1/publish/", publish.PublishHandler)

	http.HandleFunc("/api/v1/conf/", conf.ConfigHandler)

	http.HandleFunc("/api/v1/help/", help.GetHelp)

	http.HandleFunc("/api/v1/themes/", theme.GetThemes)

	logger.Logger("Controllers are configured!")

}

// Validate r.URL.Path for security
func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = os.Getenv("SERVER_PATH") + r.URL.Path
	http.ServeFile(w, r, r.URL.Path)
}

/*  TODO: Setup router using mux
func Setup( ) {
r := mux.NewRouter()
r.HandleFunc("/", home.GetHome)
// Api to set config
r.HandleFunc("/api/v1/conf", conf.Configure).Methods("POST")
// Api to execute commands(custom);
r.HandleFunc("/api/v1/exec/{cmd:\\d+}", execute.Execute).Methods("GET")
http.Handle("/", r)
 } */
