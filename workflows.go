package workflows

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var mainFound bool

func Exec(repoDir string) error {
	if repoDir == "" || repoDir == "." {
		repoDir, _ = os.Getwd()
	}
	err := filepath.Walk(repoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == ".github" {
			return filepath.SkipDir
		}

		if strings.Contains(path, "/workflows/") {
			fmt.Println(path)
		}
		mainFound = mainFound || info.Name() == "main.nf"

		return nil
	})
	if err != nil {
		log.Fatalf("unable to walk current dir: %s", err)
	}

	if !mainFound {
		log.Fatalln("main.nf not found")
	}
	return nil
}
