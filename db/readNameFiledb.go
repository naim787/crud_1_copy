package db 

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFileNamedb() {
  entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	found := false

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".db" {
			fmt.Println("Ditemukan file:", entry.Name())
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada file .db di direktori saat ini")
	}
}