package tailwind

import "os"

func binExists() (bool, error) {
	binPath, err := getBinPath()

	if err != nil {
		// Something's _incredibly_ bad if this errors

		panic(err)
	}

	_, statErr := os.Stat(binPath)

	if statErr != nil {
		return false, nil
	}

	return true, nil
}
