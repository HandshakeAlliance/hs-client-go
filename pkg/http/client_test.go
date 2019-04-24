package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestConfig(t *testing.T) {
	hostname := "http://localhost:13037"
	config := New(hostname)

	if config.Hostname != hostname {
		t.Error("Hostname not set correctly")
	}
}

//Currently this test expects localhost to exist
func TestCall(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/fee")
		// Send response to be tested
		rw.Write([]byte(`0`))
	}))

	// Close the server when test finishes
	defer server.Close()

	url := server.URL + "/fee"

	var fee float64

	err := call(url, &fee)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestCallFail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		// equals(t, req.URL.String(), "/fees")
		// Send response to be tested
		// rw.Write([]byte(`0`))
	}))

	// Close the server when test finishes
	defer server.Close()

	url := server.URL + "/fee"

	var fee float64

	err := call(url, &fee)

	if err == nil {
		t.Errorf("Error: %s", err)
	}

}
