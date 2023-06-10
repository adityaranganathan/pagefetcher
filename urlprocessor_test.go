package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// TestProcessURLs tests the ProcessURLs function using a mock PageFetcher and ResultPrinter.
func TestProcessURLs(t *testing.T) {
	mockFetcher := func(url string) ([]byte, error) {
		return []byte("response body of " + url), nil
	}

	gotResults := make(map[string]string)
	var mu sync.Mutex
	mockPrinter := func(url string, result []byte) {
		mu.Lock()
		defer mu.Unlock()
		gotResults[url] = string(result)
	}

	urls := []string{"url1", "url2", "url3", "url4", "url5"}
	ProcessURLs(urls, 5, mockFetcher, mockPrinter)

	for _, url := range urls {
		expectedResult := "response body of " + url
		if gotResults[url] != expectedResult {
			t.Errorf("expected %q, got %q", expectedResult, gotResults[url])
		}
	}
}

// TestFetchPage tests the FetchPage function using a mock HTTP server.
func TestFetchPage(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test response"))
	}))
	defer mockServer.Close()

	gotResult, err := FetchPage(mockServer.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedResult := "test response"
	if string(gotResult) != expectedResult {
		t.Fatalf("expected %q, got %q", expectedResult, string(gotResult))
	}
}
