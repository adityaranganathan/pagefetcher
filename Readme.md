## Go Page Fetcher

The Go Page Fetcher is a tool that allows you to fetch web pages concurrently and retrieve the URLs along with the MD5 hash of the response body. It supports both URLs with and without the URL scheme, defaulting to HTTP if no scheme is provided.

### Setup

1. Build the tool by running the command: ```go build```
2. Execute the tool with the desired URLs: ```./myhttp url1 url2```

### Usage

To use the tool, provide URLs as arguments. For example, to fetch example.com and google.com, run the following command:
```bash
$ ./myhttp example.com google.com
```
The tool will display the fetched URLs along with their corresponding MD5 hashes of the response body:
```
http://example.com 84238dfc8092e5d9c0dac8ef93371a07
http://google.com c1712dc00e8a5c96137949ada1fad872
```

#### Flags

The tool also supports a flag to specify the maximum number of parallel requests to fetch URLs. If not provided, the default value is set to 10.
```bash
$ ./myhttp -parallel=2 example.com google.com yandex.com reddit.com
```