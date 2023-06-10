## Go Page Fetcher

This tool fetches web pages concurrently and prints the URLs and the MD5 hash of the response body. 
The urls can be supplied with or without the url scheme. 
The tool will default to http if no scheme is provided.

### Setup

1. ```go build```
2. ```./myhttp url1 url2```

### Example

Given two URLs, `example.com` and `google.com`, running the tool will look like:

```bash
$ ./myhttp example.com google.com
http://example.com 84238dfc8092e5d9c0dac8ef93371a07
http://google.com c1712dc00e8a5c96137949ada1fad872
```