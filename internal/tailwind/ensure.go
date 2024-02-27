package tailwind

func Ensure() error {
	exists, err := binExists()

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	return downloadBin()
}
