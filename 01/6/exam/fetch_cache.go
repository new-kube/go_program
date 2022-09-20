package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("http get: %v\n", err)
			continue
		}

		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("fetch while reading %s: %v\n", url, err)
			continue
		}
		fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
	}
}
