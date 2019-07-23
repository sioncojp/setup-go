package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Get ... get all dir name
func (d *Dirs) Get(rn reponame) []string {
	return []string{
		"bin/",
	}
}

// Create ...create directroy
func (d *Dirs) Create(rn reponame) {
	dirs := d.Get(rn)
	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Dir(dir), 0755)
		if err != nil {
			log.Printf("[WARN] cannot create %s dir: %s\n", dir, err)
			continue
		}
		log.Printf("[INFO] create dir: %s\n", dir)
	}
}
