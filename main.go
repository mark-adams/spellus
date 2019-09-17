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

var basepath = flag.String("path", ".", "the path to scan")
var exclude = flag.String("exclude", ".git", "path prefixes to exclude from the scan")

func processFile(path string) {
	file, err := os.Open(filepath.Join(*basepath, path))
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
		for fpath := range ch {
			go func(fpath string) {
				processFile(fpath)
				wg.Done()
			}(fpath)
		}
	}()

	prefixes := strings.Split(*exclude, ",")

	// Walk the filesystem and place each path into the channel
	filepath.Walk(*basepath, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		relpath, err := filepath.Rel(*basepath, path)
		if err != nil {
			panic(err)
		}

		for _, prefix := range prefixes {
			if strings.HasPrefix(relpath, prefix) {
				return nil
			}
		}

		wg.Add(1)
		ch <- relpath
		return nil
	})

	// Close the channel to signal that there are no more
	// paths that will be added
	close(ch)

	// Wait until all the paths have been processed
	wg.Wait()
}
