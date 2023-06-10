package main

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// GetMD5Hash returns the MD5 hash of a string
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// AddHTTPScheme adds "http://" to URLs that do not have a scheme
func AddHTTPScheme(urls []string) []string {
	var result []string

	for _, url := range urls {
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			result = append(result, url)
		} else {
			result = append(result, "http://"+url)
		}
	}

	return result
}
