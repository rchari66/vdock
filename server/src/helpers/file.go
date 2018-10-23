package helpers

import "os"

// path : is a relative path to a file from src
// returns exact path of the file
func GetExactPath(path string) string {
	if path == "" {
		return ""
	}

	return os.Getenv("SERVER_PATH") + "/src/" + path
}
