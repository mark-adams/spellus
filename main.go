package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var path = flag.String("path", ".", "the path to scan")

func processFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("error reading %s: %s", path, err)
	}

	scanner := bufio.NewScanner(file)
	lineno := 0
	for scanner.Scan() {
		lineno++
		line := scanner.Text()
		for k, v := range dictionary {
			if strings.Contains(line, k) {
				fmt.Printf("%s:%d: %s: '%s' should be spelled '%s'\n", path, lineno, v.ID, k, v.CorrectSpelling)
			}
		}
	}
}

func main() {
	flag.Parse()
	log.SetPrefix("")

	ch := make(chan string)
	wg := sync.WaitGroup{}

	// Read through all the paths in the channel
	// and process them to look for matching terms
	go func() {
		for path := range ch {
			go func(path string) {
				processFile(path)
				wg.Done()
			}(path)
		}
	}()

	// Walk the filesystem and place each path into the channel
	filepath.Walk(*path, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		wg.Add(1)
		ch <- path
		return nil
	})

	// Close the channel to signal that there are no more
	// paths that will be added
	close(ch)

	// Wait until all the paths have been processed
	wg.Wait()
}
