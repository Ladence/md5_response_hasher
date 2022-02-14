package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/Ladence/md5_response_hasher/internal/hash"
	"github.com/Ladence/md5_response_hasher/internal/transport"
)

func main() {
	parallel := flag.Int("parallel", 10, "constrains maximum number of requests which could be run concurrently")
	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("usage: ./myhttp [-parallel] urls ...")
		flag.PrintDefaults()
		os.Exit(1)
	}

	concurrently := make(chan struct{}, *parallel)
	client := transport.NewClient()
	wg := sync.WaitGroup{}
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			concurrently <- struct{}{}
			workUnit(urls[i], client)
			<-concurrently
		}(i)
	}
	wg.Wait()
}

func workUnit(url string, client *transport.Client) {
	b, err := client.SendRequest(url)
	if err != nil {
		fmt.Printf("url: %v processing ended with error: %v: ", url, err)
		return
	}
	m5 := hash.Md5{}
	h := m5.Hash(b)
	fmt.Println(url + " " + h)
}
