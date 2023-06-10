package main

import "testing"

func TestGetMD5Hash(t *testing.T) {
	input := "abcd"
	expectedOutput := "e2fc714c4727ee9395f324cd2e7f331f"

	got := GetMD5Hash(input)
	if got != expectedOutput {
		t.Errorf("expected %s, got %s", expectedOutput, got)
	}
}

func TestAddHTTPScheme(t *testing.T) {
	urls := []string{"google.com", "http://facebook.com", "https://twitter.com", "yahoo.com"}
	expectedURLs := []string{"http://google.com", "http://facebook.com", "https://twitter.com", "http://yahoo.com"}

	got := AddHTTPScheme(urls)
	for i, url := range got {
		if url != expectedURLs[i] {
			t.Errorf("expected %s, got %s", expectedURLs[i], url)
		}
	}
}
