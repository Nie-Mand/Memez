package common

import (
	"fmt"
	"log"
	"os"
)

const (
	LOGO_PATH = "hack/ascii.txt"
)

func DisplayLogo() error {
	ascii, err := os.ReadFile(LOGO_PATH)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(string(ascii))
	return nil
}
