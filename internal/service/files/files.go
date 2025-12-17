package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Files struct {
}

/*
Web Crypto, Websocket, WebRTC, Locks, Workers, Cache API
*/

func New() *Files {
	return &Files{}
}

func (s *Files) Save(id, filename string, fileData io.Reader) error {
	pathToDir := filepath.Join("./uploads", id)

	if _, err := os.Stat(pathToDir); os.IsNotExist(err) {
		if err := os.MkdirAll(pathToDir, 0755); err != nil {
			return fmt.Errorf("failed to create direcotyry")
		}
	} else if err != nil {
		return fmt.Errorf("failed to check direcory: %w", err)
	} else {
		fmt.Println("The directory is already exists")
	}

	fullPathToFile := filepath.Join(pathToDir, filename)

	file, err := os.Create(fullPathToFile)
	if err != nil {
		return fmt.Errorf("failed to create file")
	}

	defer file.Close()

	if _, err := io.Copy(file, fileData); err != nil {
		return fmt.Errorf("failed to copy fileData")
	}

	return nil
}
