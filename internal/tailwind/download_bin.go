package tailwind

import (
	"errors"
	"io"
	"net/http"
	"os"
	"runtime"
)

func downloadBin() error {
	var url string

	if runtime.GOOS == "darwin" {
		url = MACOS_TAILWIND_BIN_URL
	} else {
		url = LINUX_TAILWIND_BIN_URL
	}

	response, err := http.Get(url)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("Failed to get tailwind bin")
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	binFile, err := os.Create(FILE_PATH)

	if err != nil {
		return err
	}

	binFile.Write(data)
	binFile.Chmod(0777)
	defer binFile.Close()

	return nil
}
