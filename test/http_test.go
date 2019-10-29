package test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// e.g. http.HandleFunc("/health-check", HealthCheckHandler)
func ATestHandler(w http.ResponseWriter, r *http.Request) {
	//	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func TestHttpRequest(t *testing.T) {
	client := &http.Client{}
	ts := httptest.NewServer(http.HandlerFunc(ATestHandler))
	defer ts.Close()
	fmt.Println(ts.URL)
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("resp error")
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
}
