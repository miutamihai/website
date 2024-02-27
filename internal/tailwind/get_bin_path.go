package tailwind

import (
	"os"
	"path"
)

func getBinPath() (string, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return path.Join(cwd, FILE_PATH), nil
}
