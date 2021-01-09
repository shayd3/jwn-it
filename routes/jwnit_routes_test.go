package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/shayd3/jwn-it/data"
)

func TestPingRoute(t *testing.T) {
	defer t.Cleanup(cleanup)
	// Stand up Test DB
	data.ConnectDatabase("jwnit-test.db", 0600)
	defer data.DB.Close()
	
	// Inject router into a test server
	SetupRouter()
	ts := httptest.NewServer(GetRouter()) 
	// Shut down server and block until all requests go through
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/v1/ping", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	
}

func cleanup() {
	err := os.Remove("jwnit-test.db")
	if err != nil {
		log.Fatal(err)
	}
}
