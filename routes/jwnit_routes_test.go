package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/shayd3/jwn-it/data"
	"github.com/shayd3/jwn-it/models"
)
  
func init() {
	// Inject router into a test server
	SetupRouter()
}
func TestMain(m *testing.M) {
	// Stand up Test DB
	data.ConnectDatabase("jwnit-test.db", 0600)
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func cleanup() {
	data.DB.Close()
	err := os.Remove("jwnit-test.db")
	if err != nil {
		log.Fatal(err)
	}
}

func TestAddURLEntryRoute(t *testing.T) {
	ts := httptest.NewServer(GetRouter()) 
	// Shut down server and block until all requests go through
	defer ts.Close()

	urlEntryRequest, _ := json.Marshal(models.URLEntry{
		Slug: "test",
		TargetURL: "google.com",
	})
	
	resp, err := http.Post(fmt.Sprintf("%s/v1/create", ts.URL), "application/json", bytes.NewBuffer(urlEntryRequest))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	var urlEntry models.URLEntry
	body, err := ioutil.ReadAll(resp.Body)
 	
	if json.Unmarshal(body, &urlEntry) != nil {
		t.Fatalf("Expected successful unmarshal to URLEntry object, got %v", err)
	}

	if !strings.Contains(urlEntry.TargetURL, "https://") {
		t.Fatalf("Expected TargetURL to contain 'https://', got %v", urlEntry.TargetURL)
	}
}
