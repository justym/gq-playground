package errors

import (
	"log"
	"os"
)

func Check(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s\n", msg, err)
		os.Exit(1)
	}
}
