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

// Register all routes once
func init() {
	SetupRouter()
}
// Setup/Connect to test DB and run cleanup
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

func TestGetURLEntries(t *testing.T) {
	ts := httptest.NewServer(GetRouter()) 
	// Shut down server and block until all requests go through
	defer ts.Close()

	urlEntryRequest1, _ := json.Marshal(models.URLEntry{
		Slug: "test1",
		TargetURL: "https://google.com",
	})
	urlEntryRequest2, _ := json.Marshal(models.URLEntry{
		Slug: "test2",
		TargetURL: "https://testurl.com",
	})
	resp1, err := http.Post(fmt.Sprintf("%s/v1/create", ts.URL), "application/json", bytes.NewBuffer(urlEntryRequest1))
	resp2, err := http.Post(fmt.Sprintf("%s/v1/create", ts.URL), "application/json", bytes.NewBuffer(urlEntryRequest2))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp1.StatusCode != 200 && resp2.StatusCode != 200 {
		t.Fatalf("Expected status code 200 on URLEntry creations, got %v %v", resp1.StatusCode, resp2.StatusCode)
	}

	urlEntriesResp, err := http.Get(fmt.Sprintf("%s/v1/urls", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	var urlEntries []models.URLEntry
	body, err := ioutil.ReadAll(urlEntriesResp.Body)
	if err != nil {
		t.Fatalf("Expected no error with reading data, got %v", err)
	}
	err = json.Unmarshal(body, &urlEntries)
	if err != nil {
		t.Fatalf("Expected no error with marshalling data, got %v", err)
	}
	if len(urlEntries) != 2 {
		t.Fatalf("Expected length of URLEntries slice to be 2, got %v", len(urlEntries))
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
