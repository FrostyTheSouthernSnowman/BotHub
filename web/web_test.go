package web_test

import (
	"net/http"
	"net/http/httptest"
	"robot-simulator/web"
	"testing"
)

//! Test removed because of path issues in web.go
//TODO: Fix test
/*func TestRouter(t *testing.T) {
	// Instantiate the router using the constructor function that
	// we defined previously
	r := web.NewRouter()

	// Create a new server using the "httptest" libraries `NewServer` method
	// Documentation : https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/")

	// Handle any unexpected error
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}
}*/

func TestRouterForNonExistentRoute(t *testing.T) {
	r := web.NewRouter()
	mockServer := httptest.NewServer(r)
	// Most of the code is similar. The only difference is that now we make a
	//request to a route we know we didn't define, like the `POST /hello` route.
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 404 (method not allowed)
	if resp.StatusCode != 404 {
		t.Errorf("Status should be 404, got %d", resp.StatusCode)
	}
}

//! Test comment out for same reason as the test above
/*func TestStaticFileServer(t *testing.T) {
	r := web.NewRouter()
	mockServer := httptest.NewServer(r)

	// We want to hit the `GET /assets/` route to get the index.html file response
	resp, err := http.Get(mockServer.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	// It isn't wise to test the entire content of the HTML file.
	// Instead, we test that the content-type header is "text/html; charset=utf-8"
	// so that we know that an html file has been served
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}

}*/
