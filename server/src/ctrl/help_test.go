package ctrl

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHome(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/api/v1/help/", nil)
	if err != nil {
		t.Fatal(err)
	}

	help := new(HelpController)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(help.GetHelp)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	handler.ServeHTTP(rr, req)

	// Check the status.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
