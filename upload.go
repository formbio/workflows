package upload

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Exec(repoDir string) error {
	err := filepath.Walk(repoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, "/workflows/") {
			fmt.Println(path)
		}

		return nil
	})
	if err != nil {
		log.Fatalf("unable to walk current dir: %s", err)
	}
}
