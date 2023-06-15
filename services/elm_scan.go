package services

import (
	"awesomeProject/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ElmScan(basePath string) {
	var fileCount uint64 = 0
	tx := models.DB.Begin()

	scanDir := func(fullPath string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if !info.IsDir() {
			var parent string
			var path = filepath.Dir(fullPath)
			if path == basePath {
				parent = "__ROOT__"
			} else {
				parent = filepath.Base(filepath.Dir(path))
			}
			fmt.Println(parent)
			elm := models.Elm{Path: path, Dir: filepath.Base(path), ParentDir: parent, Name: info.Name(), Checksum: "", Size: uint64(info.Size()), Date: info.ModTime()}
			tx.Create(&elm)
			fileCount++
		}
		return nil
	}

	err := filepath.Walk(basePath, scanDir)
	if err != nil {
		log.Fatal(err)
		fileCount = 0
		tx.Rollback()
	}

	if fileCount > 0 {
		tx.Commit()
	}

	// scanDir(basePath, "__ROOT__")
}
