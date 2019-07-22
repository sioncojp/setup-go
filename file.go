package main

import (
	"log"
	"os"
)

// Get ... get all file name
func (f *Files) Get(rn reponame) []string {
	return []string{
		"README.md",
		"bin/.gitkeep",
	}
}

// Create ...create files
func (f *Files) Create(rn reponame) {
	files := f.Get(rn)
	for _, file := range files {
		// check File exist
		if FileExist(file) {
			continue
		}

		func() {
			f, err := os.Create(file)
			if err != nil {
				log.Printf("[WARN] cannot create %s file: %s\n", file, err)
				return
			}
			defer f.Close()

			log.Printf("[INFO] create file: %s\n", file)
		}()
	}
}
